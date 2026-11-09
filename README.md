# imgdim

A CLI that shows the width, height and aspect ratio of an image.

Formats: PNG, JPEG, GIF and WEBP.

## Installation

```sh
go install github.com/JacobsM14/imgdim@latest
```

## Usage

```sh
imgdim photo.png
```

```
Width: 1920
Height: 1080
Aspect Ratio: 1.78
Orientation: Horizontal
Simplified: 16:9
Common Ratio: 16:9
```

## Shell completion

```sh
source <(imgdim completion zsh)
```
