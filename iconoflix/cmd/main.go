package main

import (
	"fmt"

	"github.com/derailed/nuclio/movie"
)

func main() {
	fmt.Println(movie.LoadMem())
}
