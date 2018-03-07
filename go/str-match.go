package main

import (
	"errors"
	"fmt"
)

type KMP struct {
	pattern string
	prefix  []int
	size    int
}

func NewKMP(pattern string) (*KMP, error) {
	prefix, err := computePrefix(pattern)
	if err != nil {
		return nil, err
	}
	return &KMP{
			pattern: pattern,
			prefix:  prefix,
			size:    len(pattern)},
		nil
}

// returns an array containing indexes of matches
// - error if pattern argument is less than 1 char
func computePrefix(pattern string) ([]int, error) {
	len_p := len(pattern)
	if len_p < 2 {
		if len_p == 0 {
			return nil, errors.New("'pattern' must contain at least one character")
		}
		return []int{-1}, nil
	}
	t := make([]int, len_p)
	t[0], t[1] = -1, 0

	pos, count := 2, 0
	for pos < len_p {

		if pattern[pos-1] == pattern[count] {
			count++
			t[pos] = count
			pos++
		} else {
			if count > 0 {
				count = t[count]
			} else {
				t[pos] = 0
				pos++
			}
		}
	}
	return t, nil
}

// return index of first occurence of kmp.pattern in argument 's'
// - if not found, returns -1
func (kmp *KMP) FindStringIndex(s string) int {
	// sanity check
	if len(s) < kmp.size {
		return -1
	}
	m, i := 0, 0
	for m+i < len(s) {
		if kmp.pattern[i] == s[m+i] {
			if i == kmp.size-1 {
				return m
			}
			i++
		} else {
			m = m + i - kmp.prefix[i]
			if kmp.prefix[i] > -1 {
				i = kmp.prefix[i]
			} else {
				i = 0
			}
		}
	}
	return -1
}

func matchStr(str1, str2 string) bool {
	str1 = str1 + str1
	kmp, err := NewKMP(str2)

	if err != nil {
		fmt.Print("error: %s\n", err.Error())
		return false
	}

	return kmp.FindStringIndex(str1) != -1
}

func main() {
	str1 := "abcde"
	str2 := "eab"

	if matchStr(str1, str2) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
