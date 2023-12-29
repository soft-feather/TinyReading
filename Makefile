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
	$(call pack)
	@tar zcvf TinyReading.tar.gz build/

.PHONY: clean
clean:
	@echo "start clean"
	@rm build/* -r

.PHONY: test
test:
	@echo "test run"
	$(call pack)
	cd build/TinyReading/bin > ./log && ./server

# build_bin $input $output
define build_bin
	$(GO) build -o $(1)
endef


# pack
define pack
	@mkdir -p $(OUTPUT_BIN_DIR)/doc/md
	$(call build_bin, $(OUTPUT_BIN_DIR)/server $(INPUT_SERVER_DIR))
	cp $(INPUT_FILE_LS) $(OUTPUT_BIN_DIR)
endef

.PHONY: dry-run
dry-run:
	@echo "test dry run"
	$(call pack)
	cd $(OUTPUT_BIN_DIR) > ./log && ./server --dry-run