package main

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type track struct {
	id         string
	name       string
	popularity int
}

func getById(id string) *spotify.FullTrack {

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

	return track
}

func createObject(spotifyTrack *spotify.FullTrack) track {
	t := track{
		name:       spotifyTrack.Name,
		id:         string(spotifyTrack.ID),
		popularity: spotifyTrack.Popularity,
	}

	return t
}
