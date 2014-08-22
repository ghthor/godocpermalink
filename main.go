package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"
)

func genPermalink(link string, goversion string) (permalink string, err error) {
	u, err := url.Parse(link)
	if err != nil {
		return
	}

	u.Scheme = "https"
	u.Host = "code.google.com"
	u.Path = path.Join("/p/go/source/browse", u.Path)

	q := u.Query()
	q.Del("s")
	q.Set("name", goversion)
	u.RawQuery = q.Encode()

	u.Fragment = strings.TrimPrefix(u.Fragment, "L")

	permalink = u.String()

	return
}

func main() {
	flag.Parse()

	golangUrl := flag.Args()[0]

	permalink, err := genPermalink(golangUrl, "go1.3.1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(permalink)
}
