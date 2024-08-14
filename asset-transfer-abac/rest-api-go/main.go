package main

import (
	"fabric-samples/asset-transfer-abac/rest-api-go/web"
	"fmt"
	"os"
)

// setEnvVariables sets the required environment variables programmatically.
func setEnvVariables() {
	// Set CORE_PEER_TLS_ENABLED
	os.Setenv("CORE_PEER_TLS_ENABLED", "true")

	// Set CORE_PEER_LOCALMSPID
	os.Setenv("CORE_PEER_LOCALMSPID", "Org1MSP")

	// Set CORE_PEER_MSPCONFIGPATH
	os.Setenv("CORE_PEER_MSPCONFIGPATH", "../../test-network/organizations/peerOrganizations/org1.example.com/users/creator1@org1.example.com/msp")

	// Set CORE_PEER_TLS_ROOTCERT_FILE
	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE", "../../test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt")

	// Set CORE_PEER_ADDRESS
	os.Setenv("CORE_PEER_ADDRESS", "localhost:7051")

	// Set TARGET_TLS_OPTIONS
	tlsOptions := "-o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile " +
		"\"../../test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem\" " +
		"--peerAddresses localhost:7051 --tlsRootCertFiles " +
		"\"../../test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt\" " +
		"--peerAddresses localhost:9051 --tlsRootCertFiles " +
		"\"../../test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt\""
	os.Setenv("TARGET_TLS_OPTIONS", tlsOptions)
}
func main() {
	//Initialize setup for Org1
	setEnvVariables()

	cryptoPath := "../../test-network/organizations/peerOrganizations/org1.example.com"
	orgConfig := web.OrgSetup{
		OrgName:      "Org1",
		MSPID:        "Org1MSP",
		CertPath:     cryptoPath + "/users/creator1@org1.example.com/msp/signcerts/cert.pem",
		KeyPath:      cryptoPath + "/users/creator1@org1.example.com/msp/keystore/",
		TLSCertPath:  cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt",
		PeerEndpoint: "dns:///localhost:7051",
		GatewayPeer:  "peer0.org1.example.com",
	}

	orgSetup, err := web.Initialize(orgConfig)
	if err != nil {
		fmt.Println("Error initializing setup for Org1: ", err)
	}
	web.Serve(web.OrgSetup(*orgSetup))
}
