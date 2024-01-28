package main

import (
	"fmt"

	"github.com/zmb3/spotify"
)

type playlist struct {
	tracks []spotify.FullTrack
}

func displayPlaylist(p []spotify.FullTrack) string {

	var playlist (string)

	for _, pl := range p {
		playlist += fmt.Sprintf("%s by %s (ID: %s)\n", pl.Name, pl.Artists[0].Name, pl.ID)
	}

	return playlist
}
