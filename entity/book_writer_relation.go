package entity

import (
	"github.com/uptrace/bun"
)

type BookWriterRelation struct {
	bun.BaseModel `bun:"table:books"`

	Id              int    `json:"id" bun:",pk,autoincrement"`
	Title           string `json:"title" bun:",notnull"`
	AuthorId        int    `json:"author_id" bun:",notnull"`
	PublicationYear int    `json:"publication_year" bun:",notnull"`
}
