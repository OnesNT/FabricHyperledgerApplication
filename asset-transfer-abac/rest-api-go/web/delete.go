package web

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// Delete handles chaincode delete requests.
func (setup *OrgSetup) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Received Delete Asset request")

	// Parse query parameters
	queryParams := r.URL.Query()
	chainCodeName := queryParams.Get("chaincodeid")
	channelID := queryParams.Get("channelid")
	function := queryParams.Get("function")
	args := queryParams["args"]

	// Validate required parameters
	if chainCodeName == "" || channelID == "" || function == "" || len(args) == 0 {
		http.Error(w, "Missing chaincodeid, channelid, function, or args", http.StatusBadRequest)
		return
	}

	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %v\n", channelID, chainCodeName, function, args)

	// Retrieve the network and contract from the gateway
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)

	// Create a transaction proposal
	txn_proposal, err := contract.NewProposal(function, client.WithArguments(args...))
	if err != nil {
		fmt.Fprintf(w, "Error creating transaction proposal: %s", err)
		return
	}

	// Endorse the transaction proposal
	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		fmt.Fprintf(w, "Error endorsing transaction: %s", err)
		return
	}

	// Submit the endorsed transaction
	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		fmt.Fprintf(w, "Error submitting transaction: %s", err)
		return
	}

	// Respond with the transaction ID and result
	fmt.Fprintf(w, "Transaction ID: %s, Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())
}
