package internal

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var validate = validator.New()

func checkJson(obj any) error {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		return errors.New("bind object should be a pointer")
	}

	t = t.Elem()
	fmt.Println(t.Kind())
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Struct && t.Kind() != reflect.Map && t.Kind() != reflect.Interface {
		return errors.New("can not bind json to " + t.Kind().String() + " type")
	}
	return nil
}

func (a *Api) BindJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := checkJson(obj); err != nil {
		return a.setError(err)
	}
	if err := a.c.BodyParser(obj); err != nil {
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

	if err := checkJson(obj); err != nil {
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

func (a *Api) BindUriAndJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := checkJson(obj); err != nil {
		return a.setError(err)
	}
	if err := a.c.ParamsParser(obj); err != nil {
		return a.setError(err)
	}
	if err := a.c.BodyParser(obj); err != nil {
		return a.setError(err)
	}

	if err := validate.Struct(obj); err != nil {
		return a.setError(err)
	}

	return a
}
