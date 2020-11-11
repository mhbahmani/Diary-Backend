package auth

import (
	"fmt"
	"crypto/rand"
	"encoding/json"
	"net/http"
)

func tokenGenerator() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func authentication(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqData map[string]interface{}
		json.NewDecoder(r.Body).Decode(&reqData)
		// if Authentication failed {
		// 	http.Error(w, "Forbiden", 403)
		// 	return
		// }
		h.ServeHTTP(w, r)
	}
}