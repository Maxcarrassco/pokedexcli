package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)


const BASE_URL = "https://pokeapi.co/api/v2"

var cache = NewCache(20 * time.Second)

func GetRequest[T any](url string, obj *T) error {
	if ch, ok := cache.store[url]; ok {
		json.Unmarshal(ch.val, obj)
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	cache.Add(url, resBody)
	json.Unmarshal(resBody, obj)
	return nil
}
