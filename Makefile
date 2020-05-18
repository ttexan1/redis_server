deploy:
	cd server
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build
	scp ./server fenrir:~/redis
	rm ./server
	cd ../tester
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build
	scp ./tester fenrir:~/redis
	rm ./tster
