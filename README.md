# GoImageCompressor

📦 A fast and easy-to-use CLI image compression tool written in Go.  
Supports JPEG, PNG, and WebP formats — including batch processing and rotation fix.

---

## 🚀 Features

- ✅ Supports input files and entire folders
- 🖼️ Converts images to WebP or keeps original format
- 📉 Shows original vs compressed size with compression percentage
- ⏱️ Displays processing time per image
- 🔄 Automatically corrects image rotation using EXIF
- ⚙️ Adjustable quality options for WebP

---

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

### Output Sample
🔄 Compressing: ./images/pic1.jpg
📉 Reduced: (77.74% smaller)
⏱️ Time: 58ms


## License

MIT License
