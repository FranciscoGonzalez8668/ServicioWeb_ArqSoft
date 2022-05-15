package dto

type User struct {
	Id_user  string `json:"id_user"`
	Name     string `json:"name"`
	Lname    string `json:"lname"`
	Email    string `jsno:"email"`
	Password string `json:"password"`
	Cart     string `json:"cart"`
	Token    string `json:"token"`
}

type Users []User
