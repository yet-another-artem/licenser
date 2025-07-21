# Licenser

Licenser is a lightweight tool written in Go that helps you automate copyright creditionals for your projects. It’s designed to be simple to use, fast, and easily integrated into CI/CD pipelines or local development workflows.

## Features

- **Automatic copyright creditionals** for files and dependencies
- **Fast and efficient** operation with minimal dependencies
- **Easy integration** into build systems and CI/CD

## Installation

You can install Licenser in several ways:

### 1. Prebuilt Binaries

Release binaries are available on the [GitHub Releases page](https://github.com/yet-another-artem/licenser/releases).

Download the binary for your platform, extract it, and move it to your `$PATH`.

```bash
wget https://github.com/yet-another-artem/licenser/releases/download/vX.Y.Z/licenser-linux-amd64.tar.gz
tar -xvf licenser-linux-amd64.tar.gz
sudo mv licenser /usr/local/bin/
```

### 2. Build from Source

Make sure you have [Go](https://golang.org/dl/) installed (Go 1.18+ recommended):

```bash
git clone https://github.com/yet-another-artem/licenser.git
cd licenser
go build -o licenser
```

Move the resulting binary to your path:

```bash
sudo mv licenser /usr/local/bin/
```

### 3. Docker

You can run Licenser using Docker:

```bash
docker build -t licenser .
docker run --rm -v ./:./data licenser [options]
```

## Usage

Licenser can be run from the command line:

```bash
licenser <path> <licenser_name>
```

For example, to check licenses in your current directory:

```bash
licenser . yet.another.artem
```

## Example

```bash
licenser ./src/ yet.another.artem
```

## Contributing

Contributions are welcome! Please open issues or pull requests on the [GitHub repository](https://github.com/yet-another-artem/licenser).

## License

This project is licensed under the terms specified in the [LICENSE](./LICENSE) file.

---

The original idea of the program was given to me by my gal pal, named Kira. Thanks to her! ❤️

Made by "yet.another.artem".

Copyright (c) 2025 yet.another.artem
