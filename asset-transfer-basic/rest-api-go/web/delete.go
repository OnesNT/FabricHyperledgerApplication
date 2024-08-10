package web

import (
	"fmt"
	"net/http"
)

// Delete handles chaincode delete requests.
func (setup OrgSetup) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Delete Asset request")

	queryParams := r.URL.Query()
	chainCodeName := queryParams.Get("chaincodeid")
	channelID := queryParams.Get("channelid")
	function := queryParams.Get("function")
	args := queryParams["args"]

	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)

	// Ensure required parameters are present
	if chainCodeName == "" || channelID == "" || function == "" || len(args) == 0 {
		http.Error(w, "Missing chaincodeid, channelid, function, or args", http.StatusBadRequest)
		return
	}

	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)

	// Submit the transaction
	txn_committed, err := contract.SubmitTransaction(function, args...)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	// Respond with the transaction ID
	fmt.Fprintf(w, "Transaction ID: %s", txn_committed)
}
