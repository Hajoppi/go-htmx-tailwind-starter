# Go HTMX Tailwind Starter
This is a very minimalistic starter with very few dependencies.

This starter uses go's standard library for routing and template rendering.
The only direct dev-dependencies are Tailwindcss for styling, browser-sync for hot-reloading and prettier for go-templates.
For production dependency, there is only HTMX

## Requirements

- Install Go: [https://go.dev/doc/install](https://go.dev/doc/install)
- Install Air: `go install github.com/air-verse/air@latest`
- (Optional: Add go's bin directory to path)
- Install npm [https://nodejs.org/en/download/package-manager](https://nodejs.org/en/download/package-manager)
- Download dependencies `npm install`
- Just to make sure `go mod tidy`

Note: This starter assumes `air` is installed in ´~/go/bin´, so update the location in `scripts/dev.sh` if it's not the same.

## Development

Run
```bash
./scripts/dev.sh
```
