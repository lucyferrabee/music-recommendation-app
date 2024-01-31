package auth

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type Authenticator struct {
	Client spotify.Client
}

func NewAuthenticator(clientID, clientSecret string) *Authenticator {
	authConfig := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieving access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)

	return &Authenticator{
		Client: client,
	}
}
