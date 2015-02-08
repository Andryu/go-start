package main

/*
   API 認証
   signature : 署名コード
   message   : 署名元メッサージ
*/
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/url"
)

type Authentication struct {
	RedisKey string
}

type AuthenticateInterface interface {
	Authenticate(signature string, message string)
}

func (a *Authentication) SetKeyForApi(level string) {
	a.RedisKey = "api:" + level
}

func (a *Authentication) SetKeyForApp(token string) {
	a.RedisKey = "application:" + token
}

func (a *Authentication) Authenticate(signature string, message string) bool {
	c := pool.Get()
	defer c.Close()

	k, err := redis.String(c.Do("GET", a.RedisKey))
	if err != nil {
		return false
	}
	fmt.Println(hmacSha256(k, message))
	if signature != hmacSha256(k, message) {
		return false
	}
	return true
}

// HMAC SHA256
func hmacSha256(key string, message string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(message))
	return url.QueryEscape(base64.URLEncoding.EncodeToString(hash.Sum(nil)))
}
