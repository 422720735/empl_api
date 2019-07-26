package defs


type Users struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Position int `json:"position"`
}
