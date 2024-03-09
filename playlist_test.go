package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zmb3/spotify"
)

func TestChooseSimilarPopularity(t *testing.T) {
	ps := &PlaylistService{}

	tests := []struct {
		tracks           []spotify.FullTrack
		targetPopularity int
		threshold        int
		expectedCount    int
	}{
		// All tracks being within the threshold
		{
			tracks: []spotify.FullTrack{
				{Popularity: 81},
				{Popularity: 75},
				{Popularity: 65},
				{Popularity: 60},
			},
			targetPopularity: 70,
			threshold:        10,
			expectedCount:    3,
		},
		// No tracks being within the threshold
		{
			tracks: []spotify.FullTrack{
				{Popularity: 20},
				{Popularity: 30},
				{Popularity: 40},
				{Popularity: 50},
			},
			targetPopularity: 70,
			threshold:        10,
			expectedCount:    0,
		},
		// One track being within the threshold
		{
			tracks: []spotify.FullTrack{
				{Popularity: 69},
				{Popularity: 76},
				{Popularity: 80},
				{Popularity: 85},
			},
			targetPopularity: 70,
			threshold:        5,
			expectedCount:    1,
		},
		// // One track's popularity being exactly above the threshold
		{
			tracks: []spotify.FullTrack{
				{Popularity: 81},
				{Popularity: 70},
				{Popularity: 65},
			},
			targetPopularity: 70,
			threshold:        10,
			expectedCount:    2,
		},
		// // One track's popularity being exactly below the threshold
		{
			tracks: []spotify.FullTrack{
				{Popularity: 59},
				{Popularity: 70},
				{Popularity: 65},
			},
			targetPopularity: 70,
			threshold:        10,
			expectedCount:    2,
		},
		// // Zero threshold (only tracks with same popularity will match)
		{
			tracks: []spotify.FullTrack{
				{Popularity: 70},
				{Popularity: 75},
				{Popularity: 80},
				{Popularity: 85},
			},
			targetPopularity: 70,
			threshold:        0,
			expectedCount:    1,
		},
		// // Stupidly large threshold
		{
			tracks: []spotify.FullTrack{
				{Popularity: 20},
				{Popularity: 30},
				{Popularity: 40},
				{Popularity: 50},
			},
			targetPopularity: 70,
			threshold:        100,
			expectedCount:    4,
		},
		// // Negative thresholds (should be treated as zero threshold)
		{
			tracks: []spotify.FullTrack{
				{Popularity: 70},
				{Popularity: 75},
				{Popularity: 80},
				{Popularity: 85},
			},
			targetPopularity: 70,
			threshold:        -10,
			expectedCount:    1,
		},
		// // Negative popularities (should be treated as zero popularity)
		{
			tracks: []spotify.FullTrack{
				{Popularity: -5},
				{Popularity: 75},
				{Popularity: 80},
				{Popularity: 85},
			},
			targetPopularity: -10,
			threshold:        10,
			expectedCount:    0,
		},
		// // Empty track list
		{
			tracks:           []spotify.FullTrack{},
			targetPopularity: 70,
			threshold:        10,
			expectedCount:    0,
		},
	}

	for _, test := range tests {
		result := ps.chooseSimilarPopularity(test.tracks, test.targetPopularity, test.threshold)
		assert.Len(t, result, test.expectedCount, "Unexpected number of tracks")
		for _, track := range result {
			assert.True(t, math.Abs(float64(track.Popularity)-float64(test.targetPopularity)) <= math.Abs(float64(test.threshold)), "Track popularity should be within the threshold")
		}
	}
}

func TestRemovesDuplicateWhenOneExists(t *testing.T) {
	ps := &PlaylistService{}

	tracks := []spotify.FullTrack{
		{
			SimpleTrack: spotify.SimpleTrack{ID: "1"},
			Popularity:  80,
		},
		{
			SimpleTrack: spotify.SimpleTrack{ID: "2"},
			Popularity:  70,
		},
		{
			SimpleTrack: spotify.SimpleTrack{ID: "3"},
			Popularity:  70,
		},
		{
			SimpleTrack: spotify.SimpleTrack{ID: "4"},
			Popularity:  70,
		},
		{
			SimpleTrack: spotify.SimpleTrack{ID: "4"},
			Popularity:  70,
		},
	}

	result := ps.removeDuplicates(tracks)

	// Assert
	assert.Len(t, result, 4, "Expected 4 unique tracks")

	uniqueIDs := make(map[spotify.ID]struct{})
	for _, track := range result {
		assert.NotContains(t, uniqueIDs, track.ID, "Duplicate track found")
		uniqueIDs[track.ID] = struct{}{}
	}
}

func TestRemoveDuplicatesReturnsUniqueTracksWhenAllTracksAreUnique(t *testing.T) {
	ps := &PlaylistService{}

	tracks := []spotify.FullTrack{
		{
			SimpleTrack: spotify.SimpleTrack{ID: "1"},
			Popularity:  80,
		},
		{
			SimpleTrack: spotify.SimpleTrack{ID: "2"},
			Popularity:  70,
		},
		{
			SimpleTrack: spotify.SimpleTrack{ID: "3"},
			Popularity:  70,
		},
		{
			SimpleTrack: spotify.SimpleTrack{ID: "4"},
			Popularity:  70,
		},
	}

	result := ps.removeDuplicates(tracks)

	// Assert
	assert.Len(t, result, 4, "Expected 4 unique tracks")
}
