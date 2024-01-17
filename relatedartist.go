package main

import (
	"log"

	"github.com/zmb3/spotify"
	"lucy.ferrabee.co.uk/auth"
)

type relatedartist struct {
	id         string
	name       string
	popularity int
	topTracks  []spotify.FullTrack
}

type RelatedArtistService struct {
	Auth *auth.Authenticator
}

func NewRelatedArtistService(auth *auth.Authenticator) *RelatedArtistService {
	return &RelatedArtistService{
		Auth: auth,
	}
}

func (ras *RelatedArtistService) getByArtistId(id string) []relatedartist {
	client := ras.Auth.Client

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
