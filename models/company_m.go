package models

import (
	"fmt"
	"reflect"
	"strings"
)

type Company struct {
	IdCompany string `json:"_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Alias     string `json:"alias" validate:"required"`
	Level     string `json:"iduser" validate:"required"`
	Table     string `json:"table" validate:"required"`
}
type CompanyEdit struct {
	IdCompany string `json:"_id" validate:"required"`
	Rev       string `json:"_rev"`
	Name      string `json:"name" validate:"required"`
	Alias     string `json:"alias" validate:"required"`
	Level     string `json:"iduser" validate:"required"`
	Table     string `json:"table" validate:"required"`
}

func (c *Company) Validate() error {
	var requiredFields []string

	// Using reflection to check struct tags
	rt := reflect.TypeOf(*c)
	rv := reflect.ValueOf(*c)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag.Get("validate")

		if strings.Contains(tag, "required") {
			value := rv.Field(i).Interface()
			if isEmpty(value) {
				requiredFields = append(requiredFields, field.Name)
			}
		}
	}

	if len(requiredFields) > 0 {
		return fmt.Errorf("Required fields are empty: %s", strings.Join(requiredFields, ", "))
	}

	return nil
}
