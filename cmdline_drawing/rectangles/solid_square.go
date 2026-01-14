package rectangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein Quadrat mit dieser Seitenlänge auf der Konsole.
// Das Quadrat soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidSquare(length int) {
	for k := 0; k < length; k++ {

		for i := 0; i < length; i++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
