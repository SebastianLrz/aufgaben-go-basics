package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

// HINWEIS
// - Nutzen Sie zwei ineinander geschachtelte Schleifen, um die Zeilen und Spalten des Rechtecks zu durchlaufen.
// - Die äußere Schleife sollte die Zeilen des Rechtecks durchlaufen, die innere Schleife die Spalten.
