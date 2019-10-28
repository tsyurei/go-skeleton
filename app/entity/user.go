package entity

type User struct {
	tableName struct{} `pg:"user"`

	Id             uint   `pg:"id,pk"`
	Email          string `pg:email`
	HashedPassword []byte `pg:hashed_password`
}
