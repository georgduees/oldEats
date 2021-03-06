package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const inputDir = "../source/data"
const ooutputDir = "./out"

type SubChapterItem struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Page    int    `json:"page"`
	Begin   int    `json:"begin"`
	Image   string `json:"image"`
	Content string `json:"content"`
	Note    string `json:"note"`
}
type ChapterItem struct {
	ID         int              `json:"id"`
	Title      string           `json:"title"`
	Note       string           `json:"note"`
	Page       int              `json:"page"`
	Begin      int              `json:"begin"`
	Offset     int              `json:"offset"`
	Image      string           `json:"image"`
	Subchapter []SubChapterItem `json:"subchapter"`
}
type Book struct {
	ID          string        `json:"id"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	PublishDate int64         `json:"publishdate"`
	Country     string        `json:"country"`
	City        string        `json:"city"`
	Publisher   string        `json:"publisher"`
	Offset      int           `json:"offset"`
	BasePath    string        `json:"basepath"`
	TOCPath     string        `json:"tocpath"`
	CoverImage  string        `json:"coverimage"`
	Chapter     []ChapterItem `json:"chapter"`
}

type PDF struct {
	FullPath  string
	FileName  string
	Name      string
	PageCount int
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

type TOCLine struct {
	CategoryNo int
	Category   string
	Number     int
	Title      string
	Page       int
	Image      string
	Offset     int
}

func stringToInt(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}
func unique(inSlice []TOCLine) []TOCLine {
	keys := make(map[int]bool)
	list := []TOCLine{}
	for _, entry := range inSlice {
		if _, value := keys[entry.CategoryNo]; !value {
			keys[entry.CategoryNo] = true
			list = append(list, entry)
		}
	}
	return list
}
func main() {
	var bookList []Book
	tocList, _ := WalkMatch(inputDir, "toc.csv")
	for _, v := range tocList {
		curBook := Book{Title: v}
		file, err := os.Open(v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("successfully opened file" + v)
		defer file.Close()
		//Get All lines form CSV File
		csvReader := csv.NewReader(file)
		csvReader.Comma = ';'
		csvReader.Comment = '#'

		//Read first line
		if _, err := csvReader.Read(); err != nil {
			panic(err)
		}
		var content []TOCLine
		//CategoryNo;Category;Number;Title;Page;Image;Offset

		records, err := csvReader.ReadAll()
		if err != nil {
			panic(err)
		}
		//fill book with meaningful data
		if strings.Contains(curBook.Title, "sguf") {
			curBook.Title = "Supp Gem??s' und Fleisch"
			curBook.TOCPath = "/sguf/toc.json"
			curBook.CoverImage = "/sguf/cover.jpeg"
		} else {
			if strings.Contains(curBook.Title, "kfdpk") {
				curBook.Title = "Kochb??chlein f??r die Puppen-K??che"
				curBook.TOCPath = "/kfdpk/toc.json"
				curBook.CoverImage = "/kfdpk/cover.jpeg"
			} else {

				if strings.Contains(curBook.Title, "dkut") {
					curBook.Title = "Der Kaffee- und Theetisch"
					curBook.TOCPath = "/dkut/toc.json"
					curBook.CoverImage = "/dkut/cover.jpeg"

				}
			}
		}
		for _, line := range records {
			c := TOCLine{
				CategoryNo: stringToInt(line[0]),
				Category:   line[1],
				Number:     stringToInt(line[2]),
				Title:      line[3],
				Page:       stringToInt(line[4]),
				Image:      line[5],
				Offset:     stringToInt(line[6]),
			}
			content = append(content, c)
		}
		bookContent := content
		contentUniq := unique(content)
		//GetChapters from Content
		/*type ChapterItem struct {
		ID         int
		Title      string
		Note       string
		Page       int
		Begin      int
		Subchapter []SubChapterItem
		*/
		for _, c := range contentUniq {
			curChapter := ChapterItem{
				ID:     c.CategoryNo,
				Title:  c.Category,
				Page:   c.Page,
				Begin:  c.Page,
				Offset: c.Offset,
				Image:  c.Image + ".jpeg",
			}
			curBook.Chapter = append(curBook.Chapter, curChapter)
		}

		for i := range curBook.Chapter {
			curChapter := &curBook.Chapter[i]
			//Search Through BookContents and Extract SubChapters
			for _, c := range bookContent {
				if c.CategoryNo == curChapter.ID {
					sc := SubChapterItem{
						ID:      c.Number,
						Title:   c.Title,
						Page:    c.Page,
						Begin:   c.Page,
						Image:   c.Image + ".jpeg",
						Content: c.Image + ".txt",
						Note:    c.Title,
					}
					curChapter.Subchapter = append(curChapter.Subchapter, sc)
				}
			}

		}

		//add book to booklist
		bookList = append(bookList, curBook)
	}
	b, err := json.Marshal(bookList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
