package dom

import (
	"golang.org/x/net/html"
)

func _getElementsByTagNameAndClass(e *html.Node, tagName string, class string, result *[]*html.Node) {
	if tagName == "" || e.Data == tagName {
		if class == "" {
			*result = append(*result, e)
		} else {
			for _, attr := range e.Attr {
				if attr.Key == "class" && attr.Val == class {
					*result = append(*result, e)
				}
			}
		}
	}
	for child := e.FirstChild; child != nil; child = child.NextSibling {
		_getElementsByTagNameAndClass(child, tagName, class, result)
	}
}

func GetElementsByTagNameAndClass(e *html.Node, tagName string, class string) []*html.Node {
	result := []*html.Node{}
	_getElementsByTagNameAndClass(e, tagName, class, &result)
	return result
}

func GetElementsByTagName(e *html.Node, tagName string) []*html.Node {
	result := []*html.Node{}
	_getElementsByTagNameAndClass(e, tagName, "", &result)
	return result
}

func GetElementsByClass(e *html.Node, class string) []*html.Node {
	result := []*html.Node{}
	_getElementsByTagNameAndClass(e, "", class, &result)
	return result
}
