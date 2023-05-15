package setup

import (
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

var MCACHE *memcache.Client

func RunMemChache() {
	mc := memcache.New(os.Getenv("MEMCACHE_HOST"))

	MCACHE = mc
}
