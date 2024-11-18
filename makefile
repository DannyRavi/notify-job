

# Define the applications
APP1 = mario
APP2 = luigi

APP1ADDR = ./cmd/mario/mario
APP2ADDR = ./cmd/luigi/luigi
init:
	go mod tidy
# Build both applications
build:
	go build -C ./cmd/mario/ -o $(APP1) 
	go build -C ./cmd/luigi/ -o $(APP2) 

# Run both applications
run: build
	./$(APP1ADDR) cli -d ./tmp/out.csv -r 500 & ./$(APP2ADDR) cli -d ./temp &

# Clean the build artifacts
clean:
	rm -f $(APP1) $(APP2)