# dockerfile-parser-json

A tiny Go CLI that parses a Dockerfile using BuildKit's Dockerfile parser and prints the AST as compact JSON.

## What it does

- Reads Dockerfile content from stdin.
- Parses instructions into a JSON-friendly node tree.
- Outputs compact one-line JSON.

Each node may include:

- `value`: instruction or token value
- `next`: linked-list style next token
- `children`: nested nodes (for heredoc blocks)
- `flags`: Dockerfile flags (for example `--from=...`, `--chown=...`)

## Build

```bash
go build -o dockerfile-parser-json .
```

## Usage

```bash
dockerfile-parser-json < Dockerfile
```

Or:

```bash
cat Dockerfile | dockerfile-parser-json
```

## Example

Input:

```dockerfile
FROM alpine:3.20
COPY --chown=app:app src/ /app/
```

Output (compact JSON):

```json
[{"value":"from","next":{"value":"alpine:3.20"}},{"value":"copy","next":{"value":"src/","next":{"value":"/app/"}},"flags":["--chown=app:app"]}]
```
