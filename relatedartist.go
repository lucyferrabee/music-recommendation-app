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

// func (ras *RelatedArtistService) getByArtistId(id string) []relatedartist {
// 	client := ras.Auth.Client

// 	relatedArtists, err := client.GetRelatedArtists(spotify.ID(id))
// 	if err != nil {
// 		log.Fatalf("error retrieving related artist data: %v", err)
// 	}

// 	artists := make([]relatedartist, len(relatedArtists))

// 	for i, a := range relatedArtists {
// 		topTracks, err := client.GetArtistsTopTracks(a.ID, "US")
// 		if err != nil {
// 			log.Fatalf("error retrieving top tracks data: %v", err)
// 		}

// 		artists[i] = relatedartist{
// 			name:       a.Name,
// 			id:         string(a.ID),
// 			popularity: a.Popularity,
// 			topTracks:  topTracks,
// 		}
// 	}

// 	return artists
// }

func (ras *RelatedArtistService) getTopTracksFromRelatedArtists(id string, depth int) ([]spotify.FullTrack, error) {
	client := ras.Auth.Client

	// Fetch top tracks from the original artist
	topTracks, err := client.GetArtistsTopTracks(spotify.ID(id), "US")
	if err != nil {
		log.Fatalf("error retrieving top tracks data: %v", err)
		return nil, err
	}

	// Fetch related artists
	relatedArtists, err := client.GetRelatedArtists(spotify.ID(id))
	if err != nil {
		log.Fatalf("error retrieving related artist data: %v", err)
		return nil, err
	}

	// Fetch top tracks from related artists
	for _, artist := range relatedArtists {
		artistTopTracks, err := client.GetArtistsTopTracks(artist.ID, "US")
		if err != nil {
			log.Fatalf("error retrieving top tracks data for related artist %s: %v", artist.Name, err)
			return nil, err
		}
		topTracks = append(topTracks, artistTopTracks...)

		// Optionally, fetch top tracks from related artists' related artists recursively
		if depth > 1 {
			// Make a recursive call with reduced depth
			relatedArtistTopTracks, err := ras.getTopTracksFromRelatedArtists(artist.ID.String(), depth-1)
			if err != nil {
				return nil, err
			}
			topTracks = append(topTracks, relatedArtistTopTracks...)
		}
	}

	return topTracks, nil
}
