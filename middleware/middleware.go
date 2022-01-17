package middleware

import (
	"fmt"
	"net/http"
)

// TODO may not needed
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func AdminOnly(next http.Handler) http.Handler {
	fmt.Println("Admin verification")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})

}

// todo check Token
//func CheckToken(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Add("Vary", "Authorization")
//		authHeader := r.Header.Get("Authorization")
//
//		if authHeader == "" {
//
//		}
//
//		headerParts := strings.Split(authHeader, " ")
//		if len(headerParts) != 2 {
//			utils.ErrorJSON(w, errors.New("invalid auth header"))
//			return
//		}
//
//		if headerParts[0] != "Bearer" {
//			utils.ErrorJSON(w, errors.New("unauthorized - no bearer"))
//			return
//		}
//
//		tokenString := headerParts[1]
//
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//				utils.ErrorJSON(w, errors.New("unexpected signing method"))
//			}
//
//			//return secretmac
//
//		})
//		next.ServeHTTP(w, r)
//	})
//}
