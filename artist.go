package main

import (
	"lucy.ferrabee.co.uk/auth"
)

type ArtistService struct {
	Auth *auth.Authenticator
}

func NewArtistService(auth *auth.Authenticator) *ArtistService {
	return &ArtistService{
		Auth: auth,
	}
}
