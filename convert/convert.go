package convert

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
)

func ToWebp(src, dst string) {
	buffer, err := bimg.Read(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if bimg.NewImage(newImage).Type() == "webp" {
		fmt.Fprintln(os.Stderr, "The image was converted into webp")
	}

	bimg.Write(dst, newImage)
}
