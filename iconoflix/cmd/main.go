package main

import (
	"flag"
	"fmt"

	"github.com/derailed/pkg/iconoflix"
)

func main() {
	v := flag.String("v", "v1", "Specify a movie sample version [v1|v2]")
	flag.Parse()

	corpus, err := iconoflix.LoadMem(*v)
	if err != nil {
		panic(err)
	}

	fmt.Println("----------------------------------------")
	fmt.Printf("Iconoflix Movie Database [%s]\n", *v)
	fmt.Println("----------------------------------------")
	for i, m := range corpus.Movies {
		fmt.Printf("%2d %s\n", i, m.Name)
		for j, e := range m.Icons {
			fmt.Printf("  %2d %s\n", j, e.Emoji)
		}
	}
}
