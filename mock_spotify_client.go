package main

import (
	"github.com/zmb3/spotify"
)

type SpotifyClient interface {
	GetArtistsTopTracks(artistID spotify.ID, country string) ([]spotify.FullTrack, error)
	GetRelatedArtists(artistID spotify.ID) ([]spotify.FullArtist, error)
}

type MockSpotifyClient struct {
	TopTracks      []spotify.FullTrack
	RelatedArtists []spotify.FullArtist
}

func NewMockSpotifyClient() SpotifyClient {
	return &MockSpotifyClient{}
}

func (m *MockSpotifyClient) GetArtistsTopTracks(artistID spotify.ID, country string) ([]spotify.FullTrack, error) {
	return m.TopTracks, nil
}

func (m *MockSpotifyClient) GetRelatedArtists(artistID spotify.ID) ([]spotify.FullArtist, error) {
	return m.RelatedArtists, nil
}
