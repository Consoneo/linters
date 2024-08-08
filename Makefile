.PHONY: build

build: clean
	@goreleaser build --snapshot
clean:
	@rm -Rf dist||true
