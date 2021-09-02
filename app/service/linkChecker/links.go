package linkChecker

import (
	"strings"

	"golang.org/x/net/html"
)

func getDomain(link string) string {
	s := strings.Split(link, "/")
	return s[0] + "//" + s[2]
}

func fixProtocolPrefix(link string) string {
	if !strings.Contains(link, "://") {
		link = "http://" + link
	}
	return link
}

func getLinks(htmlTokens *html.Tokenizer) (newLink []string) {
	for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := htmlTokens.Token()
			if t.Data == "a" {
				newLink = append(newLink, t.Attr[0].Val)
			}
		}
	}
}