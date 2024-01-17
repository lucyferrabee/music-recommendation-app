package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

	id, err := getInput("Input the id of the song and we'll generate a playlist for you: ", reader)
	if err != nil {
		log.Fatal(err)
	}

	auth := auth.NewAuthenticator("4779b9533e004287b6536fd8c5325adf", "8dbf0a481f8b4de0901fbc9661f6036c")

	artistService := NewArtistService(auth)
	relatedArtistService := NewRelatedArtistService(auth)

	artist := artistService.getByTrackId(id)
	relatedArtists := relatedArtistService.getByArtistId(artist.id)

	fmt.Println("Here's your playlist: ")

	playlist := generateFromRelatedArtists(relatedArtists)
	fmt.Println(displayPlaylist(playlist))
}
