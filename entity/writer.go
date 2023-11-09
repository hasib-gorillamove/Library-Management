package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Writer struct {
	bun.BaseModel `bun:"table:writer"`

	Name        string     `json:"name" bun:"name"`
	Nationality string     `json:"nationality" bun:"nationality"`
	Age         int        `json:"age" bun:"age"`
	Address     string     `json:"address" bun:"address"`
	CreatedAt   time.Time  `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdateAt    *time.Time `json:"update_at" bun:",nullzero"`
	DeletedAt   *time.Time `json:"-" bun:",soft_delete"`
	CreatedBy   *string    `json:"created_by" bun:"type:uuid,default:uuid_generate_v4()"`
	UpdatedBy   *string    `json:"updated_by" bun:"type:uuid,default:uuid_generate_v4()"`
}

func (p *Writer) Validate() []FieldError {
	return validate(p)
}

type ListAllWriterResponse struct {
	Total  int      `json:"total"`
	Page   int      `json:"page"`
	Writer []Writer `json:"writer"`
}
type WriterFilter struct {
	Keyword     string `query:"keyword"`
	Name        string `query:"name"`
	Nationality string `query:"nationality"`
	PaginationRequest
}
