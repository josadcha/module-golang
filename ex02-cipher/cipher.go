package cipher

import (
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct {
	c string
}

type Shift struct {
	s string
	x int
}

type Vigenere struct {
	s    string
	keys []int
}

func RemoveCharacters(input string, characters string) string {
	if input == "" {
		return input
	}
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.ToLower(strings.Map(filter, input))

}

func (c Caesar) Encode(s string) string {
	s = RemoveCharacters(s, "!-,.'?/^:;* 123456789")
	arr := make([]byte, len(s))
	for i := range arr {
		k := s[i]
		if k >= 120 && k <= 122 {
			k -= 23
		} else {
			k += 3
		}
		arr[i] = k
	}
	return string(arr)
}

func (c Caesar) Decode(d string) string {
	d = RemoveCharacters(d, "!-,.'?/^:;* 123456789")
	arr := make([]byte, len(d))
	for i := range arr {
		k := d[i]
		if k >= 97 && k <= 99 {
			k += 23
		} else {
			k -= 3
		}
		arr[i] = k
	}
	return string(arr)
}

func (sh Shift) Encode(s string) string {
	s = RemoveCharacters(s, "!-,.'?/^:;* 123456789")
	arr := make([]byte, len(s))
	for j := range arr {
		k := s[j]

		if sh.x < 0 {
			for i := sh.x; i < 0; i++ {
				k--
			}

			if k < 97 {
				k += 26
			}
		} else {
			for i := 0; i < sh.x; i++ {
				k++
			}
			if k > 122 {
				k -= 26
			}
		}

		arr[j] = k
	}

	return string(arr)
}

func (sh Shift) Decode(d string) string {
	d = RemoveCharacters(d, "!-,.'?/^:;* 123456789")
	arr := make([]byte, len(d))
	for j := range arr {
		k := d[j]

		if sh.x < 0 {
			for i := sh.x; i < 0; i++ {
				k++
			}

			if k > 122 {
				k -= 26
			}
		} else {
			for i := 0; i < sh.x; i++ {
				k--
			}

			if k < 97 {
				k += 26
			}
		}

		arr[j] = k
	}

	return string(arr)
}

func (v Vigenere) Encode(s string) string {
	s = RemoveCharacters(s, "!-,.'?/^:;* 123456789")
	arr := make([]byte, len(s))
	for j := range arr {
		k := s[j]
		var c Cipher
		c = Shift{string(k), v.keys[j]}
		new_k := c.Encode(string(k))
		for i := 0; i < len(new_k); i++ {
			arr[j] = new_k[i]
		}
	}
	return string(arr)

}

func (v Vigenere) Decode(d string) string {
	d = RemoveCharacters(d, "!-,.'?/^:;* 123456789")
	arr := make([]byte, len(d))

	for j := range arr {
		k := d[j]
		var c Cipher
		c = Shift{string(k), v.keys[j]}
		new_k := c.Decode(string(k))
		for i := 0; i < len(new_k); i++ {
			arr[j] = new_k[i]
		}
	}

	return string(arr)
}

func NewCaesar() Cipher {
	var c Cipher
	c = Caesar{""}
	return c
}

func NewShift(x int) Cipher {
	var c Cipher

	if x < -25 || x > 25 || x == 0 {
		return nil
	}

	c = Shift{"", x}

	return c
}

func NewVigenere(key string) Cipher {
	var c Cipher

	keys := make([]int, len(key))
	for i := 0; i < len(key); i++ {
		k := key[i]
		if k < 97 && k > 122 {
			return nil
		}

		keys[i] = int(k - 97)
	}

	cnt := 0
	for i := range keys {
		if keys[i] == 0 {
			cnt++
		}
	}
	if cnt == len(keys) {
		return nil
	}

	c = Vigenere{"", keys}

	return c
}
