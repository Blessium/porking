package model

type User struct {
	ID    uint   `gorm:"primaryKey"`
    Name  string `validate:"required"`
	Email string `validate:"required"`
    Phone string `validate:"required"`
    Password string `validate:"required"`
    Role string `validate:"required"`
    Cars []Car
    Reservations []Reservation
}

type UserInfo struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Phone string `json:"phone"`
}

type UserLogin struct {
    Email string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required"`
}

func (u *UserInfo) UpdateUser(user *User) *User {
    if (u.Name != "") { user.Name = u.Name }
    if (u.Email != "") { user.Email = u.Email }
    if (u.Phone != "") { user.Phone = u.Phone }
    return user
} 

func (u *User) CleanUser() UserInfo {
    var user UserInfo
    if (u.Name != "") { user.Name = u.Name }
    if (u.Email != "") { user.Email = u.Email }
    if (u.Phone != "") { user.Phone = u.Phone }
    return user
}
