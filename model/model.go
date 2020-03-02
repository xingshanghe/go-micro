package model

import "github.com/micro/go-micro/v2/store"

// NewModel returns an initialised model
func NewModel(opts ...Option) Model {
	var model Model
	for _, o := range opts {
		o(&model)
	}
	return model
}

// Model is a type of data
type Model struct {
	Name       string
	Store      store.Store
	Attributes []Attribute
}

// Option return a function which when envoked sets an option
type Option func(*Model)

// Name which represents the model
func Name(n string) Option {
	return func(m *Model) {
		m.Name = n
	}
}

// Store which will back the model
func Store(s store.Store) Option {
	return func(m *Model) {
		m.Store = s
	}
}
