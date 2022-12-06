package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
)

type AccessLevelCache struct {
	Address     string
	AccessLevel int
}

var AccessLevelCacheExpiry time.Duration = time.Duration(24) * time.Hour // 6 hours

func CacheAccessLevelFor(address string, accessLevel int) (bool, error) {
	ctx := context.TODO()
	key := fmt.Sprintf("accessLevel-%s", address)
	obj := &AccessLevelCache{
		Address:     address,
		AccessLevel: accessLevel,
	}

	if err := Cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   AccessLevelCacheExpiry,
	}); err != nil {
		fmt.Println("Error getting access level for", address, err)
		return false, err
	}

	return true, nil
}

func AccessLevelFor(address string) (int, error) {
	ctx := context.TODO()
	key := fmt.Sprintf("accessLevel-%s", address)

	var wanted AccessLevelCache
	if err := Cache.Get(ctx, key, &wanted); err != nil {
		return 0, err
	}

	return wanted.AccessLevel, nil
}

func PurgeAccessLevelCacheFor(address string) (bool, error) {
	ctx := context.TODO()
	key := fmt.Sprintf("accessLevel-%s", address)

	if err := Cache.Delete(ctx, key); err != nil {
		return false, err
	}

	return true, nil
}
