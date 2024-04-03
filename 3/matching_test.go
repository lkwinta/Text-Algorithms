package matching

import (
	"os"
	"slices"
	"testing"
)

func TestPreprocess(t *testing.T) {
	data := []string{
		"aaaaaaa",
		"pies",
		"dźwiedź",
		"owocowo",
		"indianin",
		"nienapełnienie",
	}
	for _, in := range data {
		got := Preprocess([]byte(in))
		want := SimplePreprocess([]byte(in))
		if !slices.Equal(got, want) {
			t.Errorf(`Preprocess(%#v) == %#v want %#v`,
				in, got, want)
		}
	}
}

func TestNaive(t *testing.T) {
	pat := []byte("abac")
	text := []byte("asaaddabacasda")
	got := []int{}
	Naive(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		// Zgłoś błąd, korzystając z funkcji `t.Errorf`
		t.Errorf(`Naive(%#v, %#v) == %#v want %#v`, text, pat, got, want)
	}
}

func TestBackwardNaive(t *testing.T) {
	pat := []byte("abac")
	text := []byte("asaaddabacasda")
	got := []int{}
	BackwardNaive(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		// Zgłoś błąd, korzystając z funkcji `t.Errorf`
		t.Errorf(`BackwardNaive(%#v, %#v) == %#v want %#v`, text, pat, got, want)
	}
}

func TestBoyerMoore(t *testing.T) {
	pat := []byte("abac")
	text := []byte("asaaddabacasda")
	got := []int{}
	BoyerMoore(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		// Zgłoś błąd, korzystając z funkcji `t.Errorf`
		t.Errorf(`BoyerMoore(%#v, %#v) == %#v want %#v`, text, pat, got, want)
	}
}

func TestKMP(t *testing.T) {
	pat := []byte("abac")
	text := []byte("asaaddabacasda")
	got := []int{}
	KMP(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		// Zgłoś błąd, korzystając z funkcji `t.Errorf`
		t.Errorf(`KMP(%#v, %#v) == %#v want %#v`, text, pat, got, want)
	}
}

func TestKarpRabin(t *testing.T) {
	pat := []byte("abac")
	text := []byte("asaaddabacasda")
	got := []int{}
	KarpRabin(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		// Zgłoś błąd, korzystając z funkcji `t.Errorf`
		t.Errorf(`KarpRabin(%#v, %#v) == %#v want %#v`, text, pat, got, want)
	}
}

func TestShiftOr(t *testing.T) {
	pat := []byte("abac")
	text := []byte("asaaddabacasda")
	got := []int{}
	ShiftOr(pat, text, func(n int) { got = append(got, n) })
	want := indices(pat, text)
	if !slices.Equal(got, want) {
		// Zgłoś błąd, korzystając z funkcji `t.Errorf`
		t.Errorf(`ShiftOr(%#v, %#v) == %#v want %#v`, text, pat, got, want)
	}

}

var short_word = []byte("drogą")
var long_word = []byte("królewskich")

func BenchmarkNaiveShort(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = short_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Naive(pat, text, func(int) {})
	}
}

func BenchmarkNaiveLong(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = long_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Naive(pat, text, func(int) {})
	}
}

func BenchmarkBackwardNaiveShort(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = short_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BackwardNaive(pat, text, func(int) {})
	}
}

func BenchmarkBackwardNaiveLong(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = long_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BackwardNaive(pat, text, func(int) {})
	}
}

func BenchmarkBoyerMooreShort(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = short_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMoore(pat, text, func(int) {})
	}
}

func BenchmarkBoyerMooreLong(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = long_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BoyerMoore(pat, text, func(int) {})
	}
}

func BenchmarkKMPShort(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = short_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMP(pat, text, func(int) {})
	}
}

func BenchmarkKMPLong(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = long_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KMP(pat, text, func(int) {})
	}
}

func BenchmarkKarpRabinShort(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = short_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KarpRabin(pat, text, func(int) {})
	}
}

func BenchmarkKarpRabinLong(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = long_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KarpRabin(pat, text, func(int) {})
	}
}

func BenchmarkShiftOrShort(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = short_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftOr(pat, text, func(int) {})
	}
}

func BenchmarkShiftOrLong(b *testing.B) {
	text, err := os.ReadFile("kordian.txt")
	if err != nil {
		b.Fatal(err)
	}

	var pat = long_word
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShiftOr(pat, text, func(int) {})
	}
}

func indices(pat, text []byte) []int {
	r := []int{}
	for i := 0; i+len(pat) <= len(text); i++ {
		if slices.Equal(text[i:i+len(pat)], pat) {
			r = append(r, i)
		}
	}
	return r
}
