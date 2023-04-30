package model

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
    Password string `json:"password"`
}

type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserLogin struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

func (u *UserInfo) UpdateUser(user *User) *User {
    user.Name = u.Name
    user.Email = u.Email
    user.Phone = u.Phone
    return user
} 

func (u *User) CleanUser() UserInfo {
    var user UserInfo
    user.Name = u.Name
    user.Email = u.Email
    user.Phone = u.Phone
    return user
}
