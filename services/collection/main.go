package collection

import (
	"go-rarity/lib/cache"
	"go-rarity/models"
	"go-rarity/services/s3"
	"log"

	"encoding/json"
)

var CollectionCacheKey []byte = []byte("published_collections")
var CollectionCacheExpiry int = 5 * 60 // 5 minutes

func MapCollection(vs []models.Collection, f func(models.Collection) models.Collection) []models.Collection {
	vsm := make([]models.Collection, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func TransformToS3Urls(c models.Collection) models.Collection {
	c.Logos = Map(c.Logos, s3.GetPublicReadUrl)
	c.Research = Map(c.Research, s3.GetPublicReadUrl)

	return c
}

func MapToS3Urls(cs []models.Collection) []models.Collection {
	cs = MapCollection(cs, TransformToS3Urls)

	return cs
}

func GetPublishedCollections() []models.Collection {
	// var network bytes.Buffer
	var collections []models.Collection

	got, err := cache.LocalCache.Get(CollectionCacheKey)
	if err != nil {
		collections = models.GetPublishedCollections()

		encodedCollections, errMarshal := json.Marshal(collections)
		if errMarshal != nil {
			log.Println("encode error:", err)
		}

		cache.LocalCache.Set(CollectionCacheKey, encodedCollections, CollectionCacheExpiry)
	} else {

		err = json.Unmarshal(got, &collections)
		if err != nil {
			log.Println("decode error:", err)
		}
	}

	return collections
}
