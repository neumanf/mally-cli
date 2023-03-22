<h1 align="center">
  <img alt="mally logo" src="https://raw.githubusercontent.com/neumanf/mally/main/frontend/public/logo.svg" width="80px"/><br/>
  Mally CLI
</h1>
<p align="center">Create short URLs and pastes directly from your terminal</p>

## Quick start

First, [download](https://golang.org/dl/) and install **Go**. Version `1.19` or higher is required.

Then, install `mally-cli` by running:

```sh
go install github.com/neumanf/mally-cli@latest
```

## Usage

```sh
Usage:
  mally-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  login       Login into the Mally website
  paste       Creates a pastes from a code or text snippet
  shorten     Shortens a URL

Flags:
  -h, --help   help for mally-cli
```

### Login

To login, first you need an account in [Mally](https://mally.neumanf.com). That said, run:

```sh
mally-cli login
```

You will be asked for your credentials, if they are correct, you will be logged successfully shortly after.

### URL Shortener

To shorten an URL, do:

```sh
mally-cli shorten <url>
```

> Example: `mally-cli shorten https://github.com/neumanf/mally`

### Pastebin

To paste a text or code snippet, do:

```sh
# Using text
mally-cli paste <syntax> "<text>"
# Using a file
mally-cli paste <syntax> -f /path/to/file
```

> Example: `mally-cli paste text "Hello, world!"`

> Example: `mally-cli paste go -f ./main.go`

## License

Mally CLI is licensed under the [MIT License](LICENSE).