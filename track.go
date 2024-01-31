package main

import (
	"lucy.ferrabee.co.uk/auth"
)

type TrackService struct {
	Auth *auth.Authenticator
}

func NewTrackService(auth *auth.Authenticator) *TrackService {
	return &TrackService{
		Auth: auth,
	}
}
