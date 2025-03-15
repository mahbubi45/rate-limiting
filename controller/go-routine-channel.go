package controller

import (
	"fmt"
	"runtime"
)

type User struct {
	Name string
	Age  int
}

func (s *Server) isiDatauser() []User {
	dataUser := []User{
		{
			Name: "user1",
			Age:  15,
		},
		{
			Name: "user2",
			Age:  16,
		},
		{
			Name: "user3",
			Age:  17,
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
	// var userStructSlice []string //penampung data yang hanya di ambil fieldnya saja
	// for _, users := range dataUserStruct {
	// 	valueUser := reflect.ValueOf(users)
	// 	for i := 0; i < valueUser.NumField(); i++ {
	// 		fieldValue := fmt.Sprintf("%v", valueUser.Field(i).Interface())
	// 		userStructSlice = append(userStructSlice, fieldValue)
	// 	}
	// }

	// dataUserArrayString := userStructSlice

	for _, users := range dataUserStruct {
		go func(user User) {
			var data = fmt.Sprintf("Go routine \n Name : %s,\n Age: %d, \n", user.Name, user.Age)
			message <- data
		}(users)
	}

	jumlahDataUser := len(dataUserStruct)

	//jalankan perulangan sesuai jumlahDataUser dengan cara menghitung len
	for i := 0; i < jumlahDataUser; i++ {
		s.printMessage(message)
	}
}
