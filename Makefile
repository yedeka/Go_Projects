build-ccwc:
	go build -o bin/ccwc cmd/ccwc/main.go

run-ccwc-l: 
	go run cmd/ccwc/main.go -l cmd/ccwc/input/test.txt

run-ccwc-w: 
	go run cmd/ccwc/main.go -w cmd/ccwc/input/test.txt

run-ccwc-c: 
	go run cmd/ccwc/main.go -c cmd/ccwc/input/test.txt

run-ccwc-m: 
	go run cmd/ccwc/main.go -m cmd/ccwc/input/test.txt
