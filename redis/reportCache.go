package redis

import (
	"api-ngmi/models"
	"context"
	"time"

	"github.com/go-redis/cache/v9"
)

type PublishedReportsCache struct {
	Reports []models.Report
}

var ReportCacheKey string = "published_reports"
var ReportCacheExpiry time.Duration = time.Duration(5) * time.Minute // 5 minutes

func CachePublishedReports(reports []models.Report) (bool, error) {
	ctx := context.TODO()
	obj := &PublishedReportsCache{
		Reports: reports,
	}

	if err := Cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   ReportCacheKey,
		Value: obj,
		TTL:   ReportCacheExpiry,
	}); err != nil {
		return false, err
	}

	return true, nil
}

func PublishedReportsFromCache() ([]models.Report, error) {
	ctx := context.TODO()

	var wanted PublishedReportsCache
	if err := Cache.Get(ctx, ReportCacheKey, &wanted); err != nil {
		return nil, err
	}

	return wanted.Reports, nil
}
