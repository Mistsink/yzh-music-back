package model

type Formable interface {
	Format() (any, error)
}
