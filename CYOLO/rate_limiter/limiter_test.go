package rate_limiter_test

import (
	"CYOLO/rate_limiter"
	"github.com/onsi/gomega"
	"testing"
)

func TestRateLimiter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	allowedPerSec := 2
	allowedPerMin := 10
	limiter := rate_limiter.NewRateLimiter(allowedPerSec, allowedPerMin)

	for i := 0; i < allowedPerSec; i++ {
		g.Expect(limiter.Allow()).To(gomega.Equal(true))
	}
	g.Expect(limiter.Allow()).To(gomega.Equal(false))
}
