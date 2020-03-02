package model

// Validation performs checks on a resource before it is written
// to the store
type Validation interface {
	Error(*Attribute) error
}

// ValidationOption appends a validation
type ValidationOption func(*[]Validation)

// ValidateUnique appends a unique validation
func ValidateUnique() ValidationOption {
	return func(vals *[]Validation) {
		*vals = append(*vals, UniqueValidation{})
	}
}

// ValidateURLSafe appends a URLSafeValidation
func ValidateURLSafe() ValidationOption {
	return func(vals *[]Validation) {
		*vals = append(*vals, URLSafeValidation{})
	}
}

// ValidateRequired appends a RequiredValidation
func ValidateRequired() ValidationOption {
	return func(vals *[]Validation) {
		*vals = append(*vals, RequiredValidation{})
	}
}

// ValidateEmail appends a EmailValidation
func ValidateEmail() ValidationOption {
	return func(vals *[]Validation) {
		*vals = append(*vals, EmailValidation{})
	}
}

// ValidateMinimumLength appends a MinimumLengthValidation
func ValidateMinimumLength(min int) ValidationOption {
	return func(vals *[]Validation) {
		*vals = append(*vals, MinimumLengthValidation{min})
	}
}

// ValidateMaximumLength appends a MaximumLengthValidation
func ValidateMaximumLength(max int) ValidationOption {
	return func(vals *[]Validation) {
		*vals = append(*vals, MaximumLengthValidation{max})
	}
}

// UniqueValidation will generate an error if the attrbiute is not unique
type UniqueValidation struct{}

// Error will return an error if the attribute does not satisfy the UniqueValidation
func (UniqueValidation) Error(att *Attribute) error {
	return nil
}

// URLSafeValidation will generate an error if the attribute is not URL safe,
// e.g. it has no special characters
type URLSafeValidation struct{}

// Error will return an error if the attribute does not satisfy the URLSafeValidation
func (URLSafeValidation) Error(att *Attribute) error {
	return nil
}

// RequiredValidation will generate an error if the attribute is not present
type RequiredValidation struct{}

// Error will return an error if the attribute does not satisfy the RequiredValidation
func (RequiredValidation) Error(att *Attribute) error {
	return nil
}

// EmailValidation will generate an error if the attribute is not an email
type EmailValidation struct{}

// Error will return an error if the attribute does not satisfy the EmailValidation
func (EmailValidation) Error(att *Attribute) error {
	return nil
}

// MinimumLengthValidation will generate an error if the attribute does not satisfy the mininumm length requirement
type MinimumLengthValidation struct {
	maximum int
}

// Error will return an error if the attribute does not satisfy the MinimumLengthValidation
func (MinimumLengthValidation) Error(att *Attribute) error {
	return nil
}

// MaximumLengthValidation will generate an error if the attribute does not satisfy the maximum length requirement
type MaximumLengthValidation struct {
	maximum int
}

// Error will return an error if the attribute does not satisfy the MaximumLengthValidation
func (MaximumLengthValidation) Error(att *Attribute) error {
	return nil
}
