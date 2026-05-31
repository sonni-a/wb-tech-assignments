package downloader

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Result holds a successful HTTP response body and metadata.
type Result struct {
	URL         string
	Body        []byte
	ContentType string
}

// Client performs HTTP GET requests with a timeout.
type Client struct {
	http *http.Client
}

// New creates a downloader with the given request timeout.
func New(timeout time.Duration) *Client {
	return &Client{
		http: &http.Client{
			Timeout: timeout,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 10 {
					return fmt.Errorf("too many redirects")
				}
				return nil
			},
		},
	}
}

// Fetch downloads the resource at rawURL.
func (c *Client) Fetch(rawURL string) (*Result, error) {
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("User-Agent", "wget-mirror/1.0")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", rawURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("GET %s: unexpected status %s", rawURL, resp.Status)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if err != nil {
		return nil, fmt.Errorf("read body %s: %w", rawURL, err)
	}

	ct := resp.Header.Get("Content-Type")
	return &Result{
		URL:         rawURL,
		Body:        body,
		ContentType: ct,
	}, nil
}
