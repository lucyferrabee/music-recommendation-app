package main

import "github.com/zmb3/spotify"

type TrackService struct {
	Client spotify.Client // Use spotify.Client directly
}

func NewTrackService(client spotify.Client) *TrackService {
	return &TrackService{
		Client: client,
	}
}
