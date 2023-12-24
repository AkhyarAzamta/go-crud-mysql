package main

import (
	"net/http"

	"go-crud/controllers"
)

func main() {

	http.HandleFunc("/", usercontroller.Index)
	http.HandleFunc("/pages", usercontroller.Index)
	http.HandleFunc("/pages/index", usercontroller.Index)
	http.HandleFunc("/pages/add", usercontroller.Add)
	http.HandleFunc("/pages/edit", usercontroller.Edit)
	http.HandleFunc("/pages/delete", usercontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
