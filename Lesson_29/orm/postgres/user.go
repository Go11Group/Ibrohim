package user

import (
	"User_Gorm/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (u *UserRepo) CreateUserTable() error {
	return u.DB.AutoMigrate(&model.User{})
}

func (u *UserRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User
	res := u.DB.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (u *UserRepo) GetUser(firstName string) (*model.User, error) {
	var user model.User
	res := u.DB.First(&user, "First_Name = ?", firstName)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u *UserRepo) Create(user model.User) error {
	res := u.DB.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserRepo) Update(firstName string, user model.User) error {
	newInfo := map[string]interface{}{
        "LastName":    user.LastName,
        "Email":       user.Email,
        "Password":    user.Password,
        "Age":         user.Age,
        "Field":       user.Field,
        "Gender":      user.Gender,
        "IsEmployee":  user.IsEmployee,
    }

	res := u.DB.Model(&model.User{}).Where("First_Name = ?",firstName).Updates(newInfo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *UserRepo) Delete(firstName string) error {
	res := u.DB.Delete(&model.User{},"First_Name = ?",firstName)
	if res.Error != nil {
		return res.Error
	}
	return nil
}