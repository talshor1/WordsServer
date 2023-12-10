package server

import (
	blockingQueue "CYOLO/blocking_queue"
	"CYOLO/config"
	"CYOLO/data_holder"
	"CYOLO/rate_limiter"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"sync"
)

type Server struct {
	Router     *gin.Engine
	Queue      *blockingQueue.BlockingQueue
	DataHolder *data_holder.DataHolder
	mu         sync.Mutex
	config     *config.Config
	limiter    *rate_limiter.RateLimiter
}

func NewServer(queue *blockingQueue.BlockingQueue, dataHolder *data_holder.DataHolder, config *config.Config) *Server {
	s := &Server{
		Router:     gin.Default(),
		Queue:      queue,
		DataHolder: dataHolder,
		config:     config,
		limiter:    rate_limiter.NewRateLimiter(config.PerSec, config.PerMin),
	}
	s.addRoutes()

	return s
}

func (s *Server) addRoutes() {
	s.Router.POST("/words", s.handleReceive)
	s.Router.GET("/stats", s.handleStats)
}

func (s *Server) handleReceive(c *gin.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()
	words := c.Query("words")

	if !isValidCommaSeparatedString(words) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comma-separated string format"})
		return
	}

	s.Queue.Enqueue(words)
	c.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})
}

func (s *Server) handleStats(c *gin.Context) {
	if !s.limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
		return
	}

	topStr := s.DataHolder.GetTopFive()
	leastStr := s.DataHolder.GetLeast()
	medianStr := s.DataHolder.GetMedian()
	returnStr := fmt.Sprintf("Top 5 words %s, Least word %s, Median word %s", topStr, leastStr, medianStr)
	c.JSON(http.StatusOK, gin.H{"message": returnStr})
}

func (s *Server) Run() {
	log.Printf("Server is starting on %s:%s", s.config.Address, s.config.Port)
	s.Router.Run(fmt.Sprintf("%s:%s", s.config.Address, s.config.Port))
}

func isValidCommaSeparatedString(s string) bool {
	return regexp.MustCompile(`^(\w+)(,\s*\w+)*$`).MatchString(s)
}
