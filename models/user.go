package models

// A user of this webservice.
type User struct {
	ID	int
	FirstName 	string
	LastName	string
}

var (
	users []*User // a slice of pointers to User structs
	nextID = 1
)

// Gets the users
func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}