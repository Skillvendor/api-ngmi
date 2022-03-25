package animal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Human struct {
	Name string
	Age  int
}

func MyFunc(w http.ResponseWriter, r *http.Request) {
	myAnimal := Human{Name: "Gigel", Age: 25}

	err := json.NewEncoder(w).Encode(myAnimal)

	if err != nil {
		fmt.Println("ERROR2", err)
	}
	// w.Write([]byte(myAnimal.Owner[0].Name))
}
