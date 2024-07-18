package component

import (
	"github.com/google/uuid"
	"math/rand"
	"testing"
)

func BenchmarkGoCacheUse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		use()
	}
}

func use() {
	requestId := uuid.New().String()
	BindGoIdWithRequestId(requestId)
	n := rand.Intn(200)
	for i := 0; i < n; i++ {
		GetRequestId()
	}
	ReleaseGoIdWithRequestId()
}
