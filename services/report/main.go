package report

import (
	"api-ngmi/lib/cache"
	"api-ngmi/models"
	"api-ngmi/services/s3"
	"log"

	"encoding/json"
)

var ReportCacheKey []byte = []byte("published_reports")
var ReportCacheExpiry int = 5 * 60 // 5 minutes

func MapReport(vs []models.Report, f func(models.Report) models.Report) []models.Report {
	vsm := make([]models.Report, len(vs))
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

// func TransformToS3Urls(c models.Report) models.Report {
// 	c.Logo = Map(c.Logos, s3.GetPublicReadUrl)

// 	return c
// }

func TransformToS3Urls(c models.Report) models.Report {
	c.LogoUrl = s3.AssetsBucket.GetPublicReadUrl(c.Logo)

	return c
}

func MapToS3Urls(cs []models.Report) []models.Report {
	cs = MapReport(cs, TransformToS3Urls)

	return cs
}

func GetPublishedReports() []models.Report {
	var reports []models.Report

	got, err := cache.LocalCache.Get(ReportCacheKey)
	if err != nil {
		reports = models.GetPublishedReports(models.ReportFilterParams{})

		encodedReports, errMarshal := json.Marshal(reports)
		if errMarshal != nil {
			log.Println("encode error:", err)
		}

		cache.LocalCache.Set(ReportCacheKey, encodedReports, ReportCacheExpiry)
	} else {

		err = json.Unmarshal(got, &reports)
		if err != nil {
			log.Println("decode error:", err)
		}
	}

	return reports
}
