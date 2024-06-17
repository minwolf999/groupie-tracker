package errors

import (
	"fmt"
	"net/http"
)

// The check function takes http.ResponseWriter and analyse the inputed error
func Check(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println(err)
		Error(w, http.StatusInternalServerError)
		return
	}
}
