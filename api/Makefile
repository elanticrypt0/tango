BINARY_NAME=app_api
BINARY_NAME_WIN=app_api.exe

BUILD_DIR=../build

BUILD_DIR_LINUX=../build/linux
BUILD_DIR_LINUXARM64=../build/arm64
BUILD_DIR_WIN=../build/windows

tango-install:
	mkdir _db
	make go-install-deps
	make tailwind-install
	make htmlx-install
	make templates
	go mod tidy

go-install-deps:
	go install github.com/cosmtrek/air@latest
	go mod tidy

build-linux:
	# create directories
	mkdir -p ${BUILD_DIR}
	mkdir -p ${BUILD_DIR_LINUX}
	rm -rf ${BUILD_DIR_LINUX}/*
	mkdir -p ${BUILD_DIR_LINUX}/_db
	mkdir -p ${BUILD_DIR_LINUX}/logs
	mkdir -p ${BUILD_DIR_LINUX}/public

	# copy the config files and seeds
	cp -r ./config ${BUILD_DIR_LINUX}/config
	cp -r ./seeds ${BUILD_DIR_LINUX}/seeds

	# compile into binary file
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -gccgoflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR_LINUX}

build-linuxarm64:
	# create directories
	mkdir -p ${BUILD_DIR}
	mkdir -p ${BUILD_DIR_LINUXARM64}
	rm -rf ${BUILD_DIR_LINUXARM64}/*
	mkdir -p ${BUILD_DIR_LINUXARM64}/_db
	mkdir -p ${BUILD_DIR_LINUXARM64}/logs
	mkdir -p ${BUILD_DIR_LINUXARM64}/public

	# copy the config files and seeds
	cp -r ./config ${BUILD_DIR_LINUXARM64}/config
	cp -r ./seeds ${BUILD_DIR_LINUXARM64}/seeds

	# compile into binary file
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -gccgoflags "-w -s"
	chmod +x ${BINARY_NAME}
	mv ${BINARY_NAME} ${BUILD_DIR_LINUXARM64}

build-win:
	# create directories
	mkdir -p ${BUILD_DIR}
	mkdir -p ${BUILD_DIR_WIN}
	rm -rf ${BUILD_DIR_WIN/*
	mkdir -p ${BUILD_DIR_WIN}/_db
	mkdir -p ${BUILD_DIR_WIN}/logs
	mkdir -p ${BUILD_DIR_WIN}/public

	# copy the config files and seeds
	cp -r ./config ${BUILD_DIR_WIN}/config
	cp -r ./seeds ${BUILD_DIR_WIN}/seeds

	# compile into binary file
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME_WIN} -gccgoflags "-w -s"
	chmod +x ${BINARY_NAME_WIN}
	mv ${BINARY_NAME_WIN} ${BUILD_DIR_WIN}

dev:
	go run .

dev-air:
	make templates-clean
	make templates
	air

run:
	go run .

test:
	go test ./tests

clean:
	go clean
	rm -rf ${BUILD_DIR}

cli-install:
	git clone https://github.com/k23dev/tango_cli
	cd tango_cli
	GOOS=linux GOARCH=amd64 go build -o tango_cli -gccgoflags "-w -s"
	chmod +x tango_cli
	cp tango_cli ../
