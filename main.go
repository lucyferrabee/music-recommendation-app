package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	id, err := getInput("Input the id of the song and we'll generate a playlist for you: ", reader)
	if err != nil {
		log.Fatal(err)
	}

	artist := getByTrackId(id)
	relatedArtists := getByArtistId(artist.id)

	fmt.Println("Here's your playlist: ")

	playlist := generateFromRelatedArtists(relatedArtists)
	fmt.Println(displayPlaylist(playlist))
}
