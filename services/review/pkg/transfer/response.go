package transfer

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct{}

func NewResponse() *Response {
	return &Response{}
}

// OKResponse means everything went as expected
func (Response) OKResponse(c *fiber.Ctx, data any) error {
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
func (Response) BadRequestResponse(c *fiber.Ctx, err error) error {
	msg := http.StatusText(http.StatusBadRequest)
	if err != nil {
		msg = err.Error()
	}

	response := fiber.Map{"message": msg}
	return c.Status(http.StatusBadRequest).JSON(response)
}

// UnauthorizedResponse responds when user is not authorized
func (Response) UnauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(http.StatusText(http.StatusUnauthorized))
}

// ForbiddenResponse indicates that the action is not allowed
func (Response) ForbiddenResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusForbidden).JSON(http.StatusText(http.StatusForbidden))
}

// NotFoundResponse will return 404 with, if provided, the error message.
// Otherwise, it'll return the classic 404 error message.
func (Response) NotFoundResponse(c *fiber.Ctx, err error) error {
	msg := http.StatusText(http.StatusNotFound)
	if err != nil {
		msg = err.Error()
	}

	response := fiber.Map{"message": msg}
	return c.Status(http.StatusNotFound).JSON(response)
}

// MethodNotAllowedResponse is returned when the request's method is not acceptable
func (Response) MethodNotAllowedResponse(c *fiber.Ctx) error {
	response := fiber.Map{"message": fmt.Sprintf("the %s method is not supported for this content", c.Request().Header.Method())}

	return c.Status(http.StatusMethodNotAllowed).JSON(response)
}

// DuplicateRequestResponse kicks in when the request was already handled before
func (Response) DuplicateRequestResponse(c *fiber.Ctx) error {
	response := fiber.Map{"message": "the content you're trying to add already exists"}

	return c.Status(http.StatusNotAcceptable).JSON(response)
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly
func (Response) InternalServerErrorResponse(c *fiber.Ctx, err error) error {
	msg := http.StatusText(http.StatusInternalServerError)
	if err != nil {
		msg = err.Error()
	}

	response := fiber.Map{"message": msg}
	return c.Status(http.StatusInternalServerError).JSON(response)
}
