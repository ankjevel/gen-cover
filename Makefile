.DEFAULT_GOAL := bin
INSTALL=install -p -m 644

.PHONY: build clean bin release

build: bin

clean:
	@rm -r bin

bin:
	@mkdir -p bin
	@go build -o bin/gen-cover
	@echo '"bin" successful'

clean_release:
	@rm -rf release/

release: clean_release
	@mkdir -p release/gen-cover/fonts/src release/gen-cover/static
	@go build -ldflags="-s -w" -o release/gen-cover/bin
	@$(INSTALL) static/* release/gen-cover/static/
	@$(INSTALL) fonts/src/*ttf release/gen-cover/fonts/src/
	@$(INSTALL) .env-template release/gen-cover/.env-template
	@echo '"release" successful'
