package simpleconnection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(ctx context.Context, conn string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, conn)
}
