package wget

import (
	"appengine"
	"appengine/memcache"
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

func WgetGaeCached(w http.ResponseWriter, context appengine.Context, url string) []byte {
	key := "github.com/hidai/go/appengine/wget:" + url

	cachedItem, err := memcache.Get(context, key)
	if err == nil {
		return cachedItem.Value
	}

	data := WgetGae(w, context, url)
	newItem := &memcache.Item{
		Key:   key,
		Value: data,
	}
	memcache.Set(context, newItem)
	return data
}
