package crawler

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"wget/internal/downloader"
	"wget/internal/htmlparser"
	"wget/internal/urlpath"
	"wget/internal/urlutil"
)

// Config controls mirroring behavior.
type Config struct {
	StartURL string
	Output   string
	Depth    int
	Workers  int
}

type task struct {
	url   string
	depth int
}

// Mirror downloads a site (or part of it) into a local directory.
type Mirror struct {
	cfg      Config
	client   *downloader.Client
	visited  sync.Map
	errCount atomic.Int32
}

// New creates a mirror crawler.
func New(cfg Config, client *downloader.Client) (*Mirror, error) {
	if cfg.Workers <= 0 {
		cfg.Workers = 4
	}
	if cfg.Depth < 0 {
		cfg.Depth = 0
	}

	start, err := url.Parse(cfg.StartURL)
	if err != nil {
		return nil, fmt.Errorf("parse start URL: %w", err)
	}
	if start.Scheme != "http" && start.Scheme != "https" {
		return nil, fmt.Errorf("unsupported URL scheme: %s", start.Scheme)
	}
	if start.Host == "" {
		return nil, fmt.Errorf("URL must include a host")
	}

	return &Mirror{
		cfg:    cfg,
		client: client,
	}, nil
}

// Run mirrors the site using a worker pool.
func (m *Mirror) Run() error {
	if err := os.MkdirAll(m.cfg.Output, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	start := urlutil.NormalizeURL(m.cfg.StartURL)
	queue := make(chan task, m.cfg.Workers*4)
	var jobs sync.WaitGroup
	var workers sync.WaitGroup

	for i := 0; i < m.cfg.Workers; i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for job := range queue {
				m.handle(job, queue, &jobs)
			}
		}()
	}

	if m.markVisited(start) {
		jobs.Add(1)
		queue <- task{url: start, depth: 0}
	}

	go func() {
		jobs.Wait()
		close(queue)
	}()

	workers.Wait()

	if m.errCount.Load() > 0 {
		return fmt.Errorf("completed with %d download errors", m.errCount.Load())
	}
	return nil
}

func (m *Mirror) handle(job task, queue chan<- task, wg *sync.WaitGroup) {
	defer wg.Done()

	next, err := m.download(job)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: %v\n", err)
		m.errCount.Add(1)
		return
	}

	for _, t := range next {
		if m.markVisited(t.url) {
			wg.Add(1)
			queue <- t
		}
	}
}

func (m *Mirror) markVisited(rawURL string) bool {
	_, loaded := m.visited.LoadOrStore(rawURL, true)
	return !loaded
}

func (m *Mirror) download(job task) ([]task, error) {
	if !urlutil.SameHost(job.url, m.cfg.StartURL) {
		return nil, nil
	}

	result, err := m.client.Fetch(job.url)
	if err != nil {
		return nil, err
	}

	localRel, err := urlpath.LocalPath(job.url)
	if err != nil {
		return nil, fmt.Errorf("map path for %s: %w", job.url, err)
	}

	body := result.Body
	isHTML := urlutil.IsHTMLContentType(result.ContentType) || urlutil.LooksLikeHTMLPath(job.url)
	isCSS := urlutil.IsCSSContentType(result.ContentType)

	resolve := urlutil.Resolve
	localPath := urlpath.LocalPath
	relLink := urlpath.RelLink

	if isHTML {
		body, err = htmlparser.RewriteHTML(body, job.url, resolve, localPath, localRel, relLink)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: rewrite HTML %s: %v\n", job.url, err)
			body = result.Body
		}
	} else if isCSS {
		body = htmlparser.RewriteCSS(body, job.url, resolve, localPath, localRel, relLink)
	}

	dest := filepath.Join(m.cfg.Output, localRel)
	if err := writeFile(dest, body); err != nil {
		return nil, err
	}

	fmt.Printf("saved %s -> %s\n", job.url, dest)

	var refs []string
	if isHTML {
		refs = htmlparser.ExtractHTMLRefs(job.url, result.Body, resolve)
	} else if isCSS {
		refs = htmlparser.ExtractCSSRefs(job.url, result.Body, resolve)
	}

	var next []task
	for _, ref := range refs {
		normalized := urlutil.NormalizeURL(ref)
		if !urlutil.SameHost(normalized, m.cfg.StartURL) {
			continue
		}

		childDepth := job.depth
		if isHTML && urlutil.LooksLikeHTMLPath(normalized) && normalized != job.url {
			if job.depth >= m.cfg.Depth {
				continue
			}
			childDepth = job.depth + 1
		}

		next = append(next, task{url: normalized, depth: childDepth})
	}

	return next, nil
}

func writeFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
