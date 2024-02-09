package iteration

func Repeat(char string, iterationTimes int) string {
	str := char
	for i := 0; i < iterationTimes; i++ {
		str += char
	}
	return str
}
