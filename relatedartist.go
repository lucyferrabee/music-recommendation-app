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
	topTracks  []spotify.FullTrack
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
