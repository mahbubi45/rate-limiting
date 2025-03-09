package main

import (
	"bufio"
	"fmt"
	"os"
	"rate-limiting/redis"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama user: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user) // Hapus newline di akhir

	fmt.Print("Masukkan Jumlah Request: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numRequest, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("NumRequest Harus Berupa Angka yaa!!")
		return
	}

	for i := 1; i <= numRequest; i++ {
		if redis.RateLimiting(user) {
			fmt.Printf("Request %d dari %s: Akses diizinkan ✅\n", i, user)
		} else {
			fmt.Printf("Request %d dari %s: Terlalu banyak request! ❌\n", i, user)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
