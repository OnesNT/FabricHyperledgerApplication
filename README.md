# FabricHyperledgerApplication
How to run this project:

## 1. Start network and create channel from test-network

```
cd test-network
./network.sh up createChannel -c mychannel -ca
```

## 2. Deploy chaincode to peer
```
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go
```
## 3. Run application
```
cd ../rest-api-go
go mod download
go run main.go
```
## 4. Send request
#### Create asset123
```
curl --request POST \
  --url http://localhost:3000/invoke \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data = \
  --data channelid=mychannel \
  --data chaincodeid=basic \
  --data function=createAsset \
  --data args=Asset123 \
  --data args=yellow \
  --data args=54 \
  --data args=Tom \
  --data args=13005
```

#### Get query123 from ledger 
```
curl --request GET \
  --url 'http://localhost:3000/query?channelid=mychannel&chaincodeid=basic&function=ReadAsset&args=Asset123' 
```
#### Update asset123 in ledger  
```
curl --request PUT \
  --url http://localhost:3000/update \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data = \
  --data channelid=mychannel \
  --data chaincodeid=basic \
  --data function=createAsset \
  --data args=Asset123 \
  --data args=red \
  --data args=543 \
  --data args=Quang \
  --data args=666
```

#### Delete asset123 in ledger
```
curl --request DELETE \
  --url 'http://localhost:3000/delete?channelid=mychannel&chaincodeid=basic&function=DeleteAsset&args=Asset123'
```

#### Transfer asset123 to newOwner
```
curl --request POST \
  --url 'http://localhost:3000/transfer?channelid=mychannel&chaincodeid=basic&function=TransferAsset&args=Asset123&args=NewOwner'
```

## 5.Clean up

```
cd ../test-network
./test-network down
```
