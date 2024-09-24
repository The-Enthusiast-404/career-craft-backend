package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

var keycloakCertsURL = "http://localhost:8080/realms/career-craft/protocol/openid-connect/certs"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token := bearerToken[1]

		// Fetch JWK Set from Keycloak
		set, err := jwk.Fetch(r.Context(), keycloakCertsURL)
		if err != nil {
			http.Error(w, "Failed to fetch JWK Set", http.StatusInternalServerError)
			return
		}

		// Parse and validate the token
		parsedToken, err := jwt.Parse(
			[]byte(token),
			jwt.WithKeySet(set),
			jwt.WithValidate(true),
		)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, err := parsedToken.AsMap(r.Context())
		if err != nil {
			http.Error(w, "Failed to extract claims", http.StatusInternalServerError)
			return
		}

		// Add claims to request context
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}