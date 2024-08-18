package web

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// Update handles chaincode invoke requests for updating an asset.
func (setup *OrgSetup) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Received Update request")
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %s", err), http.StatusBadRequest)
		return
	}

	chainCodeName := r.FormValue("chaincodeid")
	channelID := r.FormValue("channelid")
	function := r.FormValue("function")
	args := r.Form["args"]

	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)

	// Access the network and contract
	network := setup.Gateway.GetNetwork(channelID)
	if network == nil {
		http.Error(w, fmt.Sprintf("Channel ID %s does not exist or cannot be accessed", channelID), http.StatusBadRequest)
		return
	}

	contract := network.GetContract(chainCodeName)

	// Create and submit the transaction proposal
	txn_proposal, err := contract.NewProposal(function, client.WithArguments(args...))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating txn proposal: %s", err), http.StatusInternalServerError)
		return
	}

	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error endorsing txn: %s", err), http.StatusInternalServerError)
		return
	}

	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error submitting transaction: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())
}
