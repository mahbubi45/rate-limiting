package controller

import "fmt"

type UserRBAC struct {
	UserName string
	Role     Role
}

type Role struct {
	NameRole   string
	Permission []string
}

func HashMiddlwarePermission(user UserRBAC, permission string) string {
	for _, userRolePermission := range user.Role.Permission {
		if userRolePermission == permission {
			return "200 : kamu memiliki acces ini"
		}
	}

	return "403 : maaf kamu tidak memiliki acces terhadap fitur ini ya !!"
}

func DataUserRBAC() []UserRBAC {
	dataUserRBAC := []UserRBAC{
		{
			UserName: "mahbubi",
			Role: Role{
				NameRole:   "Admin",
				Permission: []string{"Created", "Read", "Update", "Delete"},
			},
		},
		{
			UserName: "nano",
			Role: Role{
				NameRole:   "Editor",
				Permission: []string{"Read", "Update"},
			},
		},
		{
			UserName: "ajax",
			Role: Role{
				NameRole:   "User",
				Permission: []string{"Read"},
			},
		},
	}

	return dataUserRBAC
}

func (s *Server) RBACController() {
	// var admin = UserRBAC{
	// 	UserName: "mahbubi",
	// 	Role: Role{
	// 		NameRole:   "Admin",
	// 		Permission: []string{"Created", "Read", "Update", "Delete"},
	// 	},
	// }

	// var editor = UserRBAC{
	// 	UserName: "nano",
	// 	Role: Role{
	// 		NameRole:   "Editor",
	// 		Permission: []string{"Read", "Update"},
	// 	},
	// }

	// var user = UserRBAC{
	// 	UserName: "ajax",
	// 	Role: Role{
	// 		NameRole:   "User",
	// 		Permission: []string{"Read"},
	// 	},
	// }

	datauserRbac := DataUserRBAC()

	fmt.Println("admin bisa semuanya ya : ", HashMiddlwarePermission(datauserRbac[0], "Delete"))
	fmt.Println("editor bisa baca ya : ", HashMiddlwarePermission(datauserRbac[1], "Read"))
	fmt.Println("user apakah bisa creatd ya ? : ", HashMiddlwarePermission(datauserRbac[2], "Created"))

}
