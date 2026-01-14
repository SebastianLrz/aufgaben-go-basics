package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	for k := 0; k < height; k++ {

		if k == 0 {

			for i := 0; i < width; i++ {
				fmt.Print(outer)
			}
			fmt.Println()
		} else if k == height-1 {
			for i := 0; i < width; i++ {
				fmt.Print(outer)
			}
		} else {
			fmt.Print(outer)
			for i := 0; i < width-2; i++ {
				fmt.Print(inner)
			}
			fmt.Print(outer)
			fmt.Println()
		}

	}
	fmt.Println()
}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
