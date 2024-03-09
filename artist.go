package main

import "github.com/zmb3/spotify"

type ArtistService struct {
	Client spotify.Client
}

func NewArtistService(client spotify.Client) *ArtistService {
	return &ArtistService{
		Client: client,
	}
}
