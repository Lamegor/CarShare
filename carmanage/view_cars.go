package carmanage

import (
	"carrent/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

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

func RentCar(w http.ResponseWriter, r *http.Request) {

}

func TrackCar(w http.ResponseWriter, r *http.Request) {

}
