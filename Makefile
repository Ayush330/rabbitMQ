EXECUTABLE_NAME = RabbitApplication

.PHONY: run

run: build
	./$(EXECUTABLE_NAME)
    
build:
	go build -o $(EXECUTABLE_NAME) *.go
