package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test01_Calculate(t *testing.T) {
	for change := 0.25; change < 10000; change += 0.25 {
		storage := NewStorageSystemWithName("S01")
		storage.initBankNoteStorageWithDefault()
		sol, canChange := calculateSolutionFromChange(storage, float64(change))
		if !canChange {
			t.Fail()
		}

		sumChange := 0.00
		for _, item := range sol.BanknoteChange {
			sumChange += (float64(item.Quantity) * item.Value)
		}

		assert.Equal(t, change, sumChange)
	}
}
