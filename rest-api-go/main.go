package main

import (
	"fmt"
	"rest-api-go/web"
)

func main() {

	cryptoPath := "../../test-network/organizations/peerOrganizations/org1.example.com"
	orgConfig := web.OrgSetup{
		OrgName:      "Org1",
		MSPID:        "Org1MSP",
		CertPath:     cryptoPath + "/users/user7@org1.example.com/msp/signcerts/cert.pem",
		KeyPath:      cryptoPath + "/users/user7@org1.example.com/msp/keystore/",
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

// package main

// import (
// 	"rest-api-go/user"
// )

// func main() {
// 	// Start the HTTP server from the user package
// 	user.Serve()
// }