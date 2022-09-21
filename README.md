# alfred PDF to image

Convert PDF to image with Alfred.

[![Release](https://github.com/cage1016/alfred-pdf2image/actions/workflows/release.yml/badge.svg)](https://github.com/cage1016/alfred-pdf2image/actions/workflows/release.yml)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
![GitHub all releases](https://img.shields.io/github/downloads/cage1016/alfred-pdf2image/total)
[![codecov](https://codecov.io/gh/cage1016/alfred-pdf2image/branch/master/graph/badge.svg)](https://codecov.io/gh/cage1016/alfred-pdf2image)
![](https://img.shields.io/badge/Alfred-5-blueviolet)

![](screenshots/demo.gif)

## Features

- Convert PDF to image by page range

## Download
Make sure to download the latest released directly from the releases page. [Download here](https://github.com/cage1016/alfred-pdf2image/releases).

## Requires
- Preferably Alfred 5

## Usage

![](screenshots/usage.jpg)

- File Filter keyword: `pdf2img` or File action to pick up PDF file
- Enter page numbers and/or page ranges.
  - `2` means page 2
  - `5-8` means page 5 to 8
  - `-20` means page 1 to 20
  - `20-` means page 20 to last page
  - `-` means all pages

## Third Party Library

- [gen2brain/go-fitz: Golang wrapper for the MuPDF Fitz library](https://github.com/gen2brain/go-fitz)

## Change Log

### 0.1.0
- Initial release

## License
This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.