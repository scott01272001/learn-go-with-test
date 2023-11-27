package iteration

func Repeat(char rune, repeat int) string {
	repeated := make([]rune, repeat)
	for i := range repeated {
		repeated[i] = char
	}
	return string(repeated)
}
