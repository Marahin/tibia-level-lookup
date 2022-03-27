install-dependencies:
	go mod download

test:
	ginkgo -p -cover ./... 
	go tool cover -html=coverprofile.out -o coverage.html 
