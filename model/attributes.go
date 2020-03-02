package model

// Attribute is a field on a model
type Attribute interface {
	// BeforeCreate()
	// BeforeUpdate()
	// BeforeSave()
}

// UUIDAttribute is an attribute with type UUID V4
type UUIDAttribute struct {
	AttributeBase
}

// StringAttribute is an attribute with type String
type StringAttribute struct {
	AttributeBase
}

// AttributeBase stores the shared attributes for an attribute
type AttributeBase struct {
	Name        string
	PrimaryKey  bool
	Validations []Validation
}

// Init initialises the Attribute
func (a *AttributeBase) Init(opts ...AttributeOption) {
	for _, o := range opts {
		o(a)
	}
}

// AttributeOption returns a function which when envoked sets an option
type AttributeOption func(*AttributeBase)

// WithUUIDAttribute appends a UUID attribute
func WithUUIDAttribute(name string, opts ...AttributeOption) Option {
	return func(m *Model) {
		var a UUIDAttribute
		a.Init(opts...)
		a.Init(AttributeName(name))

		m.Attributes = append(m.Attributes, a)
	}
}

// WithStringAttribute appends a String attribute
func WithStringAttribute(name string, opts ...AttributeOption) Option {
	return func(m *Model) {
		var a StringAttribute
		a.Init(opts...)
		a.Init(AttributeName(name))

		m.Attributes = append(m.Attributes, a)
	}
}

// AttributeName sets the name of the attribute
func AttributeName(n string) AttributeOption {
	return func(a *AttributeBase) {
		a.Name = n
	}
}

// AttributePrimaryKey sets the attribute as the primary key
func AttributePrimaryKey() AttributeOption {
	return func(a *AttributeBase) {
		a.PrimaryKey = true
	}
}

// AttributeValidations sets the validations an attribute has
func AttributeValidations(vals ...ValidationOption) AttributeOption {
	return func(a *AttributeBase) {
		for _, v := range vals {
			v(&a.Validations)
		}
	}
}
