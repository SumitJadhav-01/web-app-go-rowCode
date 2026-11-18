package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "web-app/home.html")
}

func servicesPage(w http.ResponseWriter, r *http.Request) {
	// Render the services html page
	http.ServeFile(w, r, "web-app/services.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "web-app/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "web-app/contact.html")
}

func main() {

	http.HandleFunc("/home", homePage)
	http.HandleFunc("/services", servicesPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
