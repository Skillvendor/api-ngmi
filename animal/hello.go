package animal

import (
	"net/http"
)

func MyFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("YES"))
}
