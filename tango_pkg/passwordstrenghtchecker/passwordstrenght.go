package passwordstrenght

import (
	"regexp"
	"unicode"
)

type PasswordStrenghtChecker struct {
	Min int
	Max int
}

type PasswordEvaluation struct {
	Strenght string
	Points   int
}

func NewPasswordStrenght(min, max int) *PasswordStrenghtChecker {
	return &PasswordStrenghtChecker{
		Min: min,
		Max: max,
	}
}

func (ps *PasswordStrenghtChecker) check(password string) int {
	// Puntuación inicial
	points := 0

	// Calcular puntuación basada en la longitud de la password
	points += ps.pointsForLength(password)

	// Calcular puntuación basada en la diversidad de caracteres
	points += ps.pointsForDiversity(password)

	// Calcular puntuación por el uso de caracteres especiales
	points += ps.pointsForSpecialChars(password)

	// Calcular puntuación por el uso de combinaciones de caracteres
	points += ps.pointsForMerging(password)

	return points
}

func (ps *PasswordStrenghtChecker) evaluate(points int) string {
	if points < 0 {
		return "weak"
	} else if points < 5 {
		return "moderate"
	} else {
		return "strong"
	}
}

func (ps *PasswordStrenghtChecker) CheckAndEvaluate(password string) *PasswordEvaluation {
	points := ps.check(password)
	eval := ps.evaluate(points)
	return &PasswordEvaluation{
		Strenght: eval,
		Points:   points,
	}
}

func (ps *PasswordStrenghtChecker) HasMinLen(password string) bool {
	return len(password) >= ps.Min
}

func (ps *PasswordStrenghtChecker) pointsForLength(password string) int {
	// Asignar puntuación basada en la longitud de la password
	pLen := len(password)
	if pLen >= ps.Min {
		return min(2, pLen-ps.Min)
	} else {
		return -2
	}
}

func (ps *PasswordStrenghtChecker) pointsForDiversity(password string) int {
	// Asignar puntuación basada en la diversidad de caracteres
	uniqueChars := make(map[rune]bool)
	for _, char := range password {
		uniqueChars[char] = true
	}
	diversity := len(uniqueChars)
	return ps.min(2, diversity-5)
}

func (ps *PasswordStrenghtChecker) pointsForSpecialChars(password string) int {
	// Asignar puntuación por el uso de caracteres especiales
	specialChars := regexp.MustCompile(`[ !@#$%^&*()_+{}\[\]:;<>,.?~\\/-]`)
	if specialChars.MatchString(password) {
		return 2
	}
	return 0
}

func (ps *PasswordStrenghtChecker) pointsForMerging(password string) int {
	// Asignar puntuación por el uso de combinaciones de caracteres
	hasUppersChars := false
	hasLowersChars := false
	hasNumber := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppersChars = true
		} else if unicode.IsLower(char) {
			hasLowersChars = true
		} else if unicode.IsNumber(char) {
			hasNumber = true
		}
	}

	if hasUppersChars && hasLowersChars && hasNumber {
		return 2
	}

	return 0
}

func (ps *PasswordStrenghtChecker) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
