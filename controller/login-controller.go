package controller

import (
	"bufio"
	"fmt"
	"os"
	"rate-limiting/helper"
	"strings"
)

func (s *Server) LogInController() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("masukkan email: ")
		inputEmailUser, _ := reader.ReadString('\n')
		inputEmailUser = strings.TrimSpace(inputEmailUser) // Hapus newline di akhir (\n)

		// validatonChekEmail := isEmailValid(inputEmailUser)
		if !helper.IsEmailValid(inputEmailUser) {
			fmt.Println("email harus menggunakan @gmail.com")
			continue
		}

		fmt.Print("masukkan password user: ")
		inputPasswordUser, _ := reader.ReadString('\n')
		inputPasswordUser = strings.TrimSpace(inputPasswordUser) // Hapus newline di akhir (\n)

		if !helper.IsPasswordValid(inputPasswordUser) {
			fmt.Println("password harus menggunakan @!")
			continue
		}

		if inputEmailUser == "mahbubi@gmail.com" && inputPasswordUser == "mahbubi123!@" {
			fmt.Println("berhasil login")
			break
		} else {
			fmt.Println("gagal masuk coba lagi")
		}
	}
}
