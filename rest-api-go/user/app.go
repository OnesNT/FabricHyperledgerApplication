package user

import (
	"fmt"
	"net/http"
)

// Serve starts the HTTP web server and maps routes for user registration and enrollment.
func Serve() {
	http.HandleFunc("/register", RegisterUserHandler) // Route for user registration
	http.HandleFunc("/enroll", EnrollUserHandler)     // Route for user enrollment

	fmt.Println("Listening on http://localhost:3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
