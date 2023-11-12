package entity

import (
	"github.com/uptrace/bun"
)

type WriterBookRelation struct {
	bun.BaseModel `bun:"table:books"`

	Id              int    `json:"id" bun:",pk,autoincrement"`
	Title           string `json:"title" bun:",notnull"`
	PublicationYear int    `json:"publication_year" bun:",notnull"`
}
