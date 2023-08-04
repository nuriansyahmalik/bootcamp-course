package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
)

type Claims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret")

//func ValidateJWTMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		tokenString := r.Header.Get("Authorization")[(len("Bearer ")):]
//		if tokenString == "" {
//			log.Println("no header")
//			http.Error(w, "Unauthorized", http.StatusUnauthorized)
//		}
//
//		ctx := context.WithValue(r.Context(), "claims", claims)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}

func ValidateToken(w http.ResponseWriter, r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")[(len("Bearer ")):]
	if tokenString == "" {
		log.Println("no header")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	return tokenString, nil
}

//	func CheckRoleMiddleware(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			_, ok := r.Context().Value("claims").(*jwt.Claims)
//			if !ok {
//				http.Error(w, "Error Claims", http.StatusUnauthorized)
//				return
//			}
//
//			//if claims. != "teacher" {
//			//	http.Error(w, "Unauthorized", http.StatusUnauthorized)
//			//	return
//			}
//			ctx := context.WithValue(r.Context(), "claims", claims)
//			next.ServeHTTP(w, r.WithContext(ctx))
//		})
//	}

//func ValidateToken(tokenString string) (*Claims, error) {
//
//	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtKey, nil
//	})
//
//	if err != nil {
//		return nil, fmt.Errorf("JWT validation failed: %v", err)
//	}
//
//	if claims, ok := token.Claims.(*Claims); ok && token.Valid && claims.Role == "teacher" {
//		if claims.Role != "teacher" {
//			return nil, fmt.Errorf("not a role teacher %s", err)
//		}
//		return claims, nil
//	}
//	return nil, fmt.Errorf("JWT is not valid")
//}
