package soal

import "fmt"

func Rekasinergi(N, X, Y int16) []string {
	var s []string
	var i int16
	for i = 1; i < N; i++ {
		if i%X == 0 && i%Y == 0 {
			s = append(s, "RekaSinergi")
		} else if i%X == 0 {
			s = append(s, "Reka")
		} else if i%Y == 0 {
			s = append(s, "Sinergi")
		} else {
			s = append(s, fmt.Sprintf("%d", i))
		}
	}

	return s
}
