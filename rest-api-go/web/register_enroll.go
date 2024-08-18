package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// UserRequest defines the structure for the JSON payload for user operations
type UserRequest struct {
	Name        string `json:"name"`
	Affiliation string `json:"affiliation"`
	Secret      string `json:"secret"`
}

// Initialize initializes the OrgSetup structure and MSPClient
func Init(config OrgSetup) (*OrgSetup, error) {
	// Setup Fabric SDK to handle MSP operations like register and enroll
	sdk, err := fabsdk.New(config.FromFile("./connection-profile.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to create Fabric SDK: %v", err)
	}
	defer sdk.Close()

	// Create the MSP client
	mspClient, err := msp.New(sdk.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to create MSP client: %v", err)
	}

	config.MSPClient = mspClient

	return &config, nil
}

// RegisterUserHandler handles user registration.
func (setup *OrgSetup) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	registrationRequest := &msp.RegistrationRequest{
		Name:        req.Name,
		Type:        "client",
		Affiliation: req.Affiliation,
		Secret:      req.Secret,
		Attributes: []msp.Attribute{
			{Name: "abac.creator", Value: "true", ECert: true},
		},
	}

	secret, err := setup.MSPClient.Register(registrationRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to register user: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User registered successfully. Secret: %s", secret)
}

// EnrollUserHandler handles user enrollment.
func (setup *OrgSetup) EnrollUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := setup.MSPClient.Enroll(req.Name, msp.WithSecret(req.Secret))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to enroll user: %s", err), http.StatusInternalServerError)
		return
	}

	// Store the user's identity details
	userID, err := setup.MSPClient.GetSigningIdentity(req.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get signing identity: %s", err), http.StatusInternalServerError)
		return
	}

	// Define the user MSP directory
	userMSPDir := filepath.Join("../../test-network/organizations/peerOrganizations/org1.example.com/users", req.Name+"@org1.example.com/msp")

	// Ensure the MSP directory exists
	err = os.MkdirAll(userMSPDir, 0755)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create MSP directory: %s", err), http.StatusInternalServerError)
		return
	}

	// Store the user's certificate
	certPath := filepath.Join(userMSPDir, "signcerts")
	err = os.MkdirAll(certPath, 0755)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create signcerts directory: %s", err), http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(filepath.Join(certPath, "cert.pem"), userID.EnrollmentCertificate(), 0644)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write certificate: %s", err), http.StatusInternalServerError)
		return
	}

	// Store the user's private key
	keyPath := filepath.Join(userMSPDir, "keystore")
	err = os.MkdirAll(keyPath, 0755)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create keystore directory: %s", err), http.StatusInternalServerError)
		return
	}

	keyPEM, err := userID.PrivateKey().Bytes()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve private key: %s", err), http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile(filepath.Join(keyPath, "key.pem"), keyPEM, 0600)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write private key: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User enrolled successfully. Credentials saved to: %s", userMSPDir)
}
