package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {
	// Path to your network configuration file (connection profile)
	configPath := "./connection-profile.yaml" // Ensure this points to your connection profile

	// Create a new Fabric SDK instance
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s\n", err)
		return
	}
	defer sdk.Close()

	// Create an MSP client for Org1 using the admin context
	mspClient, err := msp.New(sdk.Context())
	if err != nil {
		fmt.Printf("Failed to create MSP client: %s\n", err)
		return
	}

	// Register a new user (creator4) with attributes and a password
	registrationRequest := &msp.RegistrationRequest{
		Name:        "creator6",         // Name of the new user
		Type:        "client",           // Type of the identity
		Affiliation: "org1.department1", // Affiliation within Org1
		Secret:      "creator6pw",       // Specify the password here
		Attributes: []msp.Attribute{
			{Name: "abac.creator", Value: "true", ECert: true}, // Add the ABAC attribute
		},
	}

	// Register the user using the admin identity
	secret, err := mspClient.Register(registrationRequest)
	if err != nil {
		fmt.Printf("Failed to register user: %s\n", err)
		return
	}
	fmt.Printf("Successfully registered user, secret: %s\n", secret)

	// Enroll the user (creator4) with the specified password (secret)
	err = mspClient.Enroll("creator6", msp.WithSecret(secret))
	if err != nil {
		fmt.Printf("Failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("Successfully enrolled user")

	// Retrieve the user's identity (certificate and private key)
	userID, err := mspClient.GetSigningIdentity("creator6")
	if err != nil {
		fmt.Printf("Failed to retrieve user identity: %s\n", err)
		return
	}

	// Define the user MSP directory
	userMSPDir := filepath.Join("../../test-network/organizations/peerOrganizations/org1.example.com/users", "creator6@org1.example.com/msp")

	// Ensure the MSP directory exists
	err = os.MkdirAll(userMSPDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create MSP directory: %s\n", err)
		return
	}

	// Store the user's certificate
	certPath := filepath.Join(userMSPDir, "signcerts")
	err = os.MkdirAll(certPath, 0755)
	if err != nil {
		fmt.Printf("Failed to create signcerts directory: %s\n", err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(certPath, "cert.pem"), userID.EnrollmentCertificate(), 0644)
	if err != nil {
		fmt.Printf("Failed to write certificate: %s\n", err)
		return
	}

	// Store the user's private key
	keyPath := filepath.Join(userMSPDir, "keystore")
	err = os.MkdirAll(keyPath, 0755)
	if err != nil {
		fmt.Printf("Failed to create keystore directory: %s\n", err)
		return
	}

	keyPEM, err := userID.PrivateKey().Bytes()
	if err != nil {
		fmt.Printf("Failed to retrieve private key bytes: %s\n", err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(keyPath, "key.pem"), keyPEM, 0600)
	if err != nil {
		fmt.Printf("Failed to write private key: %s\n", err)
		return
	}

	// Copy CA certificate to cacerts directory
	srcCACert := "/Users/ngokuang/golang/src/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem"
	destCACert := filepath.Join(userMSPDir, "cacerts", "localhost-7054-ca-org1.pem")
	err = copyFile(srcCACert, destCACert)
	if err != nil {
		fmt.Printf("Failed to copy CA certificate: %s\n", err)
		return
	}

	// Copy IssuerPublicKey, IssuerRevocationPublicKey, and config.yaml from Orderer's MSP directory
	srcOrdererMSPDir := "/Users/ngokuang/golang/src/fabric-samples/test-network/organizations/ordererOrganizations/example.com/msp"
	destMSPDir := userMSPDir

	// Copy IssuerPublicKey
	err = copyFile(filepath.Join(srcOrdererMSPDir, "IssuerPublicKey"), filepath.Join(destMSPDir, "IssuerPublicKey"))
	if err != nil {
		fmt.Printf("Failed to copy IssuerPublicKey: %s\n", err)
		return
	}

	// Copy IssuerRevocationPublicKey
	err = copyFile(filepath.Join(srcOrdererMSPDir, "IssuerRevocationPublicKey"), filepath.Join(destMSPDir, "IssuerRevocationPublicKey"))
	if err != nil {
		fmt.Printf("Failed to copy IssuerRevocationPublicKey: %s\n", err)
		return
	}

	// Copy config.yaml
	err = copyFile(filepath.Join(srcOrdererMSPDir, "config.yaml"), filepath.Join(destMSPDir, "config.yaml"))
	if err != nil {
		fmt.Printf("Failed to copy config.yaml: %s\n", err)
		return
	}

	fmt.Printf("User credentials successfully saved to: %s\n", userMSPDir)
}

// Helper function to copy a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Ensure the destination directory exists
	destDir := filepath.Dir(dst)
	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
