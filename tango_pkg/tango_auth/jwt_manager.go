package tango_auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// type Keyfunc func (*Token) (interface{}, error)

func JWTMiddleware(e *echo.Group) {
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSecretKey),
	}))
}

func GetJWTSecret() string {
	return jwtSecretKey
}

// GenerateTokensAndSetCookies generates jwt token and saves it to the http-only cookie.
func GenerateTokensAndSetCookies(auth *AuthUser, c echo.Context) (string, error) {
	accessToken, exp, err := generateAccessToken(auth)
	if err != nil {
		return "", err
	}

	setTokenCookie(AccessTokenCookieName, accessToken, exp, c)
	return accessToken, nil
}

func generateAccessToken(auth *AuthUser) (string, time.Time, error) {
	// Declare the expiration time of the token (1h).
	expirationTime := time.Now().Add(1 * time.Minute)

	return generateToken(auth, expirationTime, []byte(GetJWTSecret()))
}

// Pay attention to this function. It holds the main JWT token generation logic.
func generateToken(auth *AuthUser, expirationTime time.Time, secret []byte) (string, time.Time, error) {

	claims := jwt.MapClaims{
		"user": auth.User.Email,
	}

	// Declare the token with the HS256 algorithm used for signing, and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string.
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

// Here we are creating a new cookie, which will store the valid JWT token.
func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

// JWTErrorChecker will be executed when user try to access a protected path.
func JWTErrorChecker(err error, c echo.Context) error {
	// Redirects to the signIn form.
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
}
