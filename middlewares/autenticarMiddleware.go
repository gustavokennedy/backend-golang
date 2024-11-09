package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AutenticarMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Token de autenticação não fornecido", http.StatusUnauthorized)
			return
		}

		// Verificar se o cabeçalho começa com "Bearer"
		// if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		// 	http.Error(w, "Token de autenticação inválido", http.StatusUnauthorized)
		// 	return
		// }

		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verificar o algoritmo de assinatura do token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			// Substitua "chave-secreta" pela sua chave secreta real usada para assinar o token JWT
			return []byte("b5f1e44e6939e09d3e3ab0845ad72d5e7110a5f6b63c54782b78e1ab1cf8c54d3"), nil
		})

		if err != nil {
			log.Println(err)
			http.Error(w, "Token de autenticação inválido", http.StatusUnauthorized)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Token de autenticação inválido", http.StatusUnauthorized)
			return
		}
	})
}
