package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/zmb3/spotify"
	"lucy.ferrabee.co.uk/auth"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// get id of track
	id, err := getInput("Input the id of the song and we'll generate a playlist for you: ", reader)
	if err != nil {
		log.Fatal(err)
	}

	// get access
	auth := auth.NewAuthenticator("4779b9533e004287b6536fd8c5325adf", "8dbf0a481f8b4de0901fbc9661f6036c")

	artistService := NewArtistService(auth)
	relatedArtistService := NewRelatedArtistService(auth)

	// get artist of track
	artist := artistService.getByTrackId(id)

	// get all top tracks of artist's related artists
	allTopTracks, err := relatedArtistService.getTopTracksFromRelatedArtists(artist.id, 1)
	if err != nil {
		log.Fatal(err)
	}

	// sort by popularity
	sort.Sort(byPopularity(allTopTracks))

	// Print the sorted tracks
	for _, track := range allTopTracks {
		fmt.Printf("Name: %s, Popularity: %d\n", track.Name, track.Popularity)
	}

	fmt.Println("Here's your playlist: ")

	// print the playlist
	fmt.Println(displayPlaylist(allTopTracks))
}

// byPopularity is a type that implements the sort.Interface for []spotify.FullTrack
type byPopularity []spotify.FullTrack

func (a byPopularity) Len() int           { return len(a) }
func (a byPopularity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPopularity) Less(i, j int) bool { return a[i].Popularity > a[j].Popularity }
