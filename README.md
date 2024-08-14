# FabricHyperledgerApplication

# FabricHyperledgerApplication
How to run this project:

## 1. Start network and create channel from test-network

```
cd test-network
./network.sh up createChannel -c mychannel -ca
```

## 2. Deploy chaincode to peer
```
./network.sh deployCC -ccn abac -ccp ../asset-transfer-abac/chaincode-go/ -ccl go
```

## 3. Register identities with attributes
#### Set environment variables
```
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
```
#### Create the identities using the Org1 CA
```
export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.example.com/
```

#### Register an identity named creator1 with the attribute of abac.creator=true.
```
fabric-ca-client register --id.name creator1 --id.secret creator1pw --id.type client --id.affiliation org1 --id.attrs 'abac.creator=true:ecert' --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
```

#### Enrol the identity that registered
```
fabric-ca-client enroll -u https://creator1:creator1pw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/creator1@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
```

#### Copy the Node OU configuration file into the creator1 MSP folder.
```
cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/creator1@org1.example.com/msp/config.yaml"
```

## 4. Run application
```
cd ../rest-api-go
go mod download
go run main.go
```
## 5. Send request
#### Create asset123
```
curl --request POST \
  --url http://localhost:3000/invoke \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data channelid=mychannel \
  --data chaincodeid=abac \
  --data function=CreateAsset \
  --data args=Asset123 \
  --data args=yellow \
  --data args=54 \
  --data args=Tom
```

#### Get query123 from ledger 
```
curl --request GET \
  --url 'http://localhost:3000/query?channelid=mychannel&chaincodeid=abac&function=ReadAsset&args=Asset123' 
```
#### Update asset123 in ledger  
```
curl --request PUT \
  --url http://localhost:3000/update \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data channelid=mychannel \
  --data chaincodeid=abac \
  --data function=UpdateAsset \
  --data args=Asset123 \
  --data args=green \
  --data args=20 \
  --data args=Quang
```

#### Delete asset123 in ledger
```
curl --request DELETE \
  --url 'http://localhost:3000/delete?channelid=mychannel&chaincodeid=abac&function=DeleteAsset&args=Asset123'
```

#### Transfer asset123 to user1
```
curl --request PUT \
  --url http://localhost:3000/transfer \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data channelid=mychannel \
  --data chaincodeid=abac \
  --data assetid=Asset123 \
  --data recipientCN=user1

```

## 6.Clean up

```
cd ../test-network
./test-network down
```
