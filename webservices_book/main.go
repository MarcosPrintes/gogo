package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	name := urlParams["user"]
	helloMsg := "Hello " + name

	message := API{helloMsg}

	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Algo errado aqui na hello migão!")
	}

	fmt.Fprint(w, string(output))
}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", hello)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8083", nil)
}

/*
func main() {
		There are some limitations with the built-in mux/router provided by the net/http package. You cannot, for example, supply a wildcard or a regular expression to a route.
		http.HandleFunc("/api", handledFunc)
		http.ListenAndServe(":8080", nil)

		/*
		Gorilla allows us to use regular expressions on routes, and we will only get what we expect—digit-based request parameters. which are as follows:
		gorillaRoute := mux.NewRouter()
		gorillaRoute.HandleFunc("/prduct/{user:\+d}", productHandler)

	}

	func handledFunc(w http.ResponseWriter, r *http.Request) {

		// initialize API struct
		message := API{"Helooo pela 1000000 vez"}

		//marshal this to a JSON byte array, output, afeter sending this message to our iowriter class (in this case, an http.ResponseWriter value)
		output, err := json.Marshal(message)

		if err != nil {
			fmt.Println("Algo errado, miguxo!")
		}

		//
		fmt.Fprint(w, string(output))
	}
*/
