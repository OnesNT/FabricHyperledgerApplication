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

// // UserEnrollmentRequest defines the structure of the request body for user enrollment.
// type UserEnrollmentRequest struct {
// 	Username string `json:"username"`
// 	Secret   string `json:"secret"`
// }

// // EnrollUserHandler handles POST requests to enroll a user.
// func (setup OrgSetup) EnrollUserHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Parse the request body to get the user details.
// 	var reqBody UserEnrollmentRequest
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
// 	if reqBody.Username == "" || reqBody.Secret == "" {
// 		http.Error(w, "Username and secret are required", http.StatusBadRequest)
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

// 	// Enroll the user with the secret received from registration.
// 	err = mspClient.Enroll(reqBody.Username, msp.WithSecret(reqBody.Secret))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to enroll user: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Respond with a success message.
// 	response := map[string]string{"message": "Successfully enrolled user", "username": reqBody.Username}
// 	jsonResponse(w, response, http.StatusOK)
// }
