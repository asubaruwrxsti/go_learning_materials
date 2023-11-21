package video

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"time"

	"golang_colly_reddit/video/videoUtilities"

	"github.com/golang/freetype"
)

type RedditVideo struct {
	VideoMeta map[string]interface{} `json:"video_meta"`
	Source    string                 `json:"source"`
}

// VideoMeta holds
// length: int
// height: int
// width: int
// dpi: int
// size: int
// path: string

func (rv RedditVideo) ToString() string {
	return fmt.Sprintf("Source: %s\nVideoMeta: %s", rv.Source, rv.VideoMeta)
}

func generateBlankFrame(width int, height int) (*image.RGBA, error) {
	if width < 1 || height < 1 {
		return nil, errors.New("invalid parameters")
	}
	return image.NewRGBA(image.Rect(0, 0, width, height)), nil
}

func addTextToImage(img *image.RGBA, x_offset int, y_offset int, text string) error {
	if img == nil || x_offset < 0 || y_offset < 0 || text == "" {
		return errors.New("invalid parameters")
	}

	col := color.RGBA{200, 100, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{col}, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	font, err := videoUtilities.LoadFont("/home/arlind/Desktop/go_learning_materials/golang_colly_reddit/video/fonts/Maleha.ttf")
	if err != nil {
		return err
	}
	c.SetFont(font)
	c.SetFontSize(12)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)
	pt := freetype.Pt(x_offset, y_offset+int(c.PointToFixed(12)>>6))
	_, err = c.DrawString(text, pt)
	if err != nil {
		return err
	}
	return nil
}

func saveImage(img *image.RGBA, filename string, path string) error {
	if img == nil || filename == "" {
		return errors.New("invalid parameters")
	}

	// check if path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Path does not exist, creating path...")
		os.Mkdir(path, 0755)
	}

	// check if the path is a directory
	if !videoUtilities.IsDir(path) {
		return errors.New("path is not a directory")
	} else {
		// empty the directory
		if err := videoUtilities.EmptyDir(path); err != nil {
			return err
		}
	}

	f, err := os.Create(path + filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return nil
}

/* CreateRedditVideo creates a video from a reddit post
 * videoMeta: map[string]int
 * storyComments: []string
 */
func CreateRedditVideo(videoMeta map[string]interface{}, storyComments []string, path string) error {
	defer fmt.Println("Image created successfully !")

	if videoMeta == nil || storyComments == nil || path == "" {
		return errors.New("invalid parameters")
	}
	var videoHeight int = videoMeta["VideoMeta"].(map[string]interface{})["height"].(int)
	var videoWidth int = videoMeta["VideoMeta"].(map[string]interface{})["width"].(int)

	// Create a blank frame
	img, err := generateBlankFrame(videoWidth, videoHeight)
	if err != nil {
		return err
	}

	// Add text to the frame
	for i, cmt := range storyComments {
		if err := addTextToImage(img, 10, i*500, cmt); err != nil {
			return err
		}
	}

	// Save the image
	if err := saveImage(img, "test"+time.Now().String()+".jpg", path); err != nil {
		return err
	}
	return nil
}
