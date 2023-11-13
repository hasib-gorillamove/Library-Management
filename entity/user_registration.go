package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type UserRegistration struct {
	bun.BaseModel `bun:"table:user_registration"`

	UserID     int        `json:"user_id"bun:",pk,autoincrement"`
	FirstName  string     `json:"first_name"bun:"first_name"`
	LastName   string     `json:"last_name" bun:"last_name"`
	Occupation string     `json:"occupation"bun:"occupation"`
	Email      string     `json:"email"bun:"email"`
	Password   string     `json:"password"bun:"password"`
	CreatedAt  time.Time  `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdateAt   *time.Time `json:"update_at" bun:",nullzero"`
	DeletedAt  *time.Time `json:"-" bun:",soft_delete"`
	CreatedBy  *string    `json:"created_by" bun:"type:uuid,default:uuid_generate_v4()"`
	UpdatedBy  *string    `json:"updated_by" bun:"type:uuid,default:uuid_generate_v4()"`
}

func (p *UserRegistration) Validate() []FieldError {
	return validate(p)
}

type UserFilter struct {
	Keyword   string `query:"keyword"`
	FirstName string `query:"first_name"`
	//LastName   string `query:"last_name"`
	//Occupation string `query:"occupation"`
	PaginationRequest
}

type GetAllUserResponses struct {
	Total int                `json:"total"`
	Page  int                `json:"page"`
	Users []UserRegistration `json:"users"`
}
