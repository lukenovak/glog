package models

// User is used solely for login/signup
type User struct {
	Username string // primary key
	HashedPW []byte
}

// Author is used to get metadata about a user- Mostly to associate their name with an article
type Author struct {
	Id int
	Username string
	FullName string
}
