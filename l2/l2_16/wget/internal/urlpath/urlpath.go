package urlpath

import (
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

// LocalPath maps a remote URL to a relative file path inside the mirror directory.
// Host name becomes the top-level folder; directory URLs map to index.html.
func LocalPath(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	host := u.Hostname()
	if host == "" {
		return "", err
	}

	p := u.Path
	if p == "" || p == "/" {
		return filepath.Join(host, "index.html"), nil
	}

	p = path.Clean(p)
	if strings.HasSuffix(u.Path, "/") {
		return filepath.Join(host, filepath.FromSlash(p), "index.html"), nil
	}

	ext := strings.ToLower(path.Ext(p))
	if ext == "" {
		return filepath.Join(host, filepath.FromSlash(p), "index.html"), nil
	}

	local := filepath.Join(host, filepath.FromSlash(p))
	if u.RawQuery != "" {
		local = local + "_" + sanitizeQuery(u.RawQuery)
	}

	return local, nil
}

func sanitizeQuery(q string) string {
	const maxLen = 64
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
		"&", "_",
		"=", "_",
	)
	s := replacer.Replace(q)
	if len(s) > maxLen {
		s = s[:maxLen]
	}
	return s
}

// RelLink returns a relative path from fromFile to toFile for use in HTML/CSS rewrites.
func RelLink(fromFile, toFile string) string {
	fromDir := filepath.Dir(fromFile)
	rel, err := filepath.Rel(fromDir, toFile)
	if err != nil {
		return toFile
	}
	return filepath.ToSlash(rel)
}
