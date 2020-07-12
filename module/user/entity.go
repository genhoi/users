package user

type User struct {
	Id       uint64 `json:"id" db:"id, primarykey"`
	Login    string `json:"login" db:"login"`
	FullName string `json:"full_name" db:"full_name"`
}

type GenerateUser struct {
	User User
	Err  error
}
