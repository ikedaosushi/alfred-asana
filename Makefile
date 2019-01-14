.PHONY: build
build:
	go build -o "build/${PROJECT_NAME}"
	cp assets/* build/
	cp build/* "${ALFREDWF_DIR}"