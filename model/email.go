package model

import "github.com/ungerik/go-mail"

type Email string

func (self *Email) Get() string {
	return string(*self)
}

func (self *Email) Set(value string) (err error) {
	if value != "" {
		if value, err = email.ValidateAddress(value); err != nil {
			return err
		}
	}
	*self = Email(value)
	return nil
}

func (self *Email) IsEmpty() bool {
	return *self == ""
}

func (self *Email) String() string {
	return self.Get()
}

func (self *Email) SetString(str string) (err error) {
	return self.Set(str)
}

func (self *Email) FixValue(metaData *MetaData) {
}

func (self *Email) Validate(metaData *MetaData) error {
	str := self.Get()
	if self.Required(metaData) || str != "" {
		if _, err := email.ValidateAddress(str); err != nil {
			return err
		}
	}
	return nil
}

func (self *Email) Required(metaData *MetaData) bool {
	return metaData.BoolAttrib("required")
}
