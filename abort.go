package piscine

func Abort(a, b, c, d, e int) int {
	var tri []int
	tri = append(tri, a)
	tri = append(tri, b)
	tri = append(tri, c)
	tri = append(tri, d)
	tri = append(tri, e)
	i := 1
	for i < len(tri) {
		if tri[i-1] > tri[i] {
			t := tri[i]
			tri[i] = tri[i-1]
			tri[i-1] = t
			i = 1
		} else {
			i++
		}
	}
	return tri[2]
}
