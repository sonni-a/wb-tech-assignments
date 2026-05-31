package htmlparser

import (
	"bytes"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var cssURLRe = regexp.MustCompile(`url\(\s*['"]?([^'")]+)['"]?\s*\)`)

// ExtractHTMLRefs returns fetchable URLs from HTML attributes and inline CSS.
func ExtractHTMLRefs(baseURL string, data []byte, resolve func(base, ref string) string) []string {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil
	}

	seen := make(map[string]struct{})
	var refs []string

	add := func(ref string) {
		abs := resolve(baseURL, ref)
		if abs == "" {
			return
		}
		if _, ok := seen[abs]; ok {
			return
		}
		seen[abs] = struct{}{}
		refs = append(refs, abs)
	}

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				switch attr.Key {
				case "href", "src", "action", "poster", "data-src":
					add(attr.Val)
				case "srcset":
					for _, part := range strings.Split(attr.Val, ",") {
						token := strings.TrimSpace(strings.Fields(part)[0])
						add(token)
					}
				case "style":
					for _, u := range extractCSSURLs(attr.Val) {
						add(u)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	return refs
}

// ExtractCSSRefs returns URLs referenced inside a CSS file.
func ExtractCSSRefs(baseURL string, data []byte, resolve func(base, ref string) string) []string {
	seen := make(map[string]struct{})
	var refs []string

	for _, u := range extractCSSURLs(string(data)) {
		abs := resolve(baseURL, u)
		if abs == "" {
			continue
		}
		if _, ok := seen[abs]; ok {
			continue
		}
		seen[abs] = struct{}{}
		refs = append(refs, abs)
	}

	return refs
}

func extractCSSURLs(content string) []string {
	matches := cssURLRe.FindAllStringSubmatch(content, -1)
	urls := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) > 1 {
			urls = append(urls, strings.TrimSpace(m[1]))
		}
	}
	return urls
}

// RewriteHTML replaces remote URLs with local relative paths in HTML.
func RewriteHTML(data []byte, baseURL string, resolve func(base, ref string) string, localPath func(rawURL string) (string, error), fromLocal string, relLink func(from, to string) string) ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return data, err
	}

	rewriteAttr := func(val string) string {
		abs := resolve(baseURL, val)
		if abs == "" {
			return val
		}
		target, err := localPath(abs)
		if err != nil {
			return val
		}
		return relLink(fromLocal, target)
	}

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for i, attr := range n.Attr {
				switch attr.Key {
				case "href", "src", "action", "poster", "data-src":
					n.Attr[i].Val = rewriteAttr(attr.Val)
				case "srcset":
					parts := strings.Split(attr.Val, ",")
					for j, part := range parts {
						fields := strings.Fields(strings.TrimSpace(part))
						if len(fields) == 0 {
							continue
						}
						fields[0] = rewriteAttr(fields[0])
						parts[j] = strings.Join(fields, " ")
					}
					n.Attr[i].Val = strings.Join(parts, ", ")
				case "style":
					n.Attr[i].Val = rewriteCSS(attr.Val, baseURL, resolve, localPath, fromLocal, relLink)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		return data, err
	}
	return buf.Bytes(), nil
}

// RewriteCSS replaces url(...) references with local relative paths.
func RewriteCSS(data []byte, baseURL string, resolve func(base, ref string) string, localPath func(rawURL string) (string, error), fromLocal string, relLink func(from, to string) string) []byte {
	return []byte(rewriteCSS(string(data), baseURL, resolve, localPath, fromLocal, relLink))
}

func rewriteCSS(content, baseURL string, resolve func(base, ref string) string, localPath func(rawURL string) (string, error), fromLocal string, relLink func(from, to string) string) string {
	return cssURLRe.ReplaceAllStringFunc(content, func(match string) string {
		sub := cssURLRe.FindStringSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		raw := strings.TrimSpace(sub[1])
		abs := resolve(baseURL, raw)
		if abs == "" {
			return match
		}
		target, err := localPath(abs)
		if err != nil {
			return match
		}
		rel := relLink(fromLocal, target)
		if strings.ContainsAny(raw, " '\"") {
			return "url('" + rel + "')"
		}
		return "url(" + rel + ")"
	})
}
