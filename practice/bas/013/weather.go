package practice

// Weather returns a label for the given temperature.
func Weather(t int) string {
	if t < 20 {
		return "too cold"
	} else if t <= 30 {
		return "perfect"
	} else {
		return "so hot"
	}
}
