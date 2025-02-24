package users

import (
	"carrent/db"
	"crypto/rand"
	"encoding/base64"
	"time"
)

var (
	// Максимальное количество попыток
	maxLoginAttempts uint = 5
	// Время блокировки после превышения попыток
	lockoutDuration = 5 * time.Minute
	// Длина куки
	COOKIE_LENGTH int = 256
)

// Функция для генерациия куки, возвращается в кодировке base64
func GenerateRandomToken() (string, error) {
	bytes := make([]byte, COOKIE_LENGTH)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// Функция для проверки попыток входа
func checkLoginAttempts(user db.User) bool {
	if user.LastFailedAttempt.IsZero() {
		return true // Нет неудачных попыток, можно пробовать
	}
	if time.Since(user.LastFailedAttempt) > lockoutDuration {
		// Блокировка снята, можно снова пробовать вход
		return true
	}
	return false
}
