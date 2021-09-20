package convert

import (
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image/jpeg"
	"imgoptimizer/config"
	"os"
)

type Conv struct {
	quality float32
}

func (c *Conv) ToWebp(src, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	output, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, c.quality)
	if err != nil {
		return err
	}

	err = webp.Encode(output, img, options)
	if err != nil {
		return err
	}

	return nil
}

func New(conf *config.Conf) *Conv {
	var c Conv
	c.quality = conf.Converter.Quality
	return &c
}
