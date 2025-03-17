package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (s *Server) TampilkanDataUser() {
	reader := bufio.NewReader(os.Stdin)
	// ambiil field yang akan di isi
	user := User{}

	fmt.Print("masukkan nama user: ")
	inputNameUser, _ := reader.ReadString('\n')
	inputNameUser = strings.TrimSpace(inputNameUser) // Hapus newline di akhir (\n)

	fmt.Print("Masukkan Jumlah Request: ")
	inputAgeUser, _ := reader.ReadString('\n')
	inputAgeUser = strings.TrimSpace(inputAgeUser)

	user.Name = inputNameUser
	convertnNumberAge, _ := strconv.Atoi(inputAgeUser)
	user.Age = convertnNumberAge

	// ini buat (wadah)/penampung user dan field nya akan di tarok di dalam yang sudah terisi data
	penampungUsers := []User{user}

	for _, users := range penampungUsers {
		var data = fmt.Sprintf("Go routine \n Name : %s,\n Age: %d, \n", users.Name, users.Age)

		fmt.Println("ini data anda ya!!", data)

		if users.Age == 15 || users.Age <= 15 {
			fmt.Println("ternyata kamu masi bocil ya!!")
		} else if users.Age == 20 {
			fmt.Println("ternyata kamu masi remaja ya!!")
		} else if users.Age == 30 {
			fmt.Println("ternyata kamu sudah meranjak om om ya!!")
		} else {
			fmt.Println("ternyata kamu sudah tua ya !!")
		}
	}
}
