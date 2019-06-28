package dom

import (
	"golang.org/x/net/html"
)

func _getElementsByTagNameAndClass(e *html.Node, tag_name string, class string, result *[]*html.Node) {
	if tag_name == "" || e.Data == tag_name {
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
		_getElementsByTagNameAndClass(child, tag_name, class, result)
	}
}

func GetElementsByTagNameAndClass(e *html.Node, tag_name string, class string) []*html.Node {
	result := []*html.Node{}
	_getElementsByTagNameAndClass(e, tag_name, class, &result)
	return result
}

func GetElementsByTagName(e *html.Node, tag_name string) []*html.Node {
	result := []*html.Node{}
	_getElementsByTagNameAndClass(e, tag_name, "", &result)
	return result
}

func GetElementsByClass(e *html.Node, class string) []*html.Node {
	result := []*html.Node{}
	_getElementsByTagNameAndClass(e, "", class, &result)
	return result
}
