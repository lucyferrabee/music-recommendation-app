package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zmb3/spotify"
)

func TestChooseSimilarPopularity(t *testing.T) {
	ps := &PlaylistService{}

	tracks := []spotify.FullTrack{
		{Popularity: 81},
		{Popularity: 75},
		{Popularity: 65},
	}

	result := ps.chooseSimilarPopularity(tracks, 70, 10)

	assert.Len(t, result, 2, "Expected 2 tracks")

	for _, track := range result {
		assert.True(t, abs(track.Popularity-70) <= 10, "Track popularity should be within the threshold")
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tracks := []spotify.FullTrack{
		{ID: spotify.ID("1")},
		{ID: spotify.ID("2")},
		{ID: spotify.ID("1")},
		{ID: spotify.ID("3")},
	}

	// Call the function being tested
	result := removeDuplicates(tracks)

	// Assertions
	assert.Len(t, result, 3, "Expected 3 unique tracks")

	// Check if the result does not contain duplicates
	uniqueIDs := make(map[spotify.ID]struct{})
	for _, track := range result {
		assert.NotContains(t, uniqueIDs, track.ID, "Duplicate track found")
		uniqueIDs[track.ID] = struct{}{}
	}
}

// TestAbs tests the abs function
func TestAbs(t *testing.T) {
	// Call the function being tested with positive and negative values
	result1 := abs(5)
	result2 := abs(-5)

	// Assertions
	assert.Equal(t, 5, result1, "Absolute value of 5 should be 5")
	assert.Equal(t, 5, result2, "Absolute value of -5 should be 5")
}
