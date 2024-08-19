package web

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"path/filepath"

// 	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
// 	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
// 	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
// )

// // RegisterUserHandler handles HTTP requests to register a new user.
// func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("ParseForm() err: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	userName := r.FormValue("username")
// 	password := r.FormValue("password")

// 	if userName == "" || password == "" {
// 		http.Error(w, "Username and password are required", http.StatusBadRequest)
// 		return
// 	}

// 	configPath := "./connection-profile.yaml"
// 	sdk, err := fabsdk.New(config.FromFile(configPath))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create new SDK: %s", err), http.StatusInternalServerError)
// 		return
// 	}
// 	defer sdk.Close()

// 	mspClient, err := msp.New(sdk.Context())
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create MSP client: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	registrationRequest := &msp.RegistrationRequest{
// 		Name:        userName,
// 		Type:        "client",
// 		Affiliation: "org1.department1",
// 		Secret:      password,
// 		Attributes: []msp.Attribute{
// 			{Name: "abac.creator", Value: "true", ECert: true}, // Add the ABAC attribute
// 		},
// 	}

// 	secret, err := mspClient.Register(registrationRequest)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to register user: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "Successfully registered user, secret: %s", secret)
// }

// // EnrollUserHandler handles HTTP requests to enroll an existing user.
// func EnrollUserHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("ParseForm() err: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	userName := r.FormValue("username")
// 	password := r.FormValue("password")

// 	if userName == "" || password == "" {
// 		http.Error(w, "Username and password are required", http.StatusBadRequest)
// 		return
// 	}

// 	configPath := "./connection-profile.yaml"
// 	sdk, err := fabsdk.New(config.FromFile(configPath))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create new SDK: %s", err), http.StatusInternalServerError)
// 		return
// 	}
// 	defer sdk.Close()

// 	mspClient, err := msp.New(sdk.Context())
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create MSP client: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	err = mspClient.Enroll(userName, msp.WithSecret(password))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to enroll user: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	userID, err := mspClient.GetSigningIdentity(userName)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to retrieve user identity: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	userMSPDir := filepath.Join("../../test-network/organizations/peerOrganizations/org1.example.com/users", fmt.Sprintf("%s@org1.example.com/msp", userName))
// 	err = os.MkdirAll(userMSPDir, 0755)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create MSP directory: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	certPath := filepath.Join(userMSPDir, "signcerts")
// 	err = os.MkdirAll(certPath, 0755)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create signcerts directory: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	err = ioutil.WriteFile(filepath.Join(certPath, "cert.pem"), userID.EnrollmentCertificate(), 0644)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to write certificate: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	keyPath := filepath.Join(userMSPDir, "keystore")
// 	err = os.MkdirAll(keyPath, 0755)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create keystore directory: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	keyPEM, err := userID.PrivateKey().Bytes()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to retrieve private key bytes: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	err = ioutil.WriteFile(filepath.Join(keyPath, "key.pem"), keyPEM, 0600)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to write private key: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	srcCACert := "../../test-network/organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem"
// 	destCACert := filepath.Join(userMSPDir, "cacerts", "ca-cert.pem")
// 	err = copyFile(srcCACert, destCACert)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to copy CA certificate: %s", err), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "User credentials successfully saved to: %s", userMSPDir)
// }

// // Helper function to copy a file from src to dst
// func copyFile(src, dst string) error {
// 	sourceFile, err := os.Open(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer sourceFile.Close()

// 	destDir := filepath.Dir(dst)
// 	err = os.MkdirAll(destDir, 0755)
// 	if err != nil {
// 		return err
// 	}

// 	destFile, err := os.Create(dst)
// 	if err != nil {
// 		return err
// 	}
// 	defer destFile.Close()

// 	_, err = io.Copy(destFile, sourceFile)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
