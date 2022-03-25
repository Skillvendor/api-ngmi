package animal

import (
	"encoding/json"
	"log"
	"net/http"
)

type Human struct {
	Name string
	Age  int
}

func MyFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling animal")
	myAnimal := Human{Name: "Gigel", Age: 25}

	err := json.NewEncoder(w).Encode(myAnimal)

	if err != nil {
		log.Println("Error loading .env file", err)
	}
	// w.Write([]byte(myAnimal.Owner[0].Name))
}
