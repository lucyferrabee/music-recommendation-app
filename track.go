package main

import (
	"log"

	"github.com/zmb3/spotify"
	"lucy.ferrabee.co.uk/auth"
)

type track struct {
	id         string
	name       string
	popularity int
}

type TrackService struct {
	Auth *auth.Authenticator
}

func NewTrackService(auth *auth.Authenticator) *TrackService {
	return &TrackService{
		Auth: auth,
	}
}

func (ts *TrackService) getById(id string) track {
	client := ts.Auth.Client

	trackID := spotify.ID(id)
	spotifyTrack, err := client.GetTrack(trackID)
	if err != nil {
		log.Fatalf("error retrieving track data: %v", err)
	}

	return track{
		name:       spotifyTrack.Name,
		id:         string(spotifyTrack.ID),
		popularity: spotifyTrack.Popularity,
	}
}
