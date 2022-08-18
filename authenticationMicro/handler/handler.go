package handler

type User struct {
	Email        string
	Username     string
	Passwordhash string
	Fullname     string
	CreateDate   string
	Role         int
}

func GetUserObject(Email string) (User, bool) {
	//need to be replace using database
	for _, User := range userList {
		if User.Email == Email {
			return User, true
		}
	}
	return User{}, false

}

//check if password hash is valid

func (u *User) ValidatePasswordHash(pswhash string) bool {

	return u.Passwordhash == pswhash
}

//this simple adds the User to the list
func AddUserObject(Email string, Username string, Passwordhash string, Fullname string, Role int) bool {

	//declare the new User object
	newUser := User{
		Email:        Email,
		Passwordhash: Passwordhash,
		Username:     Username,
		Fullname:     Fullname,
		Role:         Role,
	}
	//check if a user already exist
	for _, ele := range userList {
		if ele.Email == Email || ele.Username == Username {
			return false
		}
	}
	userList = append(userList, newUser)
	return true
}
