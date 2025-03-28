package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Count int

func (s *Server) PenjumlahanController() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan angka awal: ")
	angkaAwal, _ := reader.ReadString('\n')
	angkaAwal = strings.TrimSpace(angkaAwal)
	angkaAwalInt, _ := strconv.Atoi(angkaAwal)

	fmt.Print("Masukkan angka kedua: ")
	angkaKedua, _ := reader.ReadString('\n')
	angkaKedua = strings.TrimSpace(angkaKedua)
	angkaKeduaInt, _ := strconv.Atoi(angkaKedua)

	var jumlah Count = Count(angkaAwalInt) + Count(angkaKeduaInt)

	fmt.Println("ini jumlah angka anda yang di jumlahkan menggunakan type alias ya: ", jumlah)
}
