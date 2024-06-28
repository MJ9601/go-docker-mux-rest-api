package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func WithJWTAuth(handleFunc http.HandlerFunc, store Store) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
	 tokenString := GetTokenFromRequest(r)
	 token, err := validateJwt(tokenString)
	 
	 if err != nil {
		permissionDenied(w, err, "Failed to authenticate token")
		return 
	 }

	 if !token.Valid {
		permissionDenied(w, err, "Failed to authenticate token")
		return
	 }

	 claims := token.Claims.(jwt.MapClaims)
	 userID := claims["userID"].(string)

	 _, err = store.GetUserByID(userID)
	 if err != nil {
		permissionDenied(w, err, "Failed to get the User")
		return
	 }
	
	}
}

func GetTokenFromRequest (r *http.Request) (string) {

	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != ""{
		return tokenAuth
	}
	if tokenQuery != ""{
		return tokenQuery
	}

	return ""
}

func permissionDenied(w http.ResponseWriter, err error, message string){
		log.Printf("%s: %v",message, err)
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{Error: fmt.Errorf("permission Denied").Error()})
}

func validateJwt (token string) (*jwt.Token, error) {

	secret := Envs.JWTSecret

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
 	})
}