package tango_auth

import (
	"html/template"

	"tango_pkg/tango_errors"
	"tango_pkg/tango_validator"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const passMinLen = 8
const passMaxLen = 36
const passCost = 10

const errorWrongLogin = "Wrong user or password"

/*
	This struct is for security
	Here comes every methods to login or check users
*/

type Auth struct {
	gorm.Model
	// UserID         uint
	UserID         uuid.UUID
	User           User `gorm:"references:UID"`
	Passwd         string
	IsActive       bool
	ActivationCode string
}
type AuthDTO struct {
	ID             uuid.UUID `json:"id" param:"id" query:"id" form:"id"`
	Email          string    `json:"email" param:"email" query:"email" form:"email"`
	Token          string    `json:"token" param:"token" query:"token" form:"token"`
	IsActive       bool      `json:"is_active" param:"is_active" query:"is_active" form:"is_active"`
	ActivationCode bool      `json:"code" param:"code" query:"code" form:"code"`
	Password       string    `json:"password" param:"password" query:"password" form:"password"`
}

type AuthDTOLogin struct {
	Email    string `json:"email" param:"email" query:"email" form:"email" validate:"required,email"`
	Password string `json:"password" param:"password" query:"password" form:"password" validate:"required"`
}

type AuthDTOCreate struct {
	UserID   uuid.UUID `json:"user_id" param:"user_id" query:"user_id" form:"user_id"`
	Password string    `json:"password" param:"password" query:"password" form:"password"`
}

type AuthDTOChangePassword struct {
	UserID      uuid.UUID `json:"id" param:"id" query:"id" form:"id" validate:"required"`
	OldPassword string    `json:"old" param:"old" query:"old" form:"old" validate:"required"`
	Password    string    `json:"new" param:"new" query:"new" form:"new" validate:"required"`
}

func NewAuth() *Auth {
	return &Auth{}
}

func (u *Auth) Create(db *gorm.DB, userID uuid.UUID, password string) error {
	authDTO := &AuthDTOCreate{
		UserID:   userID,
		Password: password,
	}

	u.satinizeDTOCreate(authDTO)

	newAuth := &Auth{
		UserID: authDTO.UserID,
		Passwd: authDTO.Password,
	}

	result := db.Create(newAuth)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *Auth) UpdatePassword(db *gorm.DB, dto *AuthDTOChangePassword) (bool, error) {
	u.satinizeDTOChangePassword(dto)

	// check old password is valid
	if !u.CheckPassword(db, dto.UserID, dto.OldPassword) {
		return false, tango_errors.ReturnDefault("Change password", "Incorrect old password", 0)
	}

	db.Model(&Auth{}).Where("user_id =?", dto.UserID).Update("passwd", dto.Password)

	return true, nil
}

func (u *Auth) Login(db *gorm.DB, dto *AuthDTOLogin) (*AuthUser, error) {

	email, err := tango_validator.ValidateEmail(dto.Email)
	if err != nil {
		return &AuthUser{}, tango_errors.ReturnDefault("Login", errorWrongLogin, 0)
	}
	userModel := NewUser()
	user, err := userModel.FindOneByEmail(db, email)

	if err != nil || user.ID == 0 {
		return &AuthUser{}, tango_errors.ReturnDefault("Login", errorWrongLogin, 0)
	}

	dto.Password = template.HTMLEscapeString(dto.Password)

	if !u.CheckPassword(db, user.UID, dto.Password) {
		return &AuthUser{}, tango_errors.ReturnDefault("Login", errorWrongLogin, 1)
	}

	au := &AuthUser{
		User:  user,
		Token: "",
	}

	return au, nil
}

func (u *Auth) Logout(db *gorm.DB, id int) (*User, error) {
	var user User

	return &user, nil
}

func (u *Auth) Check(db *gorm.DB, id int) (*User, error) {
	var user User

	return &user, nil
}

func (u *Auth) CheckPassword(db *gorm.DB, id uuid.UUID, password string) bool {
	auth := &Auth{}
	db.First(&auth, "user_id=?", id)

	return u.PasswdChecker(auth.Passwd, password)

}

func (u *Auth) PasswdCipher(password string) (string, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), passCost)

	return string(hashed), nil
}

func (u *Auth) PasswdChecker(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// validate and chiper the password
func (u *Auth) ProcessPasswd(password string) (string, error) {
	// Validate and Sanitize password
	pass, err := tango_validator.ValidatePassword(password, passMinLen, passMaxLen)
	if err != nil {
		return "", err
	}
	password, err = u.PasswdCipher(pass)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (u *Auth) satinizeDTOCreate(dto *AuthDTOCreate) error {
	passwd, err := u.ProcessPasswd(dto.Password)
	if err != nil {
		return err
	}
	dto.Password = passwd
	return nil
}

// func (u *Auth) satinizeDTOUpdate(dto *AuthDTO) error {
// 	// TODO
// 	return nil
// }

func (u *Auth) satinizeDTOChangePassword(dto *AuthDTOChangePassword) error {
	// old password
	oldPasswd, err := u.ProcessPasswd(dto.OldPassword)
	if err != nil {
		return err
	}
	dto.OldPassword = oldPasswd

	// old password
	passwd, err := u.ProcessPasswd(dto.Password)
	if err != nil {
		return err
	}
	dto.Password = passwd
	return nil
}
