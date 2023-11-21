package video

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type redditVideo struct {
	VideoMeta map[string]int
	Source    string
}

// VideoMeta holds
// length: int
// height: int
// width: int
// dpi: int
// size: int

func (rv redditVideo) ToString() string {
	return fmt.Sprintf("Source: %s\nVideoMeta: %s", rv.Source, rv.VideoMeta)
}

func generateBlankFrame(width int, height int) (*image.RGBA, error) {
	if width < 1 || height < 1 {
		return nil, errors.New("Invalid parameters")
	}
	return image.NewRGBA(image.Rect(0, 0, width, height)), nil
}

func addTextToImage(img *image.RGBA, x_offset int, y_offset int, text string) error {
	if img == nil || x_offset < 0 || y_offset < 0 || text == "" {
		return errors.New("Invalid parameters")
	}

	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x_offset), fixed.Int26_6(y_offset)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
	return nil
}

func saveImage(img *image.RGBA, filename string, path string) error {
	if img == nil || filename == "" {
		return errors.New("Invalid parameters")
	}
	f, err := os.Create(path + filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return err
	}
}

/* CreateRedditVideo creates a video from a reddit post
 * videoMeta: map[string]int
 * storyComments: []string
 */
func CreateRedditVideo(videoMeta map[string]int, storyComments []string, path string) error {

}
