package video

import "image"

type redditVideo struct {
	VideoMeta []string
	Source    string
}

// VideoMeta holds
// 0: video length
// 1: video height width
// 2: video dpi
// 3: video size in bytes

func generateBlankFrame(videoMeta []string) (image.RGBA, error) {

}

func addTextToImage(img *image.RGBA, x_offset int, y_offset int, text string) (image.RGBA, error) {

}

func CreateRedditVideo(videoMeta []string) (redditVideo, error) {

}
