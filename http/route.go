package http

import (
	"fiber-mongo/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

type IError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var Validator = validator.New()

func ValidateAddUser(c *fiber.Ctx) error {
	var errors []*IError
	body := new(domain.User)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			log.Println(err)
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}

func ValidateUpdateUser(c *fiber.Ctx) error {
	var errors []*IError
	body := new(domain.User)
	c.BodyParser(&body)

	err := Validator.StructPartial(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			log.Println(err)
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}

//type Op func(v *Validate) (s interface{}, fields string) error

func CreateAndUpdateValidation(partial bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var errors []*IError
		body := new(domain.User)
		c.BodyParser(&body)
		if partial {
			err := Validator.StructPartial(body)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					var el IError
					el.Field = err.Field()
					el.Tag = err.Tag()
					el.Value = err.Param()
					log.Println(err)
					errors = append(errors, &el)
				}
				return c.Status(fiber.StatusBadRequest).JSON(errors)
			}
		} else {
			err := Validator.Struct(body)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					var el IError
					el.Field = err.Field()
					el.Tag = err.Tag()
					el.Value = err.Param()
					log.Println(err)
					errors = append(errors, &el)
				}
				return c.Status(fiber.StatusBadRequest).JSON(errors)
			}
		}
		return c.Next()
	}
}

func Routes(r fiber.Router, h *Handler) {
	//r.Route("/users", func(r fiber.Router) {
	r.Post("/users", ValidateAddUser, h.AddUser)
	r.Get("/users", h.ListUsers)
	r.Get("/users/:id", h.GetUser)
	r.Delete("/users/:id", h.DeleteUser)
	r.Put("/users/:id", CreateAndUpdateValidation(true), ValidateUpdateUser, h.Update)
	//r.With(ValidateQueryParam(ListPetsQuery{})).Get("/", h.ListPets)
	//r.With(ValidateBody(domain.Pet{})).Post("/", h.AddPet)
	//r.Route("/{id}", func(r chi.Router) {
	//	r.Use(ValidateURLParam("id"))
	//})
	//})
}
