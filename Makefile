tidy:
	go mod tidy
	npm i

run:
	templ generate
	esbuild --bundle ./react/index.ts --outdir=./static --minify
	go run .
