//+build integration

package rdb

import "testing"

var (
	testRepo *MySQLRepository
)


func TestMain(m *testing.M) {
	testRepo = NewMySQLRepository(Config{
		DSN:             "root:root@tcp(127.0.0.1:3306)/test?parseTime=true",
		ConnMaxLifetime: "30m",
		MaxIdleConns:    "10",
		SetMaxOpenConns: "20",
	})
	defer testRepo.Close()
	m.Run()
}
