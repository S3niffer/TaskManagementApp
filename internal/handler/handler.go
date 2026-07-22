package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/s3niffer/taskmanagementapp/internal/app"
	"github.com/s3niffer/taskmanagementapp/internal/models"
)

func New(app app.Application) *chi.Mux {
	r := chi.NewRouter()

	// r.Use(AuthMiddleware)

	// r.With(AuthMiddleware).Get()

	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/me", app.UserApi.GetUserInfo)
	})

	// r.Route("/protected",func(r chi.Router) {})

	r.Get("/health", app.HealthCheck)

	r.Post("/register", app.UserApi.RegisterUser)
	r.Post("/login", app.UserApi.LoginUser)

	return r
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		if tokenString == "" || tokenString == "Bearer" {
			http.Error(w, "Your'e not logged in.", http.StatusForbidden)
			return
		}

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("JWT_SECRET"), nil
			},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userId, ok := claims["userID"]

		if ok {
			next.ServeHTTP(
				w,
				r.WithContext(
					context.WithValue(
						r.Context(),
						models.AuthMiddleUserIdKey{},
						userId,
					),
				),
			)
			return
		}

		http.Error(w, "you must to be logged in. to use this route", http.StatusForbidden)
	})
}
