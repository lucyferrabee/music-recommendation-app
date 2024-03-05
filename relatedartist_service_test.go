package main

import (
	"testing"
)

type RelatedArtistServiceTest struct {
	*testing.T
	Client SpotifyClient
}

// func TestGetFirstRelatedArtistByTrackId(t *testing.T) {
// 	mockClient := new(MockSpotifyClient)

// 	service := &RelatedArtistService{
// 		Client: mockClient,
// 	}

// 	expectedTopTracks := []spotify.FullTrack{
// 		{
// 			SimpleTrack: spotify.SimpleTrack{ID: "1", Name: "Track 1"},
// 			Popularity:  80,
// 		},
// 		{
// 			SimpleTrack: spotify.SimpleTrack{ID: "2", Name: "Track 2"},
// 			Popularity:  70,
// 		},
// 	}

// 	mockClient.On("GetArtistsTopTracks", mock.AnythingOfType("spotify.ID"), mock.AnythingOfType("string")).Return(expectedTopTracks, nil)

// 	artist := service.getFirstRelatedArtistByTrackId("1a2b3c4d5e6f7g")

//		// Assert
//		mockClient.AssertExpectations(t)
//		assert.NotNil(t, artist, "Expected non-nil result")
//		assert.NoError(t, err, "Unexpected error")
//	}
func TestGetTopTracksFromRelatedArtists(t *testing.T) {
	// mockClient := new(MockSpotifyClient)

	// service := &RelatedArtistServiceTest{
	// 	T:      t,
	// 	Client: mockClient,
	// }

	// expectedRelatedArtists := []spotify.FullArtist{
	// 	{
	// 		SimpleArtist: spotify.SimpleArtist{ID: "123", Name: "related artist 1"},
	// 	},
	// 	{
	// 		SimpleArtist: spotify.SimpleArtist{ID: "345", Name: "related artist 2"},
	// 	},
	// }

	// expectedTopTracks := []spotify.FullTrack{
	// 	{
	// 		SimpleTrack: spotify.SimpleTrack{ID: "1", Name: "Track 1"},
	// 		Popularity:  80,
	// 	},
	// 	{
	// 		SimpleTrack: spotify.SimpleTrack{ID: "2", Name: "Track 2"},
	// 		Popularity:  70,
	// 	},
	// }

	// mockClient.On("GetRelatedArtists", mock.AnythingOfType("spotify.ID")).Return(expectedRelatedArtists, nil)
	// mockClient.On("GetArtistsTopTracks", mock.AnythingOfType("spotify.ID"), mock.AnythingOfType("string")).Return(expectedTopTracks, nil)

	// artist := service.getFirstRelatedArtistByTrackId("1a2b3c4d5e6f7g")

	// // Assert
	// mockClient.AssertExpectations(t)
	// assert.NotNil(t, artist, "Expected non-nil result")
}
