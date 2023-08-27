package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func InArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func CopyFile(src, dest string, wg *sync.WaitGroup) {
	defer wg.Done()

	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dest)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	fmt.Println("Copied:", src, "to", dest)
}

func CopyDir(srcDir, destDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", err)
			return err
		}

		relPath, _ := filepath.Rel(srcDir, path)
		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			err := os.MkdirAll(destPath, info.Mode())
			fmt.Println("Created directory:", destPath)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return err
			}
		} else {
			wg.Add(1)
			go CopyFile(path, destPath, wg)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking directory:", err)
	}
}

//	func ResourcesFolder(folder ...string) {
//		_, filename, _, _ := runtime.Caller(0)
//		currentDir := filepath.Dir(filename)
//
//		// Calculate the path to the 'laravel' folder
//		dir := filepath.Join(currentDir, "..", "resources", ...folder)
//	}
func ResourcesFolder(folders ...string) string {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	// Prepend ".." to navigate up from the current directory
	pathComponents := append([]string{currentDir, "..", "..", "resources"}, folders...)

	// Calculate the final path
	dir := filepath.Join(pathComponents...)

	return dir
}
