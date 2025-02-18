package main

import (
	"carrent/db"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

var (
	PORT          string = "8080"
	COOKIE_LENGTH int    = 256
)

// TODO:
// 1. Регистрация пользователя 		- сделано
// 2. Вход пользователя				- сделано
// 3. Вывод всех машин 				- сделано
// 4. Вывод конкретной машины по id	- сделано
// 5. Вывод профиля пользователя	- сделано
// 6. Отправка лицензии
// 7. Проверка куки					- сделано
// 8. Генерация куки 				- сделано
// 9. CORS настройки
// 10. Установка куки				- сделано
// 11. Развернуть БД            	- сделано
// 12. Настройки приложения
// 	через конфиг
// 13. Валидация
// 14. Возможность бронирования
// 	машины
// 15. Настройка ssl-соединения с бд

func main() {
	r := chi.NewRouter()
	db.DBConnect()

	// CORS настройки
	/*
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
			Debug:            true,
		}))*/

	// Обработчики маршрутов
	r.Post("/login", Login)
	r.Post("/register", Register)
	r.Get("/cars", ListCars)
	r.Get("/cars/{id}", GetCarById)
	r.Post("/cars/{id}/rent", SessionMiddleware(RentCar))
	r.Get("/cars/{id}/track", SessionMiddleware(TrackCar))
	r.Get("/user/profile", SessionMiddleware(ViewProfile))
	r.Post("/user/license/send", SessionMiddleware(SendLicense))

	// Запуск сервера
	log.Printf("Server started at 8080")
	http.ListenAndServe(":8080", r)
}

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
	if user.ID == 0 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(credentials["password"]))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Костыль или нет, если такая кука уже есть или появляется ошибка при ее генерации 10 раз, то отправляем 500 ошибку
	// Ограничение в количество генераций сделано для избежания бесконечного цикла
	not_gen := true
	var cookieValue string
	count := 10
	for not_gen && count != 0 {
		cookieValue, err = GenerateRandomToken(COOKIE_LENGTH)
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
	user.Cookie = cookieValue
	user.SessionExpires = time.Now().Add(30 * 24 * time.Hour)
	db.DB.Save(&user)

	//Устанавливаем куку пользователю
	http.SetCookie(w, &http.Cookie{
		Name:     "cookie",
		Value:    cookieValue,
		HttpOnly: true,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		//Secure:   true,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

// Функция, которая выводит все машины
func ListCars(w http.ResponseWriter, r *http.Request) {
	var cars []db.Car
	result := db.DB.Find(&cars)
	if result.Error != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var obj_responses []map[string]interface{}
	for _, car := range cars {
		response := make(map[string]interface{})
		response["id"] = car.ID
		response["Manufacturer"] = car.Manufacturer
		response["Model"] = car.Model
		response["Year"] = car.YearOfManufacture
		response["Condition"] = car.Condition
		obj_responses = append(obj_responses, response)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(obj_responses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// Выводим конкретную машину по ее id
func GetCarById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be integer", http.StatusBadRequest)
		return
	}
	var car db.Car
	res := db.DB.Where(&db.Car{ID: id}).First(&car)
	if res.RowsAffected == 0 {
		http.Error(w, "No car by such id", http.StatusNotFound)
		return
	}

	//Выводим информацию по машине
	response := make(map[string]interface{})
	response["Manufacturer"] = car.Manufacturer
	response["Model"] = car.Model
	response["Year"] = car.YearOfManufacture
	response["Condition"] = car.Condition

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Функция для генерациия куки, возвращается в кодировке base64
func GenerateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(db.User)

	response := make(map[string]interface{})
	response["username"] = user.Login
	response["id"] = user.ID
	response["first_name"] = user.FirstName
	response["second_name"] = user.SecondName
	response["last_name"] = user.LastName
	response["phone"] = user.ContactPhone
	response["e-mail"] = user.Email
	response["gender"] = user.Gender

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendLicense(w http.ResponseWriter, r *http.Request) {

}

func RentCar(w http.ResponseWriter, r *http.Request) {

}

func TrackCar(w http.ResponseWriter, r *http.Request) {

}

func SessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("cookie")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var user db.User
		db.DB.Where(&db.User{Cookie: cookie.Value}).First(&user)
		if user.ID == 0 || time.Now().After(user.SessionExpires) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Здесь можно сохранить информацию о пользователе в контексте запроса
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
