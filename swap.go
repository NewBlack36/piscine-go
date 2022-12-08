package piscine

func Swap(a *int, b *int) {
	v := *a
	*a = *b
	*b = v
}
