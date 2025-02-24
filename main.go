package main

import (
	"carrent/carmanage"
	"carrent/db"
	"carrent/users"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/gorilla/csrf"
	"github.com/hirochachacha/go-smb2"
	_ "github.com/wagslane/go-password-validator"
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

	CSRF := csrf.Protect(
		[]byte("a-32-byte-long-key-goes-here"),
		csrf.CookieName("X-CSRF-Token"),
		csrf.Secure(false), // Set to true in production
		csrf.Path("/auth"),
		//csrf.ErrorHandler(http.HandlerFunc(serverError(403))),
	)
	// Обработчики маршрутов
	r.Post("/auth/login", users.Login)
	r.Get("/auth/login", users.GetCSRFToken)
	r.Post("/auth/register", users.Register)
	r.Get("/auth/register", users.GetCSRFToken)
	r.Get("/cars", carmanage.ListCars)
	r.Get("/cars/{id}", carmanage.GetCarById)
	r.Post("/cars/{id}/rent", SessionMiddleware(carmanage.RentCar))
	r.Get("/cars/{id}/track", SessionMiddleware(carmanage.TrackCar))
	r.Get("/user/profile", SessionMiddleware(users.ViewProfile))
	r.Post("/user/license/send", SessionMiddleware(users.SendLicense))

	// Запуск сервера
	log.Printf("Server started at 8080")
	http.ListenAndServe(":8080", CSRF(r))
}

func SessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
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

func isValidFileType(file []byte) bool {
	fileType := http.DetectContentType(file)
	return strings.HasPrefix(fileType, "image/") // Only allow images
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Ограничиваем размер до 10 MB

	// Получаем файл из формы
	file, handler, err := r.FormFile("license")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Пишем в логи информацию о файле
	fmt.Fprintf(w, "Uploaded File: %s\n", handler.Filename)
	fmt.Fprintf(w, "File Size: %d\n", handler.Size)
	fmt.Fprintf(w, "MIME Header: %v\n", handler.Header)

	// Проверяем тип файла
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	if !isValidFileType(fileBytes) {
		http.Error(w, "Invalid file type", http.StatusUnsupportedMediaType)
		return
	}

	// Сохраняем
	dst, err := SaveFile(handler.Filename)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Proceed with saving the file
	if _, err := dst.Write(fileBytes); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
	}
}

func SaveFile(name string) (*os.File, error) {
	dst, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func saveToSmbServer(filename string, file io.Reader) error {
	// Подключение к SMB серверу
	conn, err := net.Dial("tcp", "SERVERNAME:445")

	_ = filename
	_ = file

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "USERNAME",
			Password: "PASSWORD",
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		fmt.Printf("Error occured while connecting to file server: %v", err)
	}
	defer s.Logoff()

	fs, err := s.Mount("CARSHARE")
	if err != nil {
		fmt.Printf("Error occured while mounting: %v", err)
	}
	defer fs.Umount()

	return nil
}
