package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (s *Server) PenjumlahanController(jumlahAwal, jumlahAkhir int) int {

	total := jumlahAwal + jumlahAkhir
	fmt.Print("ini hasil perhitungan anda: ", total, "\n")

	return total
}

func (s *Server) PerhitunganController() {
	reader := bufio.NewReader(os.Stdin)

	// var jumlah int

	fmt.Print("Masukkan nama nilai awal: ")
	nilaiAwal, _ := reader.ReadString('\n')
	nilaiAwal = strings.TrimSpace(nilaiAwal)
	intNilaiAwal, _ := strconv.Atoi(nilaiAwal)

	fmt.Print("Masukkan nama nilai kedua: ")
	nilaiKedua, _ := reader.ReadString('\n')
	nilaiKedua = strings.TrimSpace(nilaiKedua)
	intNilaiKedua, _ := strconv.Atoi(nilaiKedua)

	// jumlah = intNilaiAwal +
	s.PenjumlahanController(intNilaiAwal, intNilaiKedua)

	// fmt.Println("ini hasil perhitungan anda: ", jumlah)

}
