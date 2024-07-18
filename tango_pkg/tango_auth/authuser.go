package tango_auth

type AuthUser struct {
	User  *User
	Perms AuthPerms
	Token string
}

func NewAuthUser(user *User, token string) *AuthUser {
	return &AuthUser{
		User:  user,
		Token: token,
	}
}
