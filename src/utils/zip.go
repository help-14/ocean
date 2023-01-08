package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//func ZipFolder(inputPath string, outputPath string) error {
// file, err := os.Create(outputPath)
// if err != nil {
// 	return err
// }
// defer file.Close()

// w := zip.NewWriter(file)
// defer w.Close()

// walker := func(path string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	if info.IsDir() {
// 		return nil
// 	}
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Ensure that `path` is not absolute; it should not start with "/".
// 	// This snippet happens to work because I don't use
// 	// absolute paths, but ensure your real-world code
// 	// transforms path into a zip-root relative path.
// 	f, err := w.Create(path)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = io.Copy(f, file)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// err = filepath.Walk(inputPath, walker)
// if err != nil {
// 	return err
// }
// return nil
//}

func CompressZip(inputPath string, outputPath string) error {
	archive, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	w := zip.NewWriter(archive)
	stat, err := os.Stat(inputPath)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		dir := inputPath
		err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			arPath := strings.ReplaceAll(path, dir, "")
			arPath = strings.TrimPrefix(arPath, "/")
			f, err := w.Create(arPath)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			in, err := os.Open(path)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			_, err = io.Copy(f, in)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		path := inputPath
		dir := filepath.Dir(path)
		arPath := strings.ReplaceAll(path, dir, "")
		arPath = strings.TrimPrefix(arPath, "/")
		f, err := w.Create(arPath)
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, in)
		if err != nil {
			return err
		}
	}

	w.Close()
	archive.Close()
	return nil
}
