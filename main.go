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

func generatePlaylist() (track, artist, []relatedartist) {
	reader := bufio.NewReader(os.Stdin)

	id, _ := getInput("Input the id of the song and we'll generate a playlist for you: ", reader)

	track := getById(id)
	artist := getByTrackId(id)
	relatedArtists := getByArtistId(artist.id)

	fmt.Println("Here's your playlist: ")

	for _, relatedartist := range relatedArtists {
		fmt.Println(relatedartist.topTracks[0], relatedartist.name)
	}

	return track, artist, relatedArtists
}

func main() {

	generatePlaylist()
}
