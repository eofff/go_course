package models

type User struct {
	FirstName string
	Email     string
	Username  string
}

func GetUsers() []User {
	users := make([]User, 3)
	users[0] = User{"Some name", "test@test.com", "fdscxv"}
	users[1] = User{"Another name", "hz@hz.hz", "aasd"}
	users[2] = User{"Afdsi", "fff@asczx.net", "7asfgvsd"}
	return users
}
