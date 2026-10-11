package simpleconnection

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.TODO()

	conn, err := pgx.Connect(ctx, "postgres://postgres:86850@localhost:5432/postgres")
	if err != nil {
		log.Fatal("database connecting fail", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("database ping fail", err)
	}

	fmt.Println("DB connected")
}
