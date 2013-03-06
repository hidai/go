package dom

import (
	"code.google.com/p/go-html-transform/h5"
)

func _getElementsByTagNameAndClass(e *h5.Node, tag_name string, class string, result *[]*h5.Node) {
	if tag_name == "" || e.Data() == tag_name {
		if class == "" {
			*result = append(*result, e)
		} else {
			for _, attr := range e.Attr {
				if attr.Name == "class" && attr.Value == class {
					*result = append(*result, e)
				}
			}
		}
	}
	for _, child := range e.Children {
		_getElementsByTagNameAndClass(child, tag_name, class, result)
	}
}

func GetElementsByTagNameAndClass(e *h5.Node, tag_name string, class string) []*h5.Node {
	result := []*h5.Node{}
	_getElementsByTagNameAndClass(e, tag_name, class, &result)
	return result
}

func GetElementsByTagName(e *h5.Node, tag_name string) []*h5.Node {
	result := []*h5.Node{}
	_getElementsByTagNameAndClass(e, tag_name, "", &result)
	return result
}

func GetElementsByClass(e *h5.Node, class string) []*h5.Node {
	result := []*h5.Node{}
	_getElementsByTagNameAndClass(e, "", class, &result)
	return result
}
