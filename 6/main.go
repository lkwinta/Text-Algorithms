package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcinCiura/AT_lab/6/nilsimsa"
)

var files = []string{
	"W1.txt",
	"W2.txt",
	"W3.txt",
	"W4.txt",
	"W5.txt",
	"W6.txt",
}

func main() {
	var texts []string
	for _, fn := range files {
		f, err := os.ReadFile(fn)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		texts = append(texts, string(f))
	}

	var fps []*nilsimsa.Fingerprint
	for _, t := range texts {
		fps = append(fps, nilsimsa.Nilsimsa(t))
	}

	for _, f := range fps[1:] {
		fmt.Printf("%d ", nilsimsa.HammingDistance(fps[0], f))
	}
}
