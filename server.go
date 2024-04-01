package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    fmt.Printf("Sever listening at port 8080\n")

    // ? Create a home page handler
    http.HandleFunc("/home", homePageHandler)
    // ? Create a hello handler
    http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/login", loginHandler)

    http.HandleFunc("/auth", authHandler)

    // ? Start the server, if any error, console log it
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }

}

// ? Handler for the hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Page!")
}

// ? Handler for the home route
func homePageHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")	
}

func loginHandler(w http.ResponseWriter, r*http.Request) {
	http.ServeFile(w,r, "./static/form.html")
}

func authHandler(w http.ResponseWriter, r*http.Request) {
    fmt.Println("Authenticating...")
}