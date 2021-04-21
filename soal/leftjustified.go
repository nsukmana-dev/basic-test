package soal

import (
	"strings"
)

func Leftjustified(word string, size int) string {
	var s string
	kata := strings.Fields(word)
	var baris string
	for i := 0; i < len(kata); i++ {
		newpar := baris + kata[i]
		if len(newpar) <= 15 {
			baris += kata[i] + " "
		} else {
			if i == len(kata)-1 {
				s += baris + "\n"
				s += kata[i]
			} else {
				s += baris + "\n"
				baris = kata[i] + " "
			}

		}

	}
	return s
}
