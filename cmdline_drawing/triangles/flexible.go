package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {
	for k := 1; k <= length; k++ {

		if k == 1 {
			fmt.Println(outer)

		} else if k == length {
			for i := 0; i < length; i++ {
				fmt.Print(outer)
			}
			fmt.Println()

		} else {
			fmt.Print(outer)
			for i := 1; i < k-1; i++ {
				fmt.Print(inner)
			}
			fmt.Print(outer)
			fmt.Println()
		}

	}

}
