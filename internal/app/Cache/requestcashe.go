package cache

import (
	"sync"
	"time"
)

type ResponceCacheData struct {
	ResponceTime time.Time
	ResponceData string
}

type CacheData struct {
	AvailableData map[string]*ResponceCacheData
	CacheMutex    sync.Mutex
}

func (data *CacheData) ContainsRequestInCache(request, value string) (bool, string) {
	data.CacheMutex.Lock()
	defer data.CacheMutex.Unlock()
	if response, isOk := data.AvailableData[request+"/"+value]; isOk {
		return true, response.ResponceData
	}
	return false, ""
}

func createResponseCacheData(responseData string) *ResponceCacheData {
	currentTime := time.Now()

	return &ResponceCacheData{
		ResponceTime: currentTime,
		ResponceData: responseData,
	}
}

func (data *CacheData) AddDataInCache(request, value, response string) {
	createdResponceCache := createResponseCacheData(response)

	data.CacheMutex.Lock()
	defer data.CacheMutex.Unlock()

	data.AvailableData[request+"/"+value] = createdResponceCache

}
