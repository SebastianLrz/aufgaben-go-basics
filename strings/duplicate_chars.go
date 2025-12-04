package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	result := ""
	for _, char := range s {
		result += string(char)
		result += string(char)
	}
	return result
}

// HINWEIS
// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
// Fügen Sie in jedem Schleifendurchlauf den aktuellen Buchstaben zweimal
// an den result-String an.
