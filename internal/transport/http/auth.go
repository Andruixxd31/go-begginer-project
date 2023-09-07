package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

func JWTAuth(
    original func(w http.ResponseWriter, r *http.Request), 
) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header["Authorization"]
        if authHeader == nil {
            http.Error(w, "not Authorized", http.StatusUnauthorized)
            return
        }

        //Bearer: token-string
        authHeaderParts := strings.Split(authHeader[0], " ")
        if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
            http.Error(w, "not Authorized", http.StatusUnauthorized)
            return        
        }

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			log.Error("could not validate incoming token")
			return
		}
    }
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("elllanoenllamas")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}   
