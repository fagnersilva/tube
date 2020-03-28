package importers

import (
	"fmt"
	"strings"

	"github.com/prologic/vimeodl"
)

type VimeoImporter struct{}

func (i *VimeoImporter) GetVideoInfo(url string) (videoInfo VideoInfo, err error) {
	if strings.HasPrefix(url, "vimeo:") {
		url = strings.TrimPrefix(url, "vimeo:")
	}

	if !strings.HasPrefix(url, "http") {
		url = "https://player.vimeo.com/video/" + url
	}

	if !strings.HasPrefix(url, "https://player.vimeo.com/video/") {
		playerURL, err := vimeodl.GetPlayerURL(url)
		if err != nil {
			err := fmt.Errorf("error finding player url: %w", err)
			return VideoInfo{}, err
		}
		url = playerURL
	}

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	url += "config"

	config, err := vimeodl.GetVideoConfig(url)
	if err != nil {
		err := fmt.Errorf("error retrieving video config: %w", err)
		return VideoInfo{}, err
	}

	videoInfo.VideoURL = vimeodl.PickBestVideo(config)

	videoInfo.ThumbnailURL = vimeodl.PickBestThumbnail(config)

	videoInfo.ID = string(config.Video.Id)
	videoInfo.Title = config.Video.Title

	return
}
