package piscine

func ConcatParams(args []string) string {
	ping := ""

	for index, val := range args {
		ping += string(val)
		if index != len(args)-1 {
			ping += "\n"
		}
	}
	return ping
}
