package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // List of rock and roll lyrics
    rockAndRollLyrics := []string{
        "I see a red door and I want it painted black",
        "Born to be wild",
        "We don't need no education",
        "Sweet child o' mine",
        "I wanna rock and roll all night",
        // Add more lyrics here
    }

    // Get a random index
    randomIndex := rand.Intn(len(rockAndRollLyrics))

    // Print the random rock and roll lyric
    fmt.Println(rockAndRollLyrics[randomIndex])
}
