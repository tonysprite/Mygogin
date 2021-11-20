package models

type User struct {
	Id        int    `json:"user id"`
	Name      string `json:"user name"`
	Password  string `json:"user password"`
	Email     string `json:"user email"`
	Mobile    string `json:"user mobile"`
	Create_at int    `json:"user create time"`
}

func LoginQuery(user_name, password string) (userData *User) {
	userData = new(User)
	userData.Create_at = 123
	userData.Mobile = "183232323"
	userData.Email = "eaas@qq.com"
	userData.Password = password
	userData.Name = user_name
	userData.Id = 1232
	return
}

func LogoutAction(user_id int) (IsLogout bool) {
	IsLogout = true
	return
}

func RegistAction(user_name string, password string, email string, mobile string) int {
	return 0
}
