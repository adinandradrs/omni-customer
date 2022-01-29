package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConfigDatabase(url string, host string, port string, user string, passwd string, sch string) *pgxpool.Pool {
	url = strings.Replace(url, "{{host}}", host, -1)
	url = strings.Replace(url, "{{port}}", port, -1)
	url = strings.Replace(url, "{{username}}", user, -1)
	url = strings.Replace(url, "{{password}}", passwd, -1)
	url = strings.Replace(url, "{{schema}}", sch, -1)
	dbase, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		fmt.Println("error when try to open database ", err)
		panic(err)
	}
	return dbase
}
