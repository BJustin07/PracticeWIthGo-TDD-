package iteration

func RepeatWord(input string, timesToRepeat int) string {
	var repeat string
	for i := 0; i < timesToRepeat; i++ {
		repeat += input
	}
	return repeat
}
