
# imgtool

A CLI tool to compress images written in Go.

## Features

- Supports input files or folders
- Supports JPEG, PNG, and WebP formats
- Customizable output format and quality
- Adds suffix to compressed images
- CLI usage with simple commands

## Installation

```bash
go install github.com/dayiamin/imgtool@latest
```

Make sure your `$GOPATH/bin` is in your system's PATH.

## Usage

### Compress Images

```bash
imgtool compress [flags]

```

#### Flags:

- `--input`: Path to image file or folder (required)
- `--output`: Output directory (optional, default: `./compressed`)
- `--format`: Output format: `jpeg`, `png`, or `webp` (optional, default: same as input)
- `--quality`: Compression quality (1-100, optional, default: 80)
- `--suffix`: Suffix for compressed file names (optional, default: `_compressed`)

### Example

```bash
imgtool compress --input ./images --output ./out --format webp --quality 70 --suffix _min
imgtool compress -i ./images -o ./images -f original  -s "_min" -q 80
```

## License

MIT License
