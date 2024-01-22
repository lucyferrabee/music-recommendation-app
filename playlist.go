package main

import (
	"fmt"

	"github.com/zmb3/spotify"
)

type playlist struct {
	tracks []spotify.FullTrack
}

func generateFromRelatedArtists(ra []relatedartist) playlist {

	var tracks []spotify.FullTrack

	for _, relatedartist := range ra {
		topTrack := relatedartist.topTracks[0]

		tracks = append(tracks, topTrack)
	}

	return playlist{
		tracks: tracks,
	}
}

func displayPlaylist(p playlist) string {

	var playlist (string)

	for _, pl := range p.tracks {
		playlist += fmt.Sprintf("%s by %s (ID: %s)\n", pl.Name, pl.Artists[0].Name, pl.ID)
	}

	return playlist
}
