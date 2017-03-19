package htmlparser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type container struct {
	label  string
	links  []link
	childs []container
}

type link struct {
	label string
	url   string
}

type Parser struct {
	reader io.Reader
}

func New(input io.Reader) (*Parser, error) {
	p := &Parser{reader: input}

	return p, nil
}

/*func Crawl(url String) []byte {
	url = "foo.bar"

	response, err = http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	bytes := response.Body
	defer bytes.Close() // close Body when the function returns

	return bytes
}*/

func (p *Parser) ParseHtml() []container {
	depth := 0
	z := html.NewTokenizer(p.reader)
	c := make([]container, 0)

	for {
		tt := z.Next()
		var cTemp container

		switch tt {
		case html.ErrorToken:
			// End of the document, we're done
			return c
		case html.TextToken:
			if depth > 0 {
				cTemp.label = string(z.Text())
				fmt.Println("ddsad: " + z.Token().String())

				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				//emitBytes(z.Text())
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()

			if len(tn) == 2 && tn[0] == 'h' && tn[1] == '1' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
					c = append(c, cTemp)
				}
			}

			/*if len(tn) == 1 && tn[0] == 'a' {
				if tt == StartTagToken {
					depth++
				} else {
					depth--
				}
			}*/

			/*case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "h1" {
				cTemp := container{label: t.}
			}

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := p.getHref(t)
			if !ok {
				continue
			}

			// Make sure the url begines in http**
			hasProto := strings.Index(url, "http") == 0
			if hasProto {

				l := link{label: "bar", url: url}

				c := container{label: "foo", links: []link{l}, childs: []container{}}

				ret = append(ret, c)
			}*/
		}
	}
}

// Helper function to pull the href attribute from a Token
func (p *Parser) getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}
