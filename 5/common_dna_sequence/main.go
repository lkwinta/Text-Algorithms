package main

import (
	"fmt"
	"log"
	"os"

	suffixtree "github.com/MarcinCiura/AT_lab/5/suffixtree"
)

var files = []string{
	"human.txt",
	"lis_polarny.txt",
	"norka.txt",
	"szop_pracz.txt",
}

func lensIndex(lens []int, suffixStart int) uint {
	for i, len := range lens {
		if suffixStart < len {
			return uint(i)
		}
	}

	panic("error")
}

func dfs1(index *suffixtree.Index, node int, bits []uint, lens []int) uint {
	if index.IsLeaf(node) {
		return uint(1) << lensIndex(lens, index.SuffixStart(node))
	}

	for _, n := range index.Edges(node) {
		bits[node] |= dfs1(index, n, bits, lens)
	}

	return bits[node]
}

func dfs2(index *suffixtree.Index, node int, bits []uint, endmask uint, path string) {
	if bits[node] != endmask {
		return
	}

	if node != index.Root() {
		path += index.EdgeLabel(node)
	}

	for _, n := range index.Edges(node) {
		fmt.Printf("%v %v \n", len(path), path)
		dfs2(index, n, bits, endmask, path)
	}
}

func main() {
	text := []byte{}
	lens := []int{}
	for i, fn := range files {
		f, err := os.ReadFile(fn)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		text = append(text, f...)
		text = append(text, byte('0'+i))
		lens = append(lens, len(text))
	}

	index := suffixtree.New(text)
	bits := make([]uint, index.NumNodes())
	dfs1(index, index.Root(), bits, lens)
	dfs2(index, index.Root(), bits, 1<<len(lens)-1, "")
}
