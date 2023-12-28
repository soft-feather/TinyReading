.PHONY: release
release:
	@echo "start release server"
	@mkdir -p build/TinyReading/bin/doc/md
	@go build -o build/TinyReading/bin/server ./cmd/server
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
	@go build -o build/TinyReading/bin/server ./cmd/server
	@cp util/log/log.json ./config.ini build/TinyReading/bin/
	cd build/TinyReading/bin > ./log && bash -c ./server