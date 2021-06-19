# chain
## It is simple enough for new bird to learn blockchain.
```
go run cmd/main.go
```
### Boot nodes
```
$ go run blockchain.go -p 5001
$ go run blockchain.go -p 5002
$ go run blockchain.go -p 5003
```
### Chain info
```
curl localhost:5000/chain
```
### Add node
```
$ curl -X POST -H "Content-Type: application/json" -d '{
"nodes": ["http://localhost:5001","http://localhost:5002"]
}' localhost:5000/nodes/register
```
### New transaction
```
$ curl -X POST -H "Content-Type: application/json" -d '{
"sender": "96884cb7d98646128c25d03782a9e269",
"recipient": "5e27934f73154922b3ec936ac12e7a21",
"amount": 50
}' localhost:5000/transactions/new
```
### Mine block
```
curl localhost:5001/mine
```
### Consensus
```
curl localhost:5000/nodes/resolve
```

