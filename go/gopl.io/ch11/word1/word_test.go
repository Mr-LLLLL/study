package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindromes("detartrated") {
		t.Error(`IsPalindromes("kayak") = false`)
	}
	if !IsPalindromes("kayak") {
		t.Error(`IsPalindromes("Palindrome") = true`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindromes("palindrome") {
		t.Error(`IsPalindromes("Palindrome") = true`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if IsPalindromes(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindromes(p) {
			t.Errorf("IsPalindromes(%q) = false", p)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromes("A man, a plan, a cannal: Panama")
	}
}

func benchmark(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		IsPalindromes("A man, a plan, a cannal: Panama")
	}
}

func Benchmark10(b *testing.B) {
	benchmark(b, 10)
}

func Benchmark100(b *testing.B) {
	benchmark(b, 100)
}

func Benchmark1000(b *testing.B) {
	benchmark(b, 1000)
}

func ExampleIsPalindromes() {
	fmt.Println(IsPalindromes("A man, a plan, acanal: Panama"))
	fmt.Println(IsPalindromes("palindrome"))
	// Output:
	// true
	// false
}
