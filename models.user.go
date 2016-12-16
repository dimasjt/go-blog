package main

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var userList = []user{
	user{ID: 1, Name: "John Cena", Email: "john@cena.com", Password: "iminvisible"},
}

func createUser(name, email, password string) (*user, error) {
	u := user{ID: len(userList) + 1, Name: name, Email: email, Password: password}

	userList = append(userList, u)

	return &u, nil
}
