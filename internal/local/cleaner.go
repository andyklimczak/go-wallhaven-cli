package local

import (
	"fmt"
	"github.com/andyklimczak/go-wallhaven/internal/downloader"
	"github.com/andyklimczak/go-wallhaven/internal/logger"
	"github.com/andyklimczak/go-wallhaven/internal/wallhaven"
	"os"
)

type Cleaner struct {
	log *logger.GoHavenLogger
}

func NewDefaultCleaner(log *logger.GoHavenLogger) Cleaner {
	return Cleaner{
		log: log,
	}
}

func (c *Cleaner) RemoveExtraWallpapers(wallpapers wallhaven.SearchData, folder string) error {
	c.log.Debug("Cleaning up old wallpapers in path: %s", folder)
	wallpapersSet := map[string]bool{}
	for _, wallpaper := range wallpapers.Data {
		wallpapersSet[downloader.GetFileNameFromPath(wallpaper.Path)] = true
	}

	files, err := os.ReadDir(folder)
	if err != nil {
		return fmt.Errorf("Unable to read dir at path %s: %w", folder, err)
	}

	for _, file := range files {
		if !file.IsDir() {
			if _, ok := wallpapersSet[file.Name()]; !ok {
				c.log.Debug("Removing wallpaper that is no longer in wallhaven collection: %s", file.Name())
				path := fmt.Sprintf("%s/%s", folder, file.Name())
				err = os.Remove(path)
				if err != nil {
					c.log.Error("Unable to remove wallpaper at path %s: %w", path, err)
				}
			}
		}
	}

	return nil
}
