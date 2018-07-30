package letter

func Frequency(str string) map[rune]int {
	fmap := make(map[rune]int)

	for _, v := range str {
		fmap[v] += 1
	}
	return fmap
}

func ConcurrentFrequency(str []string) map[rune]int {
	conf := make(map[rune]int)
	c := make(chan map[rune]int, len(str))

	for _, v := range str {
		go func(val string) {
			c <- Frequency(val)
		}(v)
	}

	for i := 0; i < len(str); i++ {
		for item, cnt := range <-c {
			conf[item] += cnt
		}
	}

	return conf
}
