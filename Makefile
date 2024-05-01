OUTPUT_DIR?=./out

.PHONY: build-local
build-local: build-local-forward-proxy

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

.PHONY: build-local-forward-proxy
build-local-forward-proxy: download-core-packages download-forward-proxy-packages create-output-dir
	go build -o ${OUTPUT_DIR}/forward-proxy ./forward-proxy	

.PHONY: clean
clean:
	rm -rf ${OUTPUT_DIR}