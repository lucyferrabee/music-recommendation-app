package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/zmb3/spotify"
	"lucy.ferrabee.co.uk/auth"
)

type Config struct {
	TargetPopularity int `json:"target_popularity"`
	Threshold        int `json:"threshold"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	id, err := getInput("Input the id of the song and we'll generate a playlist for you: ", reader)
	if err != nil {
		log.Fatal(err)
	}

	auth := auth.NewSpotifyClient("4779b9533e004287b6536fd8c5325adf", "8dbf0a481f8b4de0901fbc9661f6036c")

	relatedArtistService := NewRelatedArtistService(auth)
	playlistService := NewPlaylistService(relatedArtistService)

	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	playlistTracks, err := generatePlaylist(playlistService, id, config)
	if err != nil {
		log.Fatal(err)
	}

	printPlaylist(playlistTracks)
}

func loadConfig(filename string) (Config, error) {
	var config Config

	configFile, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}

func generatePlaylist(service *PlaylistService, id string, config Config) ([]spotify.FullTrack, error) {
	return service.GeneratePlaylist(id, config.TargetPopularity, config.Threshold)
}

func printPlaylist(tracks []spotify.FullTrack) {
	fmt.Println("Here's your playlist: ")

	sort.Sort(byPopularity(tracks))

	for _, track := range tracks {
		fmt.Printf("Name: %s, Artist: %s, Popularity: %d\n", track.Name, track.Artists[0].Name, track.Popularity)
	}
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}
