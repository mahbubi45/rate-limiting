package controller

import (
	"fmt"
	"reflect"
	"runtime"
)

type User struct {
	Name string
}

func (s *Server) isiDatauser() []User {
	dataUser := []User{
		{
			Name: "user1",
		},
		{
			Name: "user2",
		},
		{
			Name: "user3",
		},
	}

	return dataUser

}

// p kecil fungsi menandakan private
func (s *Server) printMessage(what chan string) {
	fmt.Println(<-what)
}

func (s *Server) ChannelGoRoutineController() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	dataUserStruct := s.isiDatauser()

	// convert dan hanya ambil valuenya saja dari struct user
	var userStructSlice []string //penampung data yang hanya di ambil fieldnya saja
	for _, users := range dataUserStruct {
		valueUser := reflect.ValueOf(users)
		for i := 0; i < valueUser.NumField(); i++ {
			userStructSlice = append(userStructSlice, valueUser.Field(i).String())
		}
	}

	dataUserArrayString := userStructSlice

	for _, each := range dataUserArrayString {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			message <- data
		}(each)
	}

	jumlahDataUser := len(dataUserArrayString)

	//jalankan perulangan sesuai jumlahDataUser dengan cara menghitung len
	for i := 0; i < jumlahDataUser; i++ {
		s.printMessage(message)
	}
}
