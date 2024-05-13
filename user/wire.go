package user

/* build wiredinject*/

import (
	"database/sql"
	"github.com/google/wire"
)

func Wire(db *sql.DB) *hanlder {
	wire.Build(ProviderSet)
	return nil
}
