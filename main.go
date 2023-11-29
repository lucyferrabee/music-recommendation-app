package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func getTrackAndArtist() (track, artist, []relatedartist) {
	reader := bufio.NewReader(os.Stdin)

	id, _ := getInput("Input the id of the song you'd like information on: ", reader)

	track := getById(id)
	artist := getByTrackId(id)
	relatedArtists := getByArtistId(artist.id)
	fmt.Println("The name of this song is: ", track.name,
		"The popularity of the song is: ", track.popularity,
		"The name of the artist is: ", artist.name,
		"The popularity of the artist is: ", artist.popularity,
		"The first related artist is: ", relatedArtists[0].name,
	)

	return track, artist, relatedArtists
}

func main() {

	getTrackAndArtist()
}
