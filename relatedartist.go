package main

import (
	"log"

	"github.com/zmb3/spotify"
)

type relatedartist struct {
	id         string
	name       string
	popularity int
	topTracks  []spotify.FullTrack
}

type RelatedArtistService struct {
	Client spotify.Client // Use spotify.Client directly
}

func NewRelatedArtistService(client spotify.Client) *RelatedArtistService {
	return &RelatedArtistService{
		Client: client,
	}
}

func (ras *RelatedArtistService) getTopTracksFromRelatedArtists(id string, depth int) ([]spotify.FullTrack, error) {
	client := ras.Client

	topTracks, err := client.GetArtistsTopTracks(spotify.ID(id), "US")
	if err != nil {
		log.Fatalf("error retrieving top tracks data: %v", err)
		return nil, err
	}

	relatedArtists, err := client.GetRelatedArtists(spotify.ID(id))
	if err != nil {
		log.Fatalf("error retrieving related artist data: %v", err)
		return nil, err
	}

	for _, artist := range relatedArtists {
		artistTopTracks, err := client.GetArtistsTopTracks(artist.ID, "US")
		if err != nil {
			log.Fatalf("error retrieving top tracks data for related artist %s: %v", artist.Name, err)
			return nil, err
		}
		topTracks = append(topTracks, artistTopTracks...)

		if depth > 1 {
			relatedArtistTopTracks, err := ras.getTopTracksFromRelatedArtists(artist.ID.String(), depth-1)
			if err != nil {
				return nil, err
			}
			topTracks = append(topTracks, relatedArtistTopTracks...)
		}
	}

	return topTracks, nil
}

func (ras *RelatedArtistService) getFirstRelatedArtistByTrackId(id string) relatedartist {
	client := ras.Client

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

	topTracks, err := client.GetArtistsTopTracks(artistID, "US")
	if err != nil {
		log.Fatalf("error retrieving top tracks data: %v", err)
	}

	return relatedartist{
		name:       artist.Name,
		id:         string(artist.ID),
		popularity: artist.Popularity,
		topTracks:  topTracks,
	}
}
