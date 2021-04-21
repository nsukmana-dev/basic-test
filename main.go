package main

import (
	"basic/soal"
	"fmt"
)

func main() {
	tulisan, jumlah := soal.APlusB("16 17")

	fmt.Println("=== No 1 ===")
	fmt.Println(tulisan)
	fmt.Println(jumlah)
	fmt.Println("=================")

	fmt.Println("=== No 2 ===")
	rekasinergi := soal.Rekasinergi(20, 3, 5)
	fmt.Println(rekasinergi)
	fmt.Println("=================")

	fmt.Println("=== No 3 ===")
	word := "Pohon raksasa berjenis jambi-jambi tumbang dan berdiri lagi di lingkungan"
	leftjustified := soal.Leftjustified(word, 15)
	fmt.Println(leftjustified)
}
