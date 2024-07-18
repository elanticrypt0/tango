BINARY_NAME=tango_cli
BINARY_NAME_WIN=tango_cli.exe
BUILD_DIR=./build
BUILD_DIR_WIN=./build/windows

dev:
	go run .

test:
	cd tests
	go test 

build:
	# create directories
	mkdir -p ${BUILD_DIR}
	# compile into binary file
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -gccgoflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR}

build-win:
# 	# create directories
	mkdir -p ${BUILD_DIR_WIN}
# 	# compile into binary file
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME_WIN} -gccgoflags "-w -s"
