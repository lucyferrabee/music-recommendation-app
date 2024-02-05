package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zmb3/spotify"
	"lucy.ferrabee.co.uk/auth"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	id, err := getInput("Input the id of the song and we'll generate a playlist for you: ", reader)
	if err != nil {
		log.Fatal(err)
	}

	auth := auth.NewAuthenticator("4779b9533e004287b6536fd8c5325adf", "8dbf0a481f8b4de0901fbc9661f6036c")

	relatedArtistService := NewRelatedArtistService(auth)
	playlistService := NewPlaylistService(relatedArtistService)

	const targetPopularity = 70
	const threshold = 2

	playlistTracks, err := playlistService.GeneratePlaylist(id, targetPopularity, threshold)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Here's your playlist: ")

	for _, track := range playlistTracks {
		fmt.Printf("Name: %s, Artist: %s, Popularity: %d\n", track.Name, track.Artists[0].Name, track.Popularity)
	}
}

type byPopularity []spotify.FullTrack

func (a byPopularity) Len() int           { return len(a) }
func (a byPopularity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPopularity) Less(i, j int) bool { return a[i].Popularity > a[j].Popularity }

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
