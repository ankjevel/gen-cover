.DEFAULT_GOAL := bin
INSTALL=install -p -m 644

.PHONY: build clean bin release

build: bin

clean:
	@rm -r bin

bin:
	@mkdir -p bin
	@go build -o bin/gen_cover
	@echo '"bin" successful'

clean_release:
	@rm -rf release/

release: clean_release
	@mkdir -p release/gen_cover/fonts/src release/gen_cover/static
	@go build -ldflags="-s -w" -o release/gen_cover/bin
	@$(INSTALL) static/* release/gen_cover/static/
	@$(INSTALL) fonts/src/*ttf release/gen_cover/fonts/src/
	@$(INSTALL) .template.env release/gen_cover/.template.env
	@echo '"release" successful'
