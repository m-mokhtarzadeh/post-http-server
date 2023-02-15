package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
                file, err := os.OpenFile("test.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                if err != nil {
                    log.Fatalf("failed creating file: %s", err)
                }
	        defer file.Close()
		b := fmt.Sprintf("Name = %s, Address = %s\n", name, address )
	        _, err = file.WriteString(b)
                if err != nil {
                    log.Fatalf("failed writing to file: %s", err)
                }
		defer file.Close()
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

        server_port := os.Getenv("SERVER_PORT")
        server_listen := ":" + server_port
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(server_listen, nil); err != nil {
		log.Fatal(err)
	}
}
