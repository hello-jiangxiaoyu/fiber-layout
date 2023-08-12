package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Get(url string, result any) error {
	return sendFastHttpRequest(fiber.MethodGet, url, nil, result)
}

func Post(url string, data, result any) error {
	return sendFastHttpRequest(fiber.MethodPost, url, data, result)
}

func Put(url string, data, result any) error {
	return sendFastHttpRequest(fiber.MethodPut, url, data, result)
}

func Delete(url string, result any) error {
	return sendFastHttpRequest(fiber.MethodDelete, url, nil, result)
}

func sendFastHttpRequest(method, url string, data, result any) error {
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	if data != nil {
		byteData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		req.SetBody(byteData)
	}
	if err := a.Parse(); err != nil {
		return err
	}

	code, body, errs := a.Bytes()
	if err := mergeError(errs); err != nil {
		return err
	}
	if code > fiber.StatusIMUsed {
		return errors.New(fmt.Sprintf("post api err, status: %d, body: %s", code, string(body)))
	}

	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return err
		}
	}

	return nil
}

func mergeError(errs []error) error {
	msg := ""
	for _, err := range errs {
		if err == nil {
			continue
		}
		if msg == "" {
			msg = err.Error()
		} else {
			msg += ": " + err.Error()
		}
	}

	if msg == "" {
		return nil
	}
	return errors.New(msg)
}
