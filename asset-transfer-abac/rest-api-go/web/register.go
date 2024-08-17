package web

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
// 	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
// 	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
// )

// // UserRegistrationRequest defines the structure of the request body for user registration.
// type UserRegistrationRequest struct {
// 	Username    string `json:"username"`
// 	Affiliation string `json:"affiliation"`
// }

// // RegisterUserHandler handles POST requests to register a new user.
// func (setup OrgSetup) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Parse the request body to get the user details.
// 	var reqBody UserRegistrationRequest
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Failed to read request body", http.StatusBadRequest)
// 		return
// 	}
// 	if err := json.Unmarshal(body, &reqBody); err != nil {
// 		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Validate input
// 	if reqBody.Username == "" {
// 		http.Error(w, "Username is required", http.StatusBadRequest)
// 		return
// 	}

// 	// Path to your network configuration file.
// 	configPath := "./network-config.yaml"

// 	// Create a new Fabric SDK instance.
// 	sdk, err := fabsdk.New(config.FromFile(configPath))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create SDK: %s", err), http.StatusInternalServerError)
// 		return
// 	}
// 	defer sdk.Close()

// 	// Create an MSP client.
// 	mspClient, err := msp.New(sdk.Context())
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create MSP client: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Register the new user.
// 	registrationRequest := &msp.RegistrationRequest{
// 		Name:        reqBody.Username,    // Name of the new user
// 		Type:        "client",            // Type of the identity
// 		Affiliation: reqBody.Affiliation, // Affiliation, e.g., "org1.department1"
// 		Attributes: []msp.Attribute{
// 			{Name: "role", Value: "client", ECert: true},
// 		},
// 	}

// 	secret, err := mspClient.Register(registrationRequest)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to register user: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Return the secret to the client.
// 	response := map[string]string{"message": "User registered successfully", "secret": secret}
// 	jsonResponse(w, response, http.StatusOK)
// }

// // jsonResponse utility function for sending JSON responses.
// func jsonResponse(w http.ResponseWriter, payload interface{}, status int) {
// 	response, err := json.Marshal(payload)
// 	if err != nil {
// 		http.Error(w, "Failed to generate JSON response", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	w.Write(response)
// }
