package std

import (
	"errors"
	"fmt"
	"time"

	"gitlab.wallstcn.com/baoer/matrix/xgbkb/std/redislogger"
	"gopkg.in/redis.v5"
)

type MutexInterface interface {
	Lock() (bool, error)
	UnLock() error
}

// =====================================================================================================================

type SimpleRedisMutex struct {
	RedisClient *redis.Client
	Key         string
	LockTime    time.Duration
}

func NewSimpleRedisMutex(key string, lockTime time.Duration, rc *redis.Client) *SimpleRedisMutex {
	simpleRedisMutex := new(SimpleRedisMutex)
	simpleRedisMutex.Key = key
	simpleRedisMutex.LockTime = lockTime
	simpleRedisMutex.RedisClient = rc
	return simpleRedisMutex
}

func (self *SimpleRedisMutex) Lock() (bool, error) {
	if self.RedisClient == nil {
		redislogger.Errf("SimpleRedisMutex.Lock, RedisClient is null, key: %s, timeout: %v", self.Key, self.LockTime)
		return false, errors.New("SimpleRedisMutex.Lock, RedisClient is null")
	}
	success, err := self.RedisClient.SetNX(self.GetCacheKey(), "Locked", self.LockTime).Result()
	if err != nil {
		return false, err
	}
	return success, nil
}

func (self *SimpleRedisMutex) UnLock() error {
	if self.RedisClient == nil {
		return nil
	}
	_, err := self.RedisClient.Del(self.GetCacheKey()).Result()
	if err != nil {
		redislogger.Errf("SimpleRedisMutex.UnLock, err: %v", err)
		return err
	}
	return nil
}

func (self *SimpleRedisMutex) GetCacheKey() string {
	return fmt.Sprintf("mutex:%s", self.Key)
}
