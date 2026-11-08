# img-relation

A CLI that shows the width, height and aspect ratio of an image.

Formats: PNG, JPEG and GIF.

## Installation

```sh
go install github.com/JacobsM14/img-relation-cli@latest
```

`go install` names the binary `img-relation-cli`. To name it `img-relation`, build it from the repository:

```sh
go build -o img-relation .
```

## Usage

```sh
img-relation photo.png
```

```
Width: 1920, Height: 1080, Aspect Ratio: 1.78
```

## Shell completion

```sh
source <(img-relation completion zsh)
```
