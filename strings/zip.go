package strings

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	result := ""
	min := len(s1)
	if len(s2) < min {
		min = len(s2)
	}
	for i := 0; i < min; i++ {
		result += string(s1[i]) + string(s2[i])
	}
	result += s1[min:]
	result += s2[min:]
	return result
}

// HINWEIS
// Laufen Sie in einer for-Schleife über alle Buchstaben des kürzeren Strings.
// Fügen Sie in jedem Schleifendurchlauf den aktuellen Buchstaben aus s1 und s2
// an den result-String an.
// Am Ende der Schleife müssen Sie noch die restlichen Buchstaben des längeren
// Strings an den result-String anhängen.
