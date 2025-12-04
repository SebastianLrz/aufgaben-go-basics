package forms

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func PerimeterRightTriangle(a, b float64) float64 {
	return a + b + Hypotenuse(a, b)
}

// HINWEIS
// Der Umfang eines rechtwinkligen Dreiecks mit den Katheten a und b ist a + b + Hypotenuse(a, b).
