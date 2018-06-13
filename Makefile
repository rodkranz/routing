cover: cover-prof cover-func

cover-prof:
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

cover-func:
	go test -covermode=count -coverprofile=count.out fmt
	go tool cover -html=count.out