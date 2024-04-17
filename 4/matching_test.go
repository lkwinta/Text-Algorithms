package matching

import (
	"fmt"
	"os"
	"slices"
	"testing"

	ahocorasick "github.com/BobuSumisu/aho-corasick"
)

var words = []string{
	"godzina",
	"godziny",
	"godzinie",
	"godzinę",
	"godziną",
	"godzino",
}

func TestFuzzyShiftOrH(t *testing.T) {
	text, error := os.ReadFile("cierpienia-mlodego-wertera.txt")

	if error != nil {
		t.Error("Couldn't read file")
	}

	pat := []byte("godzina")
	got := []string{}

	FuzzyShiftOrH(pat, text,
		func(n int) { got = append(got, string(text[n:n+len(pat)])) })

	fmt.Printf("%v \n", got)
}

func TestFuzzyShiftOrL(t *testing.T) {
	text, error := os.ReadFile("cierpienia-mlodego-wertera.txt")

	if error != nil {
		t.Error("Couldn't read file")
	}

	pat := []byte("godzina")
	got := []string{}

	FuzzyShiftOrL(pat, text,
		func(n int) { got = append(got, string(text[n-len(pat)-1:n+1])) })

	fmt.Printf("%v \n", got)
}

func TestAhoCorasick(t *testing.T) {
	text, error := os.ReadFile("cierpienia-mlodego-wertera.txt")

	if error != nil {
		t.Error("Couldn't read file")
	}

	builder := ahocorasick.NewTrieBuilder()
	builder.AddStrings(words)
	trie := builder.Build()

	matches := trie.MatchString(string(text))

	var got [6][]int64
	for _, m := range matches {
		got[m.Pattern()] = append(got[m.Pattern()], m.Pos())
	}

	var want [6][]int64
	for i, pat := range words {
		BoyerMoore([]byte(pat), text, func(n int) {
			want[i] = append(want[i], int64(n))
		})
	}

	for i := range got {
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("got[%d] == %v want %v", i, got[i], want[i])
		}
	}

	fmt.Printf("Znalezione wystąpienia wyrazów %v:\n%v\n", words, got)
}

func BenchmarkAhoCorasick(b *testing.B) {
	text, error := os.ReadFile("cierpienia-mlodego-wertera.txt")

	if error != nil {
		b.Error("Couldn't read file")
	}

	stext := string(text)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		builder := ahocorasick.NewTrieBuilder()
		builder.AddStrings(words)
		trie := builder.Build()
		trie.MatchString(stext)
	}
}

func BenchmarkBoyerMoore(b *testing.B) {
	text, error := os.ReadFile("cierpienia-mlodego-wertera.txt")

	if error != nil {
		b.Error("Couldn't read file")
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, pat := range words {
			BoyerMoore([]byte(pat), text, func(int) {})
		}
	}
}
