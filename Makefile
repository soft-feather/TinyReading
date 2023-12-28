GO = go
PROJECT_NAME = TinyReading
PN = $(PROJECT_NAME)
BUILD_ROOT_DIR = build
BUILD_DST_DIR = $(BUILD_ROOT_DIR)/$(PN)

INPUT_SERVER_DIR = ./cmd/server
INPUT_FILE_LS = $(wildcard util/log/log.json ./config.ini)


OUTPUT_BIN_DIR = $(BUILD_DST_DIR)/bin


.PHONY: release
release:
	@echo "start release server"
	@mkdir -p build/TinyReading/bin/doc/md
	@$(GO) build -o build/TinyReading/bin/server ./cmd/server
	@cp util/log/log.json ./config.ini build/TinyReading/bin/
	@tar zcvf TinyReading.tar.gz build/

.PHONY: clean
clean:
	@echo "start clean"
	@rm build/* -r

.PHONY: test-run
test-run:
	@echo "test run"
	@mkdir -p build/TinyReading/bin/doc/md
	@$(GO) build -o build/TinyReading/bin/server ./cmd/server
	@cp util/log/log.json ./config.ini build/TinyReading/bin/
	cd build/TinyReading/bin > ./log && bash -c ./server

# build_bin $input $output
define build_bin
	$(GO) build -o $(1)
endef

.PHONY: test
test:
	@echo "test"
	@mkdir -p build/TinyReading/bin/doc/md
	$(call build_bin, $(OUTPUT_BIN_DIR)/server $(INPUT_SERVER_DIR))
	cp $(INPUT_FILE_LS) $(OUTPUT_BIN_DIR)

	cd $(OUTPUT_BIN_DIR) > ./log && bash -c ./server