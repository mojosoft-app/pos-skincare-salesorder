package middleware

import (
	"net/http"
	"sync"
	"time"

	"pos-mojosoft-so-service/internal/config"
	"pos-mojosoft-so-service/internal/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type RateLimiter struct {
	visitors map[string]*visitor
	mutex    *sync.RWMutex
	config   *config.RateLimitConfig
}

func NewRateLimiter(config *config.RateLimitConfig) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		mutex:    &sync.RWMutex{},
		config:   config,
	}

	go rl.cleanupVisitors()
	return rl
}

func (rl *RateLimiter) addVisitor(ip string) *rate.Limiter {
	limiter := rate.NewLimiter(rate.Limit(rl.config.RequestsPerMinute)/60, rl.config.BurstSize)

	rl.mutex.Lock()
	rl.visitors[ip] = &visitor{limiter, time.Now()}
	rl.mutex.Unlock()

	return limiter
}

func (rl *RateLimiter) getVisitor(ip string) *rate.Limiter {
	rl.mutex.Lock()
	v, exists := rl.visitors[ip]
	if !exists {
		rl.mutex.Unlock()
		return rl.addVisitor(ip)
	}

	v.lastSeen = time.Now()
	rl.mutex.Unlock()
	return v.limiter
}

func (rl *RateLimiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		rl.mutex.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(rl.visitors, ip)
			}
		}
		rl.mutex.Unlock()
	}
}

func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := rl.getVisitor(c.ClientIP())
		if !limiter.Allow() {
			utils.ErrorResponse(c, http.StatusTooManyRequests, "Rate limit exceeded", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func RateLimitMiddleware(config *config.RateLimitConfig) gin.HandlerFunc {
	limiter := NewRateLimiter(config)
	return limiter.Limit()
}
