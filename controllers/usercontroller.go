package usercontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"go-crud/libraries"

	"go-crud/models"

	"go-crud/entities"
)

var validation = libraries.NewValidation()
var userModel = models.NewUserModel()

func Index(response http.ResponseWriter, request *http.Request) {

	dataku, _ := userModel.FindAll()

	data := map[string]interface{}{
		"dataku": dataku,
	}

	temp, err := template.ParseFiles("views/pages/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pages/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var dataku entities.Dataku
		dataku.NamaLengkap = request.Form.Get("nama_lengkap")
		dataku.NIK = request.Form.Get("nik")
		dataku.JenisKelamin = request.Form.Get("jenis_kelamin")
		dataku.TempatLahir = request.Form.Get("tempat_lahir")
		dataku.TanggalLahir = request.Form.Get("tanggal_lahir")
		dataku.Alamat = request.Form.Get("alamat")
		dataku.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(dataku)

		if vErrors != nil {
			data["dataku"] = dataku
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil disimpan"
			userModel.Create(dataku)
		}

		temp, _ := template.ParseFiles("views/pages/add.html")
		temp.Execute(response, data)
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

		temp, err := template.ParseFiles("views/pages/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var dataku entities.Dataku
		dataku.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		dataku.NamaLengkap = request.Form.Get("nama_lengkap")
		dataku.NIK = request.Form.Get("nik")
		dataku.JenisKelamin = request.Form.Get("jenis_kelamin")
		dataku.TempatLahir = request.Form.Get("tempat_lahir")
		dataku.TanggalLahir = request.Form.Get("tanggal_lahir")
		dataku.Alamat = request.Form.Get("alamat")
		dataku.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(dataku)

		if vErrors != nil {
			data["user"] = dataku
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil diperbarui"
			userModel.Update(dataku)
		}

		temp, _ := template.ParseFiles("views/pages/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	userModel.Delete(id)

	http.Redirect(response, request, "/pages", http.StatusSeeOther)
}
