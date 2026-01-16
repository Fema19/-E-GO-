package main

import (
	"log"
	"net/http"

	"backend-event-api/internal/database"
	"backend-event-api/internal/handler"
	"backend-event-api/internal/middleware"
	"backend-event-api/internal/repository"
	"backend-event-api/internal/service"
)

func main() {
	db := database.Connect()

	userRepo := repository.NewUserRepository(db)
	eventRepo := repository.NewEventRepository(db)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)
	http.Handle("/events", middleware.JWTAuth(http.HandlerFunc(eventHandler.Create)))

	http.Handle("/me",
		middleware.JWTAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			email := middleware.GetEmail(r)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"email":"` + email + `"}`))
		})),
	)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
