package main

import (
	"fmt"
	"math"
)

func main() {

	// Data Uji 1

	// Soal Pertama (awal)
	// Data Mark
	markMassaUji1 := 78.0  // kg
	markTinggiUji1 := 1.69 // meter

	// Data John
	johnMassaUji1 := 92.0  // kg
	johnTinggiUji1 := 1.95 // meter

	// Soal Pertama (Akhir)

	// Menghitung versi 1
	markBMIversi1Uji1 := markMassaUji1 / (markTinggiUji1 * markTinggiUji1)
	johnBMIversi1Uji1 := johnMassaUji1 / (johnTinggiUji1 * johnTinggiUji1)

	// Menghitung versi 2

	johnBMIversi2Uji1 := math.Pow(johnMassaUji1/johnTinggiUji1, 2)
	markBMIversi2Uji1 := math.Pow(markMassaUji1/markTinggiUji1, 2)

	// Membandingkan BMI
	markHigherBMI := markBMIversi1Uji1 > johnBMIversi1Uji1 || markBMIversi2Uji1 > johnBMIversi2Uji1

	// Data Uji 2

	// Soal Pertama (awal)
	// Data Mark
	markMassaUji2 := 95.0  // kg
	markTinggiUji2 := 1.88 // meter

	// Data John
	johnMassaUji2 := 85.0  // kg
	johnTinggiUji2 := 1.76 // meter

	// Soal Pertama (Akhir)

	// Menghitung versi 1
	markBMIversi1Uji2 := markMassaUji2 / (markTinggiUji2 * markTinggiUji2)
	johnBMIversi1Uji2 := johnMassaUji2 / (johnTinggiUji2 * johnTinggiUji2)

	// Menghitung versi 2

	johnBMIversi2Uji2 := math.Pow(johnMassaUji2/johnTinggiUji2, 2)
	markBMIversi2Uji2 := math.Pow(markMassaUji2/markTinggiUji2, 2)

	// Membandingkan BMI
	markHigherBMIUji2 := markBMIversi1Uji2 > johnBMIversi1Uji2 || markBMIversi2Uji2 > johnBMIversi2Uji2

	fmt.Println("Soal untuk Data Uji Pertama ")
	fmt.Println("Mark Versi 1 :", markBMIversi1Uji1)
	fmt.Println("Mark Versi 2:", markBMIversi2Uji1)
	fmt.Println("John Versi 1:", johnBMIversi1Uji1)
	fmt.Println("John Versi 2:", johnBMIversi2Uji1)
	fmt.Println("Apakah BMI Mark lebih tinggi dari John?", markHigherBMI)

	fmt.Println("Soal untuk Data Uji Kedua ")
	fmt.Println("Mark Versi 1 :", markBMIversi1Uji2)
	fmt.Println("Mark Versi 2:", markBMIversi2Uji2)
	fmt.Println("John Versi 1:", johnBMIversi1Uji2)
	fmt.Println("John Versi 2:", johnBMIversi2Uji2)
	fmt.Println("Apakah BMI Mark lebih tinggi dari John?", markHigherBMIUji2)
}
