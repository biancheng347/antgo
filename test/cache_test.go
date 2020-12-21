package test

import (
	"github.com/small-ek/antgo/conv"
	"github.com/small-ek/antgo/os/cache"
	"log"
	"testing"
)

func TestCache(t *testing.T) {

	for i := 0; i < 10; i++ {
		cache.Sets("ek" + conv.String(i))
		var result = cache.Get("ek" + conv.String(i))

		log.Println(string(result))
	}
}
