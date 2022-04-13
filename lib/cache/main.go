package cache

import "github.com/coocood/freecache"

var LocalCache *freecache.Cache

func init() {
	var err error

	cacheSize := 100 * 1024 * 256
	LocalCache = freecache.NewCache(cacheSize)

	if err != nil {
		panic(err)
	}
}
