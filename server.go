package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		customHandlerPort = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", helloHandler)
	fmt.Println("Go server Listening on: ", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		w.Write([]byte(`{"message": "hello world"}`))
	} else {

		// To read form data
		// err := r.ParseForm()
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		// decode JSON - To remove spaces basically
		var jsonData map[string]interface{}
		json.Unmarshal([]byte(string(body)), &jsonData)

		// encode JSON
		rjson, _ := json.Marshal(jsonData)

		// Just returning the same body as we got in POST
		w.Write([]byte(rjson))

		// http: superfluous response.WriteHeader call from main.helloHandler
		// w.WriteHeader(http.StatusOK)
	}
}
