.PHONY: all clean serve deps

all: main.wasm serve

%.wasm: %.go
	GOOS=js GOARCH=wasm go generate
	GOOS=js GOARCH=wasm go build -o "$@" "$<"

serve:
	python3 -m http.server --bind localhost 3000

clean:
	rm -f *.wasm wasm_exec.js

deps:
	go mod tidy
	go mod vendor
