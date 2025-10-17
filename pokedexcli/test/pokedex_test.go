package test

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	pokedexcli "go-tutorial/bootdotdev/pokedexcli/pokedex"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	// This is the main testing function for pokedex file
	cases := []struct {
		input string
		expected []string
	}{
		{
			input:	"  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:	"Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander","bulbasaur","pikachu"},
		},
		{
			input:	"  hello   ",
			expected: []string{"hello"},
		},
		{
			input:	"",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := pokedexcli.CleanInput(c.input)
		//Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error mesage
		// and fail the test
		if (len(actual.CMD) + len(actual.ExtraCMD)) != len(c.expected) {
			t.Errorf("clean input length != expected length, %v != %v", actual, c.expected)
		}
		for i := range actual.ExtraCMD {
			word := actual.ExtraCMD[i]
			expectedWord := c.expected[i]
			// Check each word in the slice 
			// if they don't match, use t.Errorf to print an error mesage
			// and fail the test
			if word != expectedWord {
				t.Errorf("input: %s\nexpected: %s", word, expectedWord)
			}
		}
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := *internal.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := *internal.NewCache(baseTime)

	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}