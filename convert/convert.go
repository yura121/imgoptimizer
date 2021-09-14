package convert

import (
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image/jpeg"
	"log"
	"os"
)

func ToWebp(src, dst string) {
	file, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
			panic(err)
		}
	}(output)

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Fatalln(err)
	}
}
