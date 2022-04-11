package collection

import (
	"go-rarity/models"
	"go-rarity/services/s3"
)

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
