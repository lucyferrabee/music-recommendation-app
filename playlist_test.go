package main

import (
	"math"
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
		{Popularity: 40},
		{Popularity: 60},
	}

	result := ps.chooseSimilarPopularity(tracks, 70, 10)
	assert.Len(t, result, 3, "Expected 3 tracks")

	result2 := ps.chooseSimilarPopularity(tracks, 90, 9)
	assert.Len(t, result2, 1, "Expected 1 track")

	result3 := ps.chooseSimilarPopularity(tracks, 60, 6)
	assert.Len(t, result3, 2, "Expected 2 tracks")

	for _, track := range result {
		assert.True(t, math.Abs(float64(track.Popularity)-70) <= 10, "Track popularity should be within the threshold")
	}
}

func TestRemoveDuplicatesRemovesDuplicateWhenOneExists(t *testing.T) {
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
