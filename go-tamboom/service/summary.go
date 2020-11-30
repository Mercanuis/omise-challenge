package service

import (
	"fmt"
)

// ChargeSuccess represents a successful donation charge made
type ChargeSuccess struct {
	donor  string
	amount int64
}

// Summary represents a summary of the entire transaction process
type Summary struct {
	failures  int64
	total     int64
	successes []ChargeSuccess
}

// NewSummary returns a new Summary
func NewSummary() Summary {
	return Summary{
		failures:  int64(0),
		total:     int64(0),
		successes: make([]ChargeSuccess, 0),
	}
}

// AddSuccess creates a new ChargeSuccess and adds it to the Summary
func (s *Summary) AddSuccess(donor string, amount int64) {
	s.total = s.total + amount
	s.successes = append(s.successes, ChargeSuccess{
		donor:  donor,
		amount: amount,
	})
}

// AddToFailure adds the failed donation amount to the Summary
func (s *Summary) AddToFailure(amount int64) {
	s.total = s.total + amount
	s.failures = s.failures + amount
}

// PrintSummary prints the data of the Summary
func (s *Summary) PrintSummary() {
	fmt.Printf("[Total Donations Received]: JPY %d\n", s.total)
	fmt.Printf("[Successful Donations]: JPY %d\n", s.summarizeSuccess())
	fmt.Printf("[Failed Donations]: JPY %d\n", s.failures)

	fmt.Printf("[Average Per Person]: JPY %f\n", s.averageSuccess())
	fmt.Printf("[Top Donor]: %s\n", s.topDonor())
}

func (s *Summary) summarizeSuccess() int64 {
	total := int64(0)
	for _, sum := range s.successes {
		total = total + sum.amount
	}
	return total
}

func (s *Summary) averageSuccess() float64 {
	avg := int64(0)
	for _, sum := range s.successes {
		avg = avg + sum.amount
	}

	return float64(avg / int64(len(s.successes)))
}

func (s *Summary) topDonor() string {
	topDonor := s.successes[0]
	for _, donor := range s.successes {
		if donor.amount > topDonor.amount {
			topDonor = donor
		}
	}
	return topDonor.donor
}
