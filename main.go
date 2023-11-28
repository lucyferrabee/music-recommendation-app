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

func getTrack() (track, artist) {
	reader := bufio.NewReader(os.Stdin)

	id, _ := getInput("Input the id of the song you'd like information on: ", reader)

	sp := getById(id)
	t := createTrackObject(sp)
	spa := getByTrackId(id)
	a := createArtistObject(spa)
	fmt.Println("The name of this song is: ", t.name, "the popularity of the song is: ", t.popularity, "The name of the artist is: ", a.name, "The popularity of the artist is: ", a.popularity)

	return t, a
}

func main() {

	getTrack()
}
