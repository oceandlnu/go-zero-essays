package component

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
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

func TestReleaseGoIdWithRequestId(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genSql()
		})
	}
}

func genSql() {
	// 定义前缀部分
	prefix := "INSERT INTO ocean_test.user_info_ext (name, email, age, bio) VALUES "

	// 用于存储所有的 INSERT 语句
	var statements []string

	largeBio := strings.Repeat("This is a long bio content. ", 50) // 生成一个较大的 bio 内容

	// 生成 1000 条数据
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("User%d", i)
		email := fmt.Sprintf("user%d@example.com", i)
		age := i % 100
		bio := fmt.Sprintf("%s - Additional info for user %d", largeBio, i) // 确保每条记录的 bio 都稍有不同

		value := fmt.Sprintf("('%s', '%s', %d, '%s')", name, email, age, bio)
		statement := fmt.Sprintf("%s%s;SELECT SLEEP(6000);", prefix, value)
		statements = append(statements, statement)
	}

	// 将所有的 INSERT 语句连接起来，用换行分隔
	sql := strings.Join(statements, "")

	// 写入到文件 tmp.sql
	file, err := os.Create("tmp.sql")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(sql)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("SQL written to tmp.sql")
}

func selectBreak() {
	ticker := time.NewTicker(time.Second)
	num := 0
	for {
		fmt.Println("start at")
		select {
		case tc := <-ticker.C:
			if num == 5 {
				break
			}
			fmt.Println("Tick at", tc)
			num++
		}
		fmt.Println("end at")
	}
}
