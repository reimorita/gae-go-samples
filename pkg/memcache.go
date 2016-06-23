package pkg

import (
	"net/http"
	"appengine"
	"appengine/memcache"
	"fmt"
)

func init() {
	http.HandleFunc("/storeMem", storeMem)
}

func storeMem(w http.ResponseWriter, r *http.Request) {
	// Context
	ctx := appengine.NewContext(r)
	// Create Item
	item := &memcache.Item{
		Key:   "myKey",
		Value: []byte("FirstMemcache"),
	}
	// Add the item to the memcache
	if err := memcache.Add(ctx, item); err == memcache.ErrNotStored {
		// Key already exist
		fmt.Fprint(w,  "Key already exist. key is ", item.Key)
		return
	} else if err != nil {
		// Error Handler
	}

	// Get the item from the memcache
	if item, err := memcache.Get(ctx, "myKey"); err == memcache.ErrCacheMiss {
		// Missing cache data
	} else if err != nil {
		// Error Handler
	} else {
		fmt.Fprint(w,  "Value of myKey is ", string(item.Value))
	}
}