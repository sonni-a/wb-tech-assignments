package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"wget/internal/crawler"
	"wget/internal/downloader"
)

const defaultURL = "https://example.com"

func main() {
	url := flag.String("url", defaultURL, "start URL to mirror")
	output := flag.String("output", "mirror", "output directory")
	depth := flag.Int("depth", 0, "maximum link depth for HTML pages")
	flag.Parse()

	if flag.NArg() > 0 {
		*url = flag.Arg(0)
	}

	client := downloader.New(30 * time.Second)
	m, err := crawler.New(crawler.Config{
		StartURL: *url,
		Output:   *output,
		Depth:    *depth,
		Workers:  4,
	}, client)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	if err := m.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
