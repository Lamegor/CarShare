package users

import (
	"carrent/db"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"golang.org/x/crypto/bcrypt"
)

// Вход зарегистрированного пользователя, возвращает cookie пользователю
func Login(w http.ResponseWriter, r *http.Request) {

	var credentials map[string]string
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil || credentials["username"] == "" || credentials["password"] == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var user db.User
	db.DB.Where(&db.User{Login: credentials["username"]}).First(&user)
	if user.ID == 0 || !checkLoginAttempts(user) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(credentials["password"]))
	if err != nil {
		user.FailedAttemptsInARow++
		user.LastFailedAttempt = time.Now()
		db.DB.Save(&user)
		if user.FailedAttemptsInARow >= maxLoginAttempts {
			http.Error(w, "Too many failed attempts", http.StatusTooManyRequests)
			return
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
	}

	// Костыль или нет, если такая кука уже есть или появляется ошибка при ее генерации 10 раз, то отправляем 500 ошибку
	// Ограничение в количество генераций сделано для избежания бесконечного цикла
	not_gen := true
	var cookieValue string
	count := 10
	for not_gen && count != 0 {
		cookieValue, err = GenerateRandomToken()
		if res := db.DB.Where(&db.User{Cookie: cookieValue}); err == nil && res.RowsAffected == 0 {
			not_gen = false
			break
		}
		count--
	}
	if not_gen == true {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Login successful", "user": user.Login}
	//Сохраняем куку в базе данных
	user.FailedAttemptsInARow = 0
	user.LastFailedAttempt = time.Time{}
	user.Cookie = cookieValue
	user.SessionExpires = time.Now().Add(30 * 24 * time.Hour)
	db.DB.Save(&user)

	//Устанавливаем куку пользователю
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    cookieValue,
		HttpOnly: true,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		//Secure:   true,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

func GetCSRFToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	json.NewEncoder(w).Encode(nil)
	return
}

// Создание нового пользователя, перенаправляет на страницу логина
func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data["username"] == "" || data["1_password"] == "" || data["2_password"] == "" || data["email"] == "" {
		http.Error(w, "One of Parametres is empty", http.StatusBadRequest)
		return
	} else if data["1_password"] != data["2_password"] {
		http.Error(w, "Password missmatch", http.StatusBadRequest)
		return
	}

	var usr db.User
	result := db.DB.Where(&db.User{Login: data["username"]}).First(&usr)
	if result.RowsAffected > 0 {
		http.Error(w, "There is such user with such username", http.StatusConflict)
		return
	}
	result = db.DB.Where(&db.User{Email: data["email"]}).First(&usr)
	if result.RowsAffected > 0 {
		http.Error(w, "There is user with such e-mail", http.StatusConflict)
		return
	}

	hash_pass, err := bcrypt.GenerateFromPassword([]byte(data["1_password"]), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	user := db.User{Login: data["username"], Password: hash_pass}
	db.DB.Create(&user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(nil)
}
