package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"rate-limiting/response"
	"strconv"
)

// untuk ngebacai json data dummy
func (s *Server) ReadJsonUsersController() {

	var users response.Users
	jsonFile, err := os.Open("json/users.json")

	if err != nil {
		fmt.Println("gagal mengambail data users.json")
	}

	fmt.Println("sukses membuka file users.json")
	// setelah sukses maka di close
	defer jsonFile.Close()

	// ngebaca data json
	byteValueUsers, _ := ioutil.ReadAll(jsonFile)

	// mengubah (decode) data JSON dari byteValue menjadi struct atau slice di Go.
	json.Unmarshal(byteValueUsers, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Social Facebook: " + users.Users[i].Social.Facebook)
		fmt.Println("User Social twitter: " + users.Users[i].Social.Facebook)
	}

}
