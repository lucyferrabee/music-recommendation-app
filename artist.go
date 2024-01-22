package main

import (
	"log"

	"github.com/zmb3/spotify"
	"lucy.ferrabee.co.uk/auth"
)

type artist struct {
	id         string
	name       string
	popularity int
}

type ArtistService struct {
	Auth *auth.Authenticator
}

func NewArtistService(auth *auth.Authenticator) *ArtistService {
	return &ArtistService{
		Auth: auth,
	}
}

func (as *ArtistService) getByTrackId(id string) artist {
	client := as.Auth.Client

	trackID := spotify.ID(id)

	track, err := client.GetTrack(trackID)
	if err != nil {
		log.Fatalf("error retrieving track data: %v", err)
	}

	artistID := track.Artists[0].ID
	spotifyArtist, err := client.GetArtist(artistID)
	if err != nil {
		log.Fatalf("error retrieving artist data: %v", err)
	}

	return artist{
		name:       spotifyArtist.Name,
		id:         string(spotifyArtist.ID),
		popularity: spotifyArtist.Popularity,
	}
}
