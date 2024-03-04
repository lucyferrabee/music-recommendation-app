package auth

import (
	"context"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func NewSpotifyClient(clientID, clientSecret string) spotify.Client {
	authConfig := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}

	httpClient := authConfig.Client(context.Background())

	client := spotify.NewClient(httpClient)

	return client
}
