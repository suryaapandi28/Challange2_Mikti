package main

import (
	"fmt"
)

func Hitungratarata(skortim []int) float64 {
	total := 0
	for _, score := range skortim {
		total += score
	}
	return float64(total) / float64(len(skortim))
}

func main() {
	// menjumlahkan nilai atau akumulasi
	skortim := map[string][]int{
		"Lumba-lumba": {97, 112, 101},
		"Koala":       {109, 95, 123},
	}

	// Hitung skor rata-rata untuk setiap tim
	tim1 := "Lumba-lumba"
	tim2 := "Koala"

	ratarata1 := Hitungratarata(skortim[tim1])
	ratarata2 := Hitungratarata(skortim[tim2])

	// Hitung skor rata-rata dalam data uji 1
	fmt.Println("Soal Pertama : ")
	fmt.Println("Lumba-Lumba : ", ratarata1)
	fmt.Println("Koala : ", ratarata2)

	// Melakukan perbandingan skor rata rata dalam data uji 1
	fmt.Println("Soal Kedua : ")
	if ratarata1 > ratarata2 {
		fmt.Printf("Pemenang kompetisi adalah %s dengan skor rata-rata %.2f.\n", tim1, ratarata1)
	} else if ratarata2 > ratarata1 {
		fmt.Printf("Pemenang kompetisi adalah %s dengan skor rata-rata %.2f.\n", tim2, ratarata2)
	} else {
		fmt.Println("Hasil kompetisi adalah seri.")
	}

	// Bonus 1 & 2 menggabungkan kedua soal bonus menjadi satu

	fmt.Println("Soal Bonus Pertama & Kedua : ")
	skorLumbaLumba := []int{97, 112, 101}
	skorKoala := []int{109, 95, 123}
	rataRataLumbaLumba := hitungRataRataBonus1(skorLumbaLumba)
	rataRataKoala := hitungRataRataBonus1(skorKoala)

	pemenang := menentukanPemenangBonus1(rataRataLumbaLumba, rataRataKoala)

	fmt.Println("Tim pemenang adalah:", pemenang)
}

func hitungRataRataBonus1(skor []int) float64 {
	var total int
	for _, nilai := range skor {
		total += nilai
	}
	return float64(total) / float64(len(skor))
}

func menentukanPemenangBonus1(rataRataLumbaLumba, rataRataKoala float64) string {
	if rataRataLumbaLumba < 100 && rataRataKoala < 100 {
		return "Seri"
	} else if rataRataLumbaLumba >= 100 && rataRataKoala >= 100 {
		if rataRataLumbaLumba > rataRataKoala {
			return "Lumba-lumba"
		} else if rataRataKoala > rataRataLumbaLumba {
			return "Koala"
		} else {
			return "Seri"
		}
	} else if rataRataLumbaLumba >= 100 {
		return "Lumba-lumba"
	} else {
		return "Koala"
	}
}
