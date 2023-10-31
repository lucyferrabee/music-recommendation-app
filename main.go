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

func getTrack() track {
	reader := bufio.NewReader(os.Stdin)

	id, _ := getInput("Input the id of the song you'd like information on: ", reader)

	sp := getById(id)
	t := createObject(sp)
	fmt.Println("The name of this song is: ", t.name, "the popularity of the song is: ", t.popularity)

	return t
}

func main() {

	getTrack()
}
