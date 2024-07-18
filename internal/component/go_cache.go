package component

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func getGoroutineId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

var cacheMap = cache.New(3*time.Minute, 5*time.Minute)

func BindGoIdWithRequestId(requestId string) {
	if requestId == "" {
		requestId = uuid.New().String()
	}
	goId := getGoroutineId()
	cacheMap.Set(strconv.Itoa(goId), requestId, cache.DefaultExpiration)
}

func ReleaseGoIdWithRequestId() {
	goId := getGoroutineId()
	cacheMap.Delete(strconv.Itoa(goId))
}

func GetRequestId() string {
	goId := getGoroutineId()
	if requestId, ok := cacheMap.Get(strconv.Itoa(goId)); ok {
		return requestId.(string)
	}
	return ""
}
