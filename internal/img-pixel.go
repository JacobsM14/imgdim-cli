package internal

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type AspectInfo struct {
	Width       int
	Height      int
	Ratio       float64
	Orientation string
	Simplified  string
	Common      string
}

var aspectRatios = map[string]string{
	// Square
	"1:1": "1:1",

	// Horizontal
	"5:4":   "5:4",
	"4:3":   "4:3",
	"3:2":   "3:2",
	"8:5":   "16:10",
	"5:3":   "5:3",
	"16:9":  "16:9",
	"2:1":   "2:1",
	"20:9":  "20:9",
	"21:9":  "21:9",
	"64:27": "64:27",
	"32:9":  "32:9",

	// Vertical
	"4:5":   "4:5",
	"3:4":   "3:4",
	"2:3":   "2:3",
	"5:8":   "10:16",
	"3:5":   "3:5",
	"9:16":  "9:16",
	"1:2":   "1:2",
	"9:20":  "9:20",
	"9:21":  "9:21",
	"27:64": "27:64",
	"9:32":  "9:32",
}

func GetImageDimensions(filePath string) (width, height int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}

func getCommonRatio(simplified string) string {
	if ratio, exists := aspectRatios[simplified]; exists {
		return ratio
	}

	return "Custom"
}

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func GetImageInfo(filePath string) (info AspectInfo, err error) {
	width, height, err := GetImageDimensions(filePath)
	if err != nil {
		return AspectInfo{}, err
	}

	if width == 0 && height == 0 {
		return AspectInfo{}, fmt.Errorf("invalid img: width and height are 0")
	} else if width == 0 {
		return AspectInfo{}, fmt.Errorf("invalid img: width is 0")
	} else if height == 0 {
		return AspectInfo{}, fmt.Errorf("invalid img: height is 0")
	}

	divisor := GreatestCommonDivisor(width, height)
	ratio := float64(width) / float64(height)
	orientation := "Square"

	if width > height {
		orientation = "Horizontal"
	} else if height > width {
		orientation = "Vertical"
	}

	// use divisor to simplify the aspect ratio
	simplified := fmt.Sprintf("%d:%d",
		width/divisor,
		height/divisor,
	)

	commonRatio := getCommonRatio(simplified)

	return AspectInfo{
		Width:       width,
		Height:      height,
		Ratio:       ratio,
		Orientation: orientation,
		Simplified:  simplified,
		Common:      commonRatio,
	}, nil
}
