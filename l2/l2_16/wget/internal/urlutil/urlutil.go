package urlutil

import (
	"net/url"
	"path"
	"strings"
)

// Resolve converts a reference URL against a base page URL.
// Returns empty string for non-fetchable schemes (mailto, javascript, data, etc.).
func Resolve(baseURL, ref string) string {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return ""
	}

	if strings.HasPrefix(ref, "#") {
		return ""
	}

	lower := strings.ToLower(ref)
	switch {
	case strings.HasPrefix(lower, "mailto:"),
		strings.HasPrefix(lower, "javascript:"),
		strings.HasPrefix(lower, "data:"),
		strings.HasPrefix(lower, "tel:"):
		return ""
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	parsed, err := url.Parse(ref)
	if err != nil {
		return ""
	}

	resolved := base.ResolveReference(parsed)
	resolved.Fragment = ""

	if resolved.Scheme != "http" && resolved.Scheme != "https" {
		return ""
	}

	return resolved.String()
}

// SameHost reports whether both URLs belong to the same host (case-insensitive).
func SameHost(a, b string) bool {
	ua, err := url.Parse(a)
	if err != nil {
		return false
	}
	ub, err := url.Parse(b)
	if err != nil {
		return false
	}
	return strings.EqualFold(ua.Host, ub.Host)
}

// IsHTMLContentType reports whether the Content-Type header indicates HTML.
func IsHTMLContentType(contentType string) bool {
	ct := strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
	return ct == "text/html" || ct == "application/xhtml+xml"
}

// IsCSSContentType reports whether the Content-Type header indicates CSS.
func IsCSSContentType(contentType string) bool {
	ct := strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
	return ct == "text/css"
}

// LooksLikeHTMLPath guesses HTML documents from URL path when Content-Type is missing.
func LooksLikeHTMLPath(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	p := u.Path
	if p == "" || strings.HasSuffix(p, "/") {
		return true
	}

	ext := strings.ToLower(path.Ext(p))
	switch ext {
	case "", ".html", ".htm", ".xhtml":
		return true
	default:
		return false
	}
}

// NormalizeURL returns a canonical form used for deduplication.
func NormalizeURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	u.Fragment = ""
	u.Path = path.Clean(u.Path)
	if u.Path == "/" {
		u.Path = "/"
	} else if strings.HasSuffix(u.Path, "/") {
		// keep trailing slash for directories
	} else if u.Path == "" {
		u.Path = "/"
	}

	return u.String()
}
