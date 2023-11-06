package transfer

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct{}

func NewResponse() *Response {
	return &Response{}
}

// OkResponse means everything went as expected
func (Response) OkResponse(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(data)
}

// CreatedResponse indicates that requested resource(s) have been successfully created
func (Response) CreatedResponse(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusCreated).JSON(data)
}

// NoContentResponse means the operation was successful, and server has nothing more to say about it
func (Response) NoContentResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusNoContent).JSON(nil)
}

// BadRequestResponse indicates that the request has been deemed unacceptable by server
func (Response) BadRequestResponse(c *fiber.Ctx, err ...error) error {
	msg := http.StatusText(http.StatusBadRequest)
	if len(err) != 0 {
		msg = err[0].Error()
	}

	response := fiber.Map{"message": msg}
	return c.Status(http.StatusBadRequest).JSON(response)
}

// UnauthorizedResponse responds when user is not authorized
func (Response) UnauthorizedResponse(c *fiber.Ctx) error {
	response := fiber.Map{"message": http.StatusText(http.StatusUnauthorized)}
	return c.Status(http.StatusUnauthorized).JSON(response)
}

// ForbiddenResponse indicates that the action is not allowed
func (Response) ForbiddenResponse(c *fiber.Ctx) error {
	response := fiber.Map{"message": http.StatusText(http.StatusForbidden)}
	return c.Status(http.StatusForbidden).JSON(response)
}

// NotFoundResponse will return 404 with, if provided, the error message.
// Otherwise, it'll return the classic 404 error message.
func (Response) NotFoundResponse(c *fiber.Ctx, err ...error) error {
	msg := http.StatusText(http.StatusNotFound)
	if len(err) != 0 {
		msg = err[0].Error()
	}

	response := fiber.Map{"message": msg}
	return c.Status(http.StatusNotFound).JSON(response)
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly
func (Response) InternalServerErrorResponse(c *fiber.Ctx, err ...error) error {
	msg := http.StatusText(http.StatusInternalServerError)
	if len(err) != 0 {
		msg = err[0].Error()
	}

	response := fiber.Map{"message": msg}
	return c.Status(http.StatusInternalServerError).JSON(response)
}
