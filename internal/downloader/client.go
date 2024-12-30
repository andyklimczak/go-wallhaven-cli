package downloader

import (
	"fmt"
	"github.com/andyklimczak/go-wallhaven/internal/logger"
	"github.com/andyklimczak/go-wallhaven/internal/wallhaven"
	"github.com/cavaliergopher/grab/v3"
	"strings"
)

type Downloader struct {
	grabber *grab.Client
	host    string
	threads int
	log     *logger.GoHavenLogger
}

func NewDefaultWallhavenDownloader(threads int, log *logger.GoHavenLogger) *Downloader {
	return &Downloader{
		grabber: grab.NewClient(),
		host:    wallhaven.WallhavenApiHost,
		threads: threads,
		log:     log,
	}
}
func (d *Downloader) DownloadFromCollection(collection *wallhaven.Collection, searchData wallhaven.SearchData, destination string) error {
	d.log.Debug("Downloading collection: %s to: %s", collection.Label, destination)
	reqs := make([]*grab.Request, 0, len(searchData.Data))
	for _, listing := range searchData.Data {
		fileName := d.getFileNameFromPath(listing.Path)
		req, err := grab.NewRequest(fmt.Sprintf("%s/%s", destination, fileName), listing.Path)
		if err != nil {
			d.log.Error("Unable to download %s: %w", listing.ID, err)
		}
		reqs = append(reqs, req)
	}
	respch := d.grabber.DoBatch(d.threads, reqs...)
	for resp := range respch {
		switch resp.HTTPResponse.StatusCode {
		case 200:
			d.log.Debug("Downloaded %s into %s", resp.Request.URL(), resp.Filename)
		default:
			d.log.Error("Error downloading file %s with status code: %d", resp.Request.URL(), resp.HTTPResponse.StatusCode)
		}
	}

	return nil
}

func (d *Downloader) getFileNameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}
