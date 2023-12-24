package usercontroller

import (
	// "html/template"
	"encoding/json"
	// "fmt"
	"go-crud/libraries"
	"net/http"
	"strconv"

	"go-crud/models"

	"go-crud/entities"
)

var validation = libraries.NewValidation()
var userModel = models.NewUserModel()

func Index(response http.ResponseWriter, request *http.Request) {
	// Mengambil data dari userModel
	dataku, _ := userModel.FindAll()

	// Membuat map untuk data JSON
	jsonData := map[string]interface{}{
		"dataku": dataku,
	}

	// Mengonversi map menjadi JSON
	jsonResponse, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Mengatur header Content-Type sebagai application/json
	response.Header().Set("Content-Type", "application/json")

	// Menulis respon JSON ke http.ResponseWriter
	response.Write(jsonResponse)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		// Parse JSON from request body
		var dataku entities.Dataku
		err := json.NewDecoder(request.Body).Decode(&dataku)
		if err != nil {
			http.Error(response, "Invalid JSON", http.StatusBadRequest)
			return
		}

		var data = make(map[string]interface{})
		vErrors := validation.Struct(dataku)

		if vErrors != nil {
			data["dataku"] = dataku
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil disimpan"
			userModel.Create(dataku)
		}

		// Convert data to JSON
		jsonResponse, err := json.Marshal(data)
		if err != nil {
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set response headers
		response.Header().Set("Content-Type", "application/json")

		// Write JSON response to http.ResponseWriter
		response.Write(jsonResponse)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var user entities.Dataku
		userModel.Find(id, &user)

		data := map[string]interface{}{
			"user": user,
		}

		// Render JSON response
		renderJSON(response, data)
	} else if request.Method == http.MethodPut {
		var dataku entities.Dataku

		decoder := json.NewDecoder(request.Body)
		if err := decoder.Decode(&dataku); err != nil {
			http.Error(response, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		var data = make(map[string]interface{})

		vErrors := validation.Struct(dataku)

		if vErrors != nil {
			data["user"] = dataku
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil diperbarui"
			userModel.Update(dataku)
		}

		// Render JSON response
		renderJSON(response, data)
	}
}

// Helper function to render JSON response
func renderJSON(response http.ResponseWriter, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(data)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	// Mendapatkan ID dari parameter URL
	queryString := request.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)
	// fmt.Printf(err.Error())
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// Menghapus data dengan menggunakan UserModel
	err = userModel.Delete(id)
	if err != nil {
		// Jika terjadi kesalahan saat menghapus
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Mengembalikan respon JSON
	res := map[string]interface{}{
		"status":  "success",
		"message": "Data berhasil dihapus",
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(res)
}
