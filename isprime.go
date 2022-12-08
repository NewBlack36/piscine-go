package piscine

func IsPrime(nb int) bool {
	if nb > 0 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		if nb == 1 {
			return false
		}
		return true
	}
	return false
}
