package main

import (
	"math"
	"sort"

	"github.com/zmb3/spotify"
)

type PlaylistService struct {
	RelatedArtistService *RelatedArtistService
}

func NewPlaylistService(relatedArtistService *RelatedArtistService) *PlaylistService {
	return &PlaylistService{
		RelatedArtistService: relatedArtistService,
	}
}

func (ps *PlaylistService) GeneratePlaylist(trackID string, targetPopularity, threshold int) ([]spotify.FullTrack, error) {
	artist := ps.RelatedArtistService.getFirstRelatedArtistByTrackId(trackID)
	allTopTracks, err := ps.RelatedArtistService.getTopTracksFromRelatedArtists(artist.id, 1)
	if err != nil {
		return nil, err
	}

	sort.Sort(byPopularity(allTopTracks))

	similarTracks := ps.chooseSimilarPopularity(allTopTracks, targetPopularity, threshold)
	uniqueTracks := ps.removeDuplicates(similarTracks)

	return uniqueTracks, nil
}

func (ps *PlaylistService) chooseSimilarPopularity(tracks []spotify.FullTrack, targetPopularity, threshold int) []spotify.FullTrack {
	var selectedTracks []spotify.FullTrack

	if len(tracks) == 0 {
		return selectedTracks
	}

	for _, track := range tracks {
		if math.Abs(float64(track.Popularity)-float64(targetPopularity)) <= math.Abs(float64(threshold)) {
			selectedTracks = append(selectedTracks, track)
		}
	}

	return selectedTracks
}

func (ps *PlaylistService) removeDuplicates(tracks []spotify.FullTrack) []spotify.FullTrack {
	uniqueTrackIDs := make(map[spotify.ID]struct{})
	var uniqueTracks []spotify.FullTrack

	for _, track := range tracks {
		if _, exists := uniqueTrackIDs[track.ID]; !exists {
			uniqueTrackIDs[track.ID] = struct{}{}
			uniqueTracks = append(uniqueTracks, track)
		}
	}

	return uniqueTracks
}

type byPopularity []spotify.FullTrack

func (a byPopularity) Len() int           { return len(a) }
func (a byPopularity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPopularity) Less(i, j int) bool { return a[i].Popularity < a[j].Popularity }
