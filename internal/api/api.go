package api

import "github.com/jeffersonformigoni/tech-go-react-server/internal/store/pgstore"

type apihander struct {
	q *pgstore.Queries
}