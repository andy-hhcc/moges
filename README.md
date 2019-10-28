## Development instruction
1. Create database `moges` in mysql
2. Run `make install` to install dependencies
3. Run `go run main.go` to start running the local server
4. App will run on port `8080`

## Build instruction and run
1. Run `make build`

## Call API using curl
```
curl -X POST \
  http://localhost:8080/upload \
  -H 'Accept: */*' \
  -H 'Authorization: VALID_TOKEN' \
  -H 'Host: localhost:8080' \
  -F file=@/Users/duypham/Downloads/Banner_Multibank4.gif
```
