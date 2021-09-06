package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	infile = flag.String("infile", "", "Input file")
	data   = flag.String("data", "data", "Data output directory")
	// 640w
	sizeRE = regexp.MustCompile(`\d{3}w`)
	// https://scontent-lga3-2.cdninstagram.com/v/t51.2885-15/e15/12093244_861106520603875_155176226_n.jpg?_nc_ht=scontent-lga3-2.cdninstagram.com&_nc_cat=107&_nc_ohc=-bbIT4BxFOsAX93n1cA&edm=APU89FABAAAA&ccb=7-4&oh=13ff31d4eaa52143d1c26e74144f784a&oe=613DE971&_nc_sid=86f79a
	imageRE = regexp.MustCompile(`[0-9a-z_]+\.jpg`)
)

type post struct {
	Note string
	URL  string
	File string
}

// https://golangcode.com/download-a-file-from-a-url/
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func fileExists(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}

func realMain() error {
	b, err := ioutil.ReadFile(*infile)
	if err != nil {
		return err
	}
	s := string(b)
	lines := strings.Split(s, "\n")
	var note string
	var posts []post
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		if len(parts) < 3 {
			continue
		}
		if size := parts[1]; sizeRE.MatchString(size) {
			if size == "640w" {
				url := strings.Join(parts[2:], " ")
				file := imageRE.FindStringSubmatch(url)[0]
				p := post{
					Note: note,
					URL:  url,
					File: file,
				}
				log.Printf("post: %v", p)
				posts = append(posts, p)
			}
		} else {
			note = strings.Join(parts[1:], " ")
		}
	}
	if err := os.MkdirAll(*data, 0755); err != nil {
		return err
	}
	for _, p := range posts {
		f := p.File
		u := p.URL
		outfile := path.Join(*data, f)
		if fileExists(outfile) {
			log.Printf("%s exists", outfile)
			continue
		}
		log.Printf("downloading %s -> %s", u, outfile)
		if err := downloadFile(outfile, u); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if err := realMain(); err != nil {
		panic(err)
	}
}
