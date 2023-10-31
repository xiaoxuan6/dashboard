package tests

import (
    "dashboard/pkg/redis"
    "github.com/joho/godotenv"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

// REDIS_DB=2 go test tests/redis_test.go -v
func TestRedis(t *testing.T) {
    _ = godotenv.Load()

    err := redis.Set("go_key", "go", 1*time.Minute)
    assert.Nil(t, err)

    //res, err := redis.Get("go_key")
    //assert.Nil(t, err)
    //assert.Equal(t, "go", res)
    //t.Log("res", res)

    //err = redis.Del("go_key")
    //assert.Nil(t, err)

    //************* Hash *******************//
    //err = redis.HSet("go_hash_key1", "字符串3", "字符串4")
    //assert.Nil(t, err)
    //res, err := redis.HGet("go_hash_key1", "字符串3")
    //assert.Nil(t, err)
    //assert.Equal(t, "字符串4", res)
    //
    //result, err := redis.HGetAll("go_hash_key1")
    //assert.Nil(t, err)
    //t.Log("result", result)

    //hashData := map[string]interface{}{
    //    "field2": "value2",
    //    "field3": "value3",
    //    "field4": "value3",
    //}
    //b, _ := json.Marshal(hashData)
    //err = redis.HSet("go_hash_key1", "字符串1", string(b))
    //assert.Nil(t, err)
    //
    //res, err := redis.HMGet("go_hash_key1", "字符串1")
    //assert.Nil(t, err)
    //data := make(map[string]interface{})
    //_ = json.Unmarshal([]byte(res[0].(string)), &data)
    //t.Log("res", res)
    //t.Log("res", data)

    //status := redis.Exists("go_hash_key1")
    //assert.Equal(t, false, status)

    result, err := redis.Remember("go_hash_key11", 1*time.Minute, func() []string {
        return []string{"value1", "value3"}
    })
    assert.Nil(t, err)
    t.Log("result", result)

    // HMSet
    //err = redis.HMSet("go_hash_key2", "字符串1", "字符串2", "字符串3", "字符串4")
    //assert.Nil(t, err)
    //
    //item, err := redis.HMGet("go_hash_key2", "字符串1", "字符串3")
    //assert.Nil(t, err)
    //t.Log(item)
    //
    //item1, err := redis.HGetAll("go_hash_key2")
    //assert.Nil(t, err)
    //t.Log("item1", item1)
}
