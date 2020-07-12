package jsonl

import (
	"database/sql"
	"github.com/genhoi/users/module/db"
	"github.com/genhoi/users/module/user"
	"io"
	"os"
)

type Importer struct {
	query *db.Query
	db    *sql.DB
}

func NewImporter(query *db.Query, db *sql.DB) *Importer {
	return &Importer{
		query: query,
		db:    db,
	}
}

func (i *Importer) Import(path ...string) error {
	readers := make([]io.Reader, len(path))

	for i, filePath := range path {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		readers[i] = file
	}

	mReader := io.MultiReader(readers...)

	chunkSize := 10
	users := Generator(mReader, chunkSize)
	chunks := user.ChunkGenerate(users, chunkSize)
	for chunk := range chunks {
		var convertedChunk []interface{}
		for _, result := range chunk {
			if result.Err != nil {
				return result.Err
			}
			convertedChunk = append(convertedChunk, result.User)
			query, values, err := i.query.Upsert(convertedChunk)
			if err != nil {
				return err
			}

			_, err = i.db.Exec(query, values...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
