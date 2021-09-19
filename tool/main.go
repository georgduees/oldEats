package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const inputDir = "./data"
const ooutputDir = "./out"

type SubChapterItem struct {
	ID      int
	Title   string
	Page    int
	Begin   int
	Image   string
	Content string
	Note    string
}
type ChapterItem struct {
	ID         int
	Title      string
	Note       string
	Page       int
	Begin      int
	Subchapter []SubChapterItem
}
type Book struct {
	ID          string
	Author      string
	Title       string
	PublishDate int64
	Country     string
	City        string
	Publisher   string
	Offset      int
	Chapter     []ChapterItem
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

type PDF struct {
	FullPath  string
	FileName  string
	Name      string
	PageCount int
}

func newPDF(fullpath string) *PDF {
	p := PDF{FullPath: fullpath}
	p.FileName = filepath.Base(fullpath)
	p.Name = strings.Split(p.FileName, ".")[0]
	p.PageCount = 1337
	return &p
}

func main() {

	arg := os.Args[1]
	fmt.Println("hello world")
	pdfList, _ := WalkMatch(arg, "*.pdf")
	var l []PDF

	for _, pdf := range pdfList {
		l = append(l, *newPDF(pdf))
		fmt.Println(pdf)

	}
	for _, pdf := range l {
		fmt.Println(pdf.FileName)
		//	outputName := filepath.Dir(pdf.FullPath) + "/" + pdf.Name + ".jpeg"

		fmt.Println("success")

	}

	// Read in PDF files from Data folder
}

// Create Folder in Output folder with filename

// Split PDF into Single Page PDF

// Read Single Page PDF and split into Images and Text per Page
