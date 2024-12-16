package models

type User struct {
	Name	string	`json:"name" param:"name"`
	Email	string	`json:"email"`
}

func AllUsers() (*[]User, error) {
	u := &[]User{}
	result := db.Find(u)
	return u, result.Error
}

func FindUser(u *User) (*User, error) {
	result := &User{} 
	if err := db.Where(u).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(u *User) error {
	result := db.Create(u)
	return result.Error
}

func DeleteUser(u *User) error {
	result := db.Where("email = ?", u.Email).Delete(u)
	return result.Error
}