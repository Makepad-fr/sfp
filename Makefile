OUTPUT_DIR?=./out

.PHONY: build
build: build-forward-proxy

.PHONY: download-forward-proxy-packages
download-forward-proxy-packages:
	cd ./forward-proxy && go mod download

.PHONY: download-core-packages
download-core-packages:
	cd ./core && go mod download

.PHONY: download
download: download-core-packages download-forward-proxy-packages

.PHONY: create-output-dir
create-output-dir:
	mkdir -p ${OUTPUT_DIR}

.PHONY: build-forward-proxy
build-forward-proxy: download-core-packages download-forward-proxy-packages create-output-dir
	go build -o ${OUTPUT_DIR}/forward-proxy ./forward-proxy	

.PHONY: clean
clean:
	rm -rf ${OUTPUT_DIR}