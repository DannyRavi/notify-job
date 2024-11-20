

# Define the applications
APP1 = mario
APP2 = luigi

APP1ADDR = ./cmd/mario/
APP2ADDR = ./cmd/luigi/
init:
	go mod tidy
	mkdir /tmp/lab
	sudo chmod -R 777 /tmp/lab
# Build both applications
build:
	go build -C ./cmd/mario/ -o $(APP1) 
	go build -C ./cmd/luigi/ -o $(APP2)
	chmod +x $(APP1ADDR)/mario
	chmod +x $(APP2ADDR)/luigi




run:
	cd $(APP1ADDR) && ./$(APP1)  cli -d /tmp/lab/out.csv -r 999 &
	cd $(APP2ADDR) && sudo ./$(APP2) cli -d /tmp/lab
mario: 
	cd $(APP1ADDR) && ./$(APP1)  cli -d /tmp/lab/outx.csv -r 999 
luigi:
	cd $(APP2ADDR) && sudo ./$(APP2) cli -d /tmp/lab

testapp:
	cd $(APP1ADDR) && go test 
	cd $(APP2ADDR) && go test

# Clean the build artifacts
clean:
	rm -f $(APP1ADDR)/mario $(APP2ADDR)/luigi