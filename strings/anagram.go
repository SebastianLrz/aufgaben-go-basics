package strings

import (
	"strings"
)

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	for _, c := range s1 {
		if CountChar(s1, c) != CountChar(s2, c) {
			return false
		}
	}
	return true
}

// HINWEIS
// Lauft in einer for-Schleife über alle Buchstaben von s1.
// Prüfen Sie in jedem Schleifendurchlauf mittels der Funktion CountChar(),
// ob der aktuelle Buchstabe in s1 gleich oft vorkommt wie in s2.
// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	return IsAnagram(strings.ToLower(s1), strings.ToLower(s2))
}

// HINWEIS
// Verwenden Sie die Funktion strings.ToLower(), um s1 und s2 in Kleinbuchstaben
// umzuwandeln, bevor Sie die Funktion IsAnagram() aufrufen.
