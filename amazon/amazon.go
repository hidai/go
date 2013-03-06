package amazon

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func calcBase64HmacSha256(in string, secret_key string) string {
	var out [sha256.Size]byte
	h := hmac.New(sha256.New, []byte(secret_key))
	h.Write([]byte(in))
	h.Sum(out[:0])
	return base64.StdEncoding.EncodeToString(out[:])
}

func IsValidOrder(order string) bool {
	return order == "salesrank" ||
		order == "pricerank" ||
		order == "inverse-pricerank" ||
		order == "daterank" ||
		order == "titlerank" ||
		order == "-titlerank"
}

type KeyAndTag struct {
	PublicKey string
	SecretKey string
	Tag       string
}

func GetBrowseNodeLookupUrl(browse_node_id int, key_and_tag KeyAndTag) string {
	host := "ecs.amazonaws.jp"
	path := "/onca/xml"
	timestamp := time.Now().Format(time.RFC3339)
	params := "" +
		"AWSAccessKeyId=" + key_and_tag.PublicKey +
		"&AssociateTag=" + key_and_tag.Tag +
		"&BrowseNodeId=" + strconv.Itoa(browse_node_id) +
		"&ContentType=" + url.QueryEscape("text/xml") +
		"&Operation=BrowseNodeLookup" +
		//"&ResponseGroup=" + url.QueryEscape("MostGifted,NewReleases,MostWishedFor,TopSellers") +
		"&Timestamp=" + url.QueryEscape(timestamp) +
		"&Version=2011-08-01"
	string_to_sign := "GET\n" + host + "\n" + path + "\n" + params
	sign := calcBase64HmacSha256(string_to_sign, key_and_tag.SecretKey)
	signed_url := "http://" + host + path + "?" + params + "&Signature=" + url.QueryEscape(sign)
	return signed_url
}

func GetItemSearchUrl(query string, key_and_tag KeyAndTag, page int, order string) string {
	// Somehow Amazon API doesn't accept '+' escaping. Replacing + to %20.
	escaped_query := url.QueryEscape(query)
	escaped_query = strings.Replace(escaped_query, "+", "%20", -1)

	sort_line := ""
	if IsValidOrder(order) {
		sort_line = "&Sort=" + order
	}

	host := "ecs.amazonaws.jp"
	path := "/onca/xml"
	timestamp := time.Now().Format(time.RFC3339)
	params := "" +
		"AWSAccessKeyId=" + key_and_tag.PublicKey +
		"&AssociateTag=" + key_and_tag.Tag +
		//"&Availability=Available" +
		"&ContentType=" + url.QueryEscape("text/xml") +
		"&ItemPage=" + strconv.Itoa(page) +
		"&Keywords=" + escaped_query +
		//"&MerchantId=Amazon" +
		"&MinimumPrice=1" +
		"&Operation=ItemSearch" +
		"&ResponseGroup=" + url.QueryEscape("ItemAttributes,Images,Reviews") +
		"&SearchIndex=Books" +
		"&Service=AWSECommerceService" +
		sort_line +
		"&Timestamp=" + url.QueryEscape(timestamp) +
		"&Version=2011-08-01"
	string_to_sign := "GET\n" + host + "\n" + path + "\n" + params
	sign := calcBase64HmacSha256(string_to_sign, key_and_tag.SecretKey)
	signed_url := "http://" + host + path + "?" + params + "&Signature=" + url.QueryEscape(sign)
	return signed_url
}

type Image struct {
	URL    string
	Height int
	Width  int
}
type ItemAttributes struct {
	Title  string
	Author string
}
type CustomerReviews struct {
	IFrameURL  string
	HasReviews bool
}
type Item struct {
	ASIN            string
	ItemAttributes  ItemAttributes
	DetailPageURL   string
	LargeImage      Image
	CustomerReviews CustomerReviews
}
type Items struct {
	Item       []Item
	TotalPages int
}
type ItemSearchResponse struct {
	XMLName xml.Name `xml:"ItemSearchResponse"`
	Items   Items
}

type Children struct {
	BrowseNode []BrowseNode
}
type Ancestors struct {
	BrowseNode []BrowseNode
}
type BrowseNode struct {
	BrowseNodeId   int
	Name           string
	IsCategoryRoot int
	Children       Children
	Ancestors      Ancestors
}
type BrowseNodes struct {
	BrowseNode BrowseNode
}
type BrowseNodeLookupResponse struct {
	XMLName     xml.Name `xml:"BrowseNodeLookupResponse"`
	BrowseNodes BrowseNodes
}
