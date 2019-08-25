package router

import (
	"mux"
	"assignment/3/pkg/controller"
)

var Mux = mux.NewRouter()
// var mux = http.NewServeMux()

func Routers() {

	// mux.Handle("/", http.HandlerFunc(controller.index))
	Mux.HandleFunc("/", controller.Index)
	Mux.HandleFunc("/create", controller.Add)
	Mux.HandleFunc("/update", controller.Update)
	Mux.HandleFunc("/delete", controller.Delete)
	Mux.HandleFunc("/getbyid", controller.GetById)
	Mux.HandleFunc("/getall", controller.GetAll)
}
