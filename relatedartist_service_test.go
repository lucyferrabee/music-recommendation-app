package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zmb3/spotify"
)

type RelatedArtistServiceTest struct {
	Client SpotifyClient
}

func (s *RelatedArtistServiceTest) getFirstRelatedArtistByTrackId(id string) []relatedartist {
	client := s.Client

	relatedArtists, err := client.GetRelatedArtists(spotify.ID(id))
	if err != nil {
		log.Fatalf("error retrieving related artist data: %v", err)
	}

	artists := make([]relatedartist, len(relatedArtists))

	for i, a := range relatedArtists {
		topTracks, err := client.GetArtistsTopTracks(a.ID, "US")
		if err != nil {
			log.Fatalf("error retrieving top tracks data: %v", err)
		}

		artists[i] = relatedartist{
			name:       a.Name,
			id:         string(a.ID),
			popularity: a.Popularity,
			topTracks:  topTracks,
		}
	}

	return artists
}

// your test function
func TestGetTopTracksFromRelatedArtists(t *testing.T) {
	mockClient := new(MockSpotifyClient)

	service := &RelatedArtistServiceTest{
		Client: mockClient,
	}

	// Set up expected data for testing
	expectedRelatedArtists := []spotify.FullArtist{
		{ID: "related-artist-1", Name: "Related Artist 1"},
		{ID: "related-artist-2", Name: "Related Artist 2"},
	}
	expectedTopTracks := []spotify.FullTrack{
		{ID: "track-1", Name: "Track 1", Popularity: 80},
		{ID: "track-2", Name: "Track 2", Popularity: 70},
	}

	mockClient.On("GetRelatedArtists", mock.AnythingOfType("spotify.ID")).Return(expectedRelatedArtists, nil)
	mockClient.On("GetArtistsTopTracks", mock.AnythingOfType("spotify.ID"), mock.AnythingOfType("string")).Return(expectedTopTracks, nil)

	tracks := service.getByArtistId("1a2b3c4d5e6f7g")

	// Assert
	mockClient.AssertExpectations(t)
	assert.NotNil(t, tracks, "Expected non-nil result")
	assert.NoError(t, err, "Unexpected error")

}
