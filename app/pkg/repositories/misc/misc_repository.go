package misc

import "database/sql"

type MiscRepository struct {
	DB *sql.DB
}

func NewMiscRepository(db *sql.DB) MiscRepository {
	return MiscRepository{db}
}
