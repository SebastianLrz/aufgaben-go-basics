package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	for k := 0; k < height; k++ {

		if k == 0 {

			for i := 0; i < width; i++ {
				fmt.Print("#")
			}
			fmt.Println()
		} else if k == height-1 {
			for i := 0; i < width; i++ {
				fmt.Print("#")
			}
		} else {
			fmt.Print("#")
			for i := 0; i < width-2; i++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
			fmt.Println()
		}

	}
	fmt.Println()
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
