package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSummary(t *testing.T) {
	s := NewSummary()

	require.Equal(t, int64(0), s.failures)
	require.Equal(t, int64(0), s.total)
	require.Equal(t, 0, len(s.successes))
}

func TestSummary_AddSuccess(t *testing.T) {
	s := NewSummary()

	s.AddSuccess("Yshtola", 1000)

	require.Equal(t, int64(0), s.failures)
	require.Equal(t, int64(1000), s.summarizeSuccess())
	require.Equal(t, "Yshtola", s.successes[0].donor)
}

func TestSummary_AddToFailure(t *testing.T) {
	s := NewSummary()

	s.AddToFailure(int64(1000))

	require.Equal(t, int64(1000), s.failures)
	require.Equal(t, int64(0), s.summarizeSuccess())
	require.Equal(t, 0, len(s.successes))
}

func TestSummary_SummarizeSuccess(t *testing.T) {
	s := NewSummary()

	s.AddToFailure(int64(1000))
	s.AddSuccess("Yshtola", 1000)
	s.AddSuccess("Lyse", 2000)

	require.Equal(t, int64(3000), s.summarizeSuccess())
}

func TestSummary_AverageSuccess(t *testing.T) {
	s := NewSummary()

	s.AddToFailure(int64(1000))
	s.AddSuccess("Yshtola", 1000)
	s.AddSuccess("Lyse", 2000)

	require.Equal(t, float64(1500), s.averageSuccess())
}

func TestSummary_TopDonor(t *testing.T) {
	s := NewSummary()

	s.AddToFailure(int64(1000))
	s.AddSuccess("Yshtola", 1000)
	s.AddSuccess("Lyse", 2000)

	require.Equal(t, "Lyse", s.topDonor())
}
