package service

import (
	"database/sql"
	"github.com/hako/branca"
)

// This package contains the core domain logic
// we can change the transport method to grpc or graphql
// instead of using REST

type Service struct {
	db    *sql.DB
	codec *branca.Branca
}
