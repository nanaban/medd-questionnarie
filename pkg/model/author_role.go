package model

type AuthorRole int

const (
	AuthorRoleUnknown AuthorRole = iota
	AuthorRoleAdmin
	AuthorRoleCustomer
)
