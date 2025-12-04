package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

// HINWEIS
// - Verwenden Sie eine Hilfsvariable, um das bisher kleinste Element zu speichern.
//   Diese Variable können Sie mit dem ersten Element der Liste initialisieren.
// - Nutzen Sie dann eine Schleife, um die Liste zu durchlaufen.
//   In jedem Schleifendurchlauf prüfen Sie, ob das aktuelle Element kleiner als das bisher kleinste ist.
//   Wenn ja, aktualisieren Sie die Hilfsvariable.
