package entity

import "time"

type RefreshToken struct {
	tableName struct{} `pg:"refresh_token" json:"table_name,omitempty"`

	ID        uint      `pg:"id, pk" json:"id,omitempty"`
	Token     string    `pg:"token" json:"name,omitempty"`
	ExpiredAt time.Time `pg:"expired_at" json:"-"`
}
