package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if row == 0 || row == height-1 || col == 0 || col == width-1 {
				fmt.Print(outer)
			} else {
				fmt.Print(inner)
			}
		}
		fmt.Println()
	}
}

// HINWEIS
// - Dies ist eine Verallgemeinerung der Aufgabe "Leeres Rechteck zeichnen".
// - Sie können genau so vorgehen wie in der einfachen Variante,
//   nur dass Sie die Zeichen für Rand und Füllung als Parameter verwenden.
// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
