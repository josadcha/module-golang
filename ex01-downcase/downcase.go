package downcase

func Downcase(x string) (string, error) {
	arr := make([]byte, len(x))
	for i := 0; i < len(arr); i++ {
		k := x[i]
		if k >= 'A' && k <= 'Z' {
			k += 'a' - 'A'
		}
		arr[i] = k
	}
	return string(arr), nil
}
