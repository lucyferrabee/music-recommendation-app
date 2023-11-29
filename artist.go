package main

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type artist struct {
	id         string
	name       string
	popularity int
}

func getByTrackId(id string) artist {

	authConfig := &clientcredentials.Config{
		ClientID:     "4779b9533e004287b6536fd8c5325adf",
		ClientSecret: "8dbf0a481f8b4de0901fbc9661f6036c",
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieving access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)

	trackID := spotify.ID(id)
	track, err := client.GetTrack(trackID)
	if err != nil {
		log.Fatalf("error retrieving track data: %v", err)
	}

	artistID := track.Artists[0].ID
	artist, err := client.GetArtist(artistID)
	if err != nil {
		log.Fatalf("error retrieving artist data: %v", err)
	}

	a := createArtistObject(artist)

	return a
}

func createArtistObject(spotifyArtist *spotify.FullArtist) artist {

	a := artist{
		name:       spotifyArtist.Name,
		id:         string(spotifyArtist.ID),
		popularity: spotifyArtist.Popularity,
	}

	return a
}
