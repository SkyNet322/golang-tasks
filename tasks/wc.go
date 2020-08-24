package tasks

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	f := Fields(s)
	for _, i := range f {
		m[i]++
	}
	return m
}

func Fields(s string) []string {

	var stringSlice []string

	var runeSlice []rune

	for _, char := range s {

		if string(char) == " " {
			if len(runeSlice) > 0 {
				stringSlice = append(stringSlice, string(runeSlice))
				runeSlice = []rune{}
			}
			continue
		}

		runeSlice = append(runeSlice, char)

	}

	if len(runeSlice) > 0 {
		stringSlice = append(stringSlice, string(runeSlice))
		runeSlice = []rune{}
	}

	return stringSlice
}