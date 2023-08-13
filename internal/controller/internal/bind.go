package internal

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (a *Api) BindJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := json.Unmarshal(a.c.Body(), obj); err != nil {
		return a.setError(err)
	}
	if err := validate.Struct(obj); err != nil {
		return a.setError(err)
	}

	return a
}

func (a *Api) BindUri(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := a.c.ParamsParser(obj); err != nil {
		return a.setError(err)
	}
	if err := validate.Struct(obj); err != nil {
		return a.setError(err)
	}

	return a
}

func (a *Api) BindUriAndJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := json.Unmarshal(a.c.Body(), obj); err != nil {
		return a.setError(err)
	}
	if err := a.c.ParamsParser(obj); err != nil {
		return a.setError(err)
	}

	if err := validate.Struct(obj); err != nil {
		return a.setError(err)
	}

	return a
}
