# Morse translator

This is a very simple CLI for translating to and from morse code, written in Go

## Installation

run this command. If you don't have Go [download and install it](https://go.dev/dl) first.

`go install github.com/klrfl/morse-translator`

## Usage

Currently there is only one command in this program: `translate`. It consists of
two flags, `target` (`t` for short and also optional, defaults to `plain`) and `input`
(`i` for short).

To translate to morse code, you only need to specify the input:

`morse-translator translate -i 'string you want to translate'`

To translate from morse code, you need to specify the target like so:

`morse-translator translate --target plain -i '-- --- .-. ... . / -.-- --- ..- / .-- .- -. - / - --- / - .-. .- -. ... .-.. .- - .'`

## License

MIT
