package response

// penampung get dari json
type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}
