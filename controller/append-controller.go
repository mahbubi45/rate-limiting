package controller

import "fmt"

type Testing struct {
	NilaiA int
	NilaiB int
	NilaiC int
}

func (s *Server) AppendController() {
	penampungDataUser := []string{}
	dataUser := []string{"user 1", "user 2", "user3"}

	penggabunganData := append(penampungDataUser, dataUser...)
	fmt.Print("ini data user anda ketika sudah di gabungin ya: \n", penggabunganData, "\n")
}

// hanya 1 element
func (s *Server) SliceDataTesting() {
	testings := Testing{
		NilaiA: 20,
		NilaiB: 15,
		NilaiC: 0,
	}

	var total = testings.NilaiA + testings.NilaiB

	testings.NilaiC = total

	testing := []Testing{testings}

	// testingAppend := append(testing, testings)

	fmt.Print("ini data ya : \n", testing, "\n")

}

func (s *Server) AppendSliceTestingController() {
	testings := []Testing{
		{NilaiA: 10, NilaiB: 15, NilaiC: 20}, //element 1
		{NilaiA: 11, NilaiB: 16, NilaiC: 30}, //element 2
		{NilaiA: 12, NilaiB: 17, NilaiC: 40}, //element 3
	}

	penampungTesting := []Testing{}
	// var totalDataKeseluruhan int
	totalDataKeseluruhan := 0

	for _, testing := range testings {
		penampungTesting = append(penampungTesting, testing)

		// fmt.Print("ini data append ya : \n", penampungTesting, "\n")
		// fmt.Print("ini data append ya : \n", testing, "\n")
		// fmt.Print("ini total keseluruhan data ketika di jumlah ya : \n", totalDataKeseluruhan, "\n")

		totalDataKeseluruhan += testing.NilaiA + testing.NilaiB + testing.NilaiC
		fmt.Println("Setelah perkalian:", totalDataKeseluruhan)
	}

	fmt.Print("ini data append ya : \n", penampungTesting, "\n")

	// jika total nilai 50 maka nilai B
	if totalDataKeseluruhan == 50 && totalDataKeseluruhan <= 50 {
		fmt.Print("nilai kamu B : ", totalDataKeseluruhan, "\n")
	} else if totalDataKeseluruhan > 51 && totalDataKeseluruhan <= 80 {
		fmt.Print("nilai kamu  B+ : ", totalDataKeseluruhan, "\n")
	} else if totalDataKeseluruhan > 80 && totalDataKeseluruhan <= 90 {
		fmt.Print("nilai kamu  A : ", totalDataKeseluruhan, "\n")
	} else if totalDataKeseluruhan > 90 && totalDataKeseluruhan <= 10000000000 {
		fmt.Print("nilai kamu  A+ : ", totalDataKeseluruhan, "\n")
	}
}
