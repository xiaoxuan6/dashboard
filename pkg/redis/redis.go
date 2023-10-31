package redis

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/redis/go-redis/v9"
    "os"
    "strconv"
    "sync"
    "time"
)

var (
    one sync.Once
    rdb *redis.Client
    ctx = context.Background()
)

func init() {
    one.Do(func() {
        db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
        rdb = redis.NewClient(&redis.Options{
            Addr:     os.Getenv("REDIS_ADDR"),
            Password: os.Getenv("REDIS_PASSWORD"),
            DB:       db,
        })
    })
}

func Exists(key ...string) bool {
    return rdb.Exists(ctx, key...).Val() > 0
}

func Set(key string, value interface{}, expiration time.Duration) error {
    return rdb.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
    return rdb.Get(ctx, key).Result()
}

func Del(key string) error {
    return rdb.Del(ctx, key).Err()
}

// HSet accepts values in following formats:
//   - HSet("myhash", "key1", "value1")
func HSet(key string, values ...interface{}) error {
    return rdb.HSet(ctx, key, values).Err()
}

// HGet accepts values in following formats:
//   - HGet("myhash", "key1")
func HGet(key string, field string) (string, error) {
    return rdb.HGet(ctx, key, field).Result()
}

// HMSet accepts values in following formats:
//   - HMset("myhash", "key1", "value1", "key2", "value2")
func HMSet(key string, values ...interface{}) error {
    return rdb.HMSet(ctx, key, values).Err()
}

// HMGet accepts values in following formats:
//   - HMGet("myhash", "key1", "key2")
func HMGet(key string, fields ...string) ([]interface{}, error) {
    return rdb.HMGet(ctx, key, fields...).Result()
}

// HGetAll accepts values in following formats:
//   - HGetAll("myhash")
func HGetAll(key string) (map[string]string, error) {
    return rdb.HGetAll(ctx, key).Result()
}

// Remember accepts values in following formats:
//  - Remember("mykey", 10*time.Second, func() []string {
//      return []string{"value1", "value2"}
//    })
func Remember(key string, expiration time.Duration, f func() []string) ([]string, error) {
    hmKey := fmt.Sprintf("%s_hash", key)

    if !Exists(key) {
        sliceStr := f()

        err := Set(key, "1", expiration)
        if err != nil {
            return sliceStr, err
        }

        b, err1 := json.Marshal(sliceStr)
        if err1 != nil {
            return sliceStr, err1
        }

        err = HMSet(hmKey, key, string(b))
        if err != nil {
            return sliceStr, err
        }

        return sliceStr, nil
    } else {
        result, err1 := HMGet(hmKey, key)
        if err1 != nil {
            sliceStr := f()
            return sliceStr, err1
        }

        item := result[0].(string)
        var items []string
        err := json.Unmarshal([]byte(item), &items)
        if err != nil {
            sliceStr := f()
            return sliceStr, err
        }

        return items, nil
    }
}
