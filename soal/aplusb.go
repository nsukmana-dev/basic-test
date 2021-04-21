package soal

import (
	"fmt"
	"log"
)

func APlusB(angka string) (tulisan string, jumlah int16) {
	var a, b int16

	_, err := fmt.Sscanf(angka, "%d %d", &a, &b)

	if err != nil {
		log.Fatal(err)
	}

	tulisan = fmt.Sprintf("%d", a) + " + " + fmt.Sprintf("%d", b)
	jumlah = a + b
	return tulisan, jumlah

}
