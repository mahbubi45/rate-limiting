package controller

import (
	"bufio"
	"fmt"
	"os"
	"rate-limiting/redis"
	"strconv"
	"strings"
	"time"
)

func (s *Server) RateLimitingController() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama user: ")
	nameUser, _ := reader.ReadString('\n')
	nameUser = strings.TrimSpace(nameUser) // Hapus newline di akhir (\n)

	fmt.Print("Masukkan Jumlah Request: ")
	inputJumlahHitApi, _ := reader.ReadString('\n')
	inputJumlahHitApi = strings.TrimSpace(inputJumlahHitApi)

	numRequest, err := strconv.Atoi(inputJumlahHitApi)

	if err != nil {
		fmt.Println("NumRequest Harus Berupa Angka yaa!!")
		return
	}

	for i := 1; i <= numRequest; i++ {
		if redis.RateLimiting(nameUser) {
			fmt.Printf("Request %d dari %s: Akses diizinkan ✅\n", i, nameUser)
		} else {
			fmt.Printf("Request %d dari %s: Terlalu banyak request! ❌\n", i, nameUser)
		}
		time.Sleep(500 * time.Millisecond)
	}

}
