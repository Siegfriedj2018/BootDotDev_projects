package main

import (
	"log"
	"strings"
)

func findReplaceBadWords(badWords map[string]bool, chirpText string) (string, error) {
	chirpWords := strings.Fields(chirpText)

	for idx, word := range chirpWords {
		if _, ok := badWords[strings.ToLower(word)]; ok {
			log.Printf("Found bad word: %s, censoring now...", word)
			badWords[word] = ok
			chirpWords[idx] = "****"
		}
	}

	return strings.Join(chirpWords, " "), nil
}