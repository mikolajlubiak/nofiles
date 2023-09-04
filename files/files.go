package files

import (
	"path/filepath"
	"log"
	"os"
)

func Files(dir string) []string {
	files := make([]string, 100)
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files = append(files, path)
			return nil
	})
	if err != nil {
		log.Println(err)
	}
	return files
}
