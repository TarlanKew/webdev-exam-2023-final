package main

import (
	"html/template"
	"log"
	"map/data"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	u := data.Routes{}
	u.GetAll()
	t, _ := template.ParseFiles("LK/main.html")
	t.Execute(w, u)
}

func about(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("LK/about.html")
	t.Execute(w, nil)
}

func maps(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("LK/map.html")
	t.Execute(w, nil)
}

func route(w http.ResponseWriter, r *http.Request) {
	coordinations := strings.Split(r.URL.Query().Get("data"), "|")
	u := data.Location{
		Start: coordinations[0],
		Stop:  coordinations[1],
	}

	t, _ := template.ParseFiles("LK/route.html")
	t.Execute(w, u)
}

// api func
func getRoutests(w http.ResponseWriter, r *http.Request) {
	// array *Struct
}

func getGuides(w http.ResponseWriter, r *http.Request) {

}

func getOrders(w http.ResponseWriter, r *http.Request) {

}

func addOrders(w http.ResponseWriter, r *http.Request) {

}

func putOrders(w http.ResponseWriter, r *http.Request) {

}

func deleteOrders(w http.ResponseWriter, r *http.Request) {

}

func getOrderInfo(w http.ResponseWriter, r *http.Request) {

}

func getGuideInfo(w http.ResponseWriter, r *http.Request) {

}

func run() {
	http.Handle("/style/",
		http.StripPrefix("/style/",
			http.FileServer(http.Dir("./style/"))))
	http.Handle("/LK/",
		http.StripPrefix("/LK/",
			http.FileServer(http.Dir("./LK/"))))

	rout := mux.NewRouter()
	rout.HandleFunc("/", homePage).Methods("GET")
	rout.HandleFunc("/about/", about).Methods("GET")
	rout.HandleFunc("/map/", maps).Methods("GET")
	rout.HandleFunc("/route/", route).Methods("GET")

	// api
	rout.HandleFunc("/api/routes/", getRoutests).Methods("GET")
	rout.HandleFunc("/api/routes/{id:[0-9]+}/guides/", getGuides).Methods("GET")
	rout.HandleFunc("/api/orders/", getOrders).Methods("GET")
	rout.HandleFunc("/api/orders/", addOrders).Methods("POST")
	rout.HandleFunc("/api/orders/{id:[0-9]+}", putOrders).Methods("PUT")
	rout.HandleFunc("/api/orders/{id:[0-9]+}", deleteOrders).Methods("DELETE")
	rout.HandleFunc("/api/orders/{id:[0-9]+}", getOrderInfo).Methods("GET")
	rout.HandleFunc("/api/guides/{id:[0-9]+}/", getGuideInfo).Methods("GET")

	http.Handle("/", rout) // перенаправление на роутер
	http.ListenAndServe(":8080", nil)
}

func main() {
	log.Println("SERVER START")
	run()
}
