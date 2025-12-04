package forms

// Erwartet eine Seitenlänge a.
// Liefert die Fläche des entsprechenden Quadrats.
func AreaSquare(a float64) float64 {
	return AreaRectangle(a, a)
}

// HINWEIS
// Benutzen Sie die Funktion AreaRectangle().
