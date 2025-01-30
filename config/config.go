package config

import (
    "fmt"
    "os"

    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
    dsn := "postgres://user:password@localhost:5432/controle_trafego?sslmode=disable"
    var err error
    DB, err = sqlx.Connect("postgres", dsn)
    if err != nil {
        fmt.Println("Erro ao conectar no banco:", err)
        os.Exit(1)
    }
}