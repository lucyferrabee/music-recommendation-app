package main

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type relatedartist struct {
	id         string
	name       string
	popularity int
}

func getByArtistId(id string) []relatedartist {

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

	relatedArtists, err := client.GetRelatedArtists(spotify.ID(id))
	if err != nil {
		log.Fatalf("error retrieving related artist data: %v", err)
	}

	artists := []relatedartist{}

	for _, a := range relatedArtists {

		artists = append(artists, createRelatedArtistObject(a))
	}

	return artists
}

func createRelatedArtistObject(spotifyArtist spotify.FullArtist) relatedartist {

	a := relatedartist{
		name:       spotifyArtist.Name,
		id:         string(spotifyArtist.ID),
		popularity: spotifyArtist.Popularity,
	}

	return a
}
