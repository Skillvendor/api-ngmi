package report

import (
	"api-ngmi/models"
	"api-ngmi/redis"
	"api-ngmi/services/s3"
)

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
	c.AssetsUrls = Map(c.Assets, s3.AssetsBucket.GetPublicReadUrl)

	return c
}

func MapToS3Urls(cs []models.Report) []models.Report {
	cs = MapReport(cs, TransformToS3Urls)

	return cs
}

func AddAvgScore(c models.Report) models.Report {
	sum := 0.0
	for _, score := range c.Scores {
		sum += score.Content
	}

	c.AverageScore = sum / float64(len(c.Scores))
	return c
}

func ProcessReport(report models.Report) models.Report {
	report = TransformToS3Urls(report)
	report = AddAvgScore(report)

	return report
}

func ProcessReports(reports []models.Report) []models.Report {
	return MapReport(reports, ProcessReport)
}

func GetPublishedReports() []models.Report {
	var reports []models.Report

	reports, err := redis.PublishedReportsFromCache()
	if err != nil || reports == nil {
		reports = models.GetPublishedReports()

		redis.CachePublishedReports(reports)
	}

	return reports
}
