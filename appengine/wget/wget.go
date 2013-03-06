package wget

import (
	"appengine"
	"appengine/urlfetch"
	"io/ioutil"
	"net/http"
)

func WgetGae(w http.ResponseWriter, context appengine.Context, url string) []byte {
	client := urlfetch.Client(context)
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return []byte("")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

var proxyCache map[string][]byte

func WgetGaeCached(w http.ResponseWriter, context appengine.Context, url string) []byte {
	if proxyCache == nil {
		proxyCache = make(map[string][]byte)
	}
	cachedResult, hit := proxyCache[url]
	if hit {
		return cachedResult
	}

	data := WgetGae(w, context, url)
	proxyCache[url] = data
	return data
}
