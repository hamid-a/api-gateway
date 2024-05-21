package circuitbreaker

import (
	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/sony/gobreaker"
)

type CircuitBreaker struct {
	cb *gobreaker.TwoStepCircuitBreaker
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(name string, cfg config.CircuitBreaker) *CircuitBreaker {
	var st gobreaker.Settings
	st.Name = name
	st.MaxRequests = cfg.MaxRequests
	st.Interval = cfg.ResetInterval
	st.Timeout = cfg.OpenTimeout

	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		if !cfg.Enabled {
			return false
		}

		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= cfg.MinRequests && failureRatio >= cfg.FailureRatioThereshold
	}

	return &CircuitBreaker{
		cb: gobreaker.NewTwoStepCircuitBreaker(st),
	}
}

func (cb *CircuitBreaker) Allow() (func(success bool), error) {
	return cb.cb.Allow()
}

func (cb *CircuitBreaker) IsOpen() bool {
	return cb.cb.State() == gobreaker.StateOpen
}
