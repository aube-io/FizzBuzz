package api

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

var (
	methods = []string{"Get", "Post", "Put", "Patch", "Delete"}
)

type API struct {
	Namespace string
	Version   string
	Get       map[string][]fiber.Handler
	Post      map[string][]fiber.Handler
	Put       map[string][]fiber.Handler
	Patch     map[string][]fiber.Handler
	Delete    map[string][]fiber.Handler
}

func (api *API) Register(router fiber.Router) {
	apiValue := reflect.ValueOf(*api)
	for _, method := range methods {
		methodValue := apiValue.FieldByName(method)
		for _, key := range methodValue.MapKeys() {
			route := key.String()
			handlersValue := methodValue.MapIndex(key)
			handlers := handlersValue.Interface().([]fiber.Handler)
			router.Add(method, route, handlers...)
		}
	}
}
