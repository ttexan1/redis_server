deploy:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ./server ./server/ 
	scp ./server/server fenrir:~/redis
	rm ./server/server
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ./tester ./tester/ 
	scp ./tester fenrir:~/redis
	rm ./tester/tester

test:
	go test ./... -coverprofile=coverage.out