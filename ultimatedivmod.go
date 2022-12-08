package piscine

func UltimateDivMod(a *int, b *int) {
	v := *a
	*a = *a / *b
	*b = v % *b
}
