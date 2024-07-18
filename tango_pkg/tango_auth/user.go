package tango_auth

import (
	"fmt"
	"strings"
	"text/template"

	"tango_pkg/tango_errors"
	"tango_pkg/tango_validator"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
	This struct is only for user information like personal data
*/

type User struct {
	gorm.Model
	UID      uuid.UUID `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	IsActive bool
}
type UserDTO struct {
	UID   uuid.UUID `json:"id" param:"id" query:"id" form:"id"`
	Name  string    `json:"name" param:"name" query:"name" form:"name"`
	Email string    `json:"email" param:"email" query:"email" form:"email"`
}

type UserDTOCreate struct {
	Name     string `json:"name" param:"name" query:"name" form:"name" validate:"required"`
	Email    string `json:"email" param:"email" query:"email" form:"email" validate:"required,email"`
	Password string `json:"password" param:"password" query:"password" form:"password" validate:"required"`
}

type UserCounter struct {
	Total int
}

func NewUser() *User {
	return &User{}
}

func (u *User) ConvertToDTO() *UserDTO {
	return &UserDTO{
		UID:   u.UID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (u *User) Count(db *gorm.DB) (int, error) {
	counter := &UserCounter{}
	db.Model(&User{}).Select("count(ID) as total").Find(&counter)
	return counter.Total, nil
}

func (u *User) FindOne(db *gorm.DB, id int) (*User, error) {
	var user User
	db.First(&user, id)
	if user.ID == 0 {
		return nil, tango_errors.ReturnModel("User", tango_errors.MsgNotFound(), 0)
	}
	return &user, nil
}

func (u *User) FindOneByUID(db *gorm.DB, uid uuid.UUID) (*User, error) {
	var user User
	db.First(&user, "uid=?", uid)
	if user.ID == 0 {
		return nil, tango_errors.ReturnModel("User", tango_errors.MsgNotFound(), 0)
	}
	return &user, nil
}

func (u *User) IsEmailAvailable(db *gorm.DB, email string) bool {
	var user User
	db.First(&user, "email=?", email)
	if user.ID == 0 {
		return true
	} else {
		return false
	}
}

func (u *User) FindOneByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	db.First(&user, "email=?", email)
	if user.ID == 0 {
		return nil, tango_errors.ReturnModel("User", tango_errors.MsgNotFound(), 0)
	}
	return &user, nil
}

func (u *User) FindAll(db *gorm.DB) (*[]User, error) {
	var users []User
	db.Order("created_at ASC").Find(&users)
	if len(users) <= 0 {
		return nil, &tango_errors.ModelError{
			ModelName: "User",
			Code:      0,
			Message:   tango_errors.MsgZeroRecordsFound(),
		}
	}
	return &users, nil
}

func (u *User) FindAllPagination(db *gorm.DB, itemsPerPage, currentPage int) (*[]User, error) {
	users := []User{}

	db.Order("created_at ASC").Limit(itemsPerPage).Offset(itemsPerPage * currentPage).Find(&users)
	if len(users) <= 0 {
		return nil, tango_errors.ReturnModel("User", tango_errors.MsgZeroRecordsFound(), 0)
	}
	return &users, nil
}

func (u *User) Create(db *gorm.DB, dto UserDTOCreate) (*User, error) {
	u.satinizeDTOCreate(&dto)
	user := &User{
		UID:   uuid.New(),
		Name:  dto.Name,
		Email: dto.Email,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return &User{}, result.Error
	}
	// create auth
	auth := NewAuth()
	err := auth.Create(db, user.UID, dto.Password)
	if err != nil {
		return &User{}, nil
	}
	return user, nil
}

func (u *User) Update(db *gorm.DB, id uuid.UUID, dto UserDTO) (*User, error) {
	u.satinizeDTOUpdate(&dto)

	user := &User{}
	db.First(user, "uid=?", id)
	if user.ID == 0 {
		return user, tango_errors.ReturnModel("User", tango_errors.MsgUIDNotFound(id.String()), 0)
	}

	user.Name = dto.Name
	user.Email = dto.Email

	db.Save(user)

	return user, nil
}

func (u *User) Delete(db *gorm.DB, id uuid.UUID) (*User, error) {
	user, err := u.FindOneByUID(db, id)
	if err != nil {
		return nil, err
	}
	db.Delete(&user)
	return user, nil
}

func (u *User) GetIDAsString() string {
	return fmt.Sprintf("%d", u.ID)
}

func (u *User) satinizeDTOCreate(dto *UserDTOCreate) error {
	// name
	dto.Name = strings.TrimSpace(template.HTMLEscapeString(dto.Name))
	// validate and satinize
	email, err := tango_validator.ValidateEmail(dto.Email)
	if err != nil {
		return err
	}
	dto.Email = email
	return nil
}

func (u *User) satinizeDTOUpdate(dto *UserDTO) error {
	dto.Name = strings.TrimSpace(dto.Name)
	email, err := tango_validator.ValidateEmail(dto.Email)
	if err != nil {
		return err
	}
	dto.Email = email
	return nil
}
