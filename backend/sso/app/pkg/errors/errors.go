package errors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type CustomError struct {
	desc       string
	statusCode int
	showToUser bool
	errConst   string
}

func NewCustomError(desc string, status int, show bool, errConst string) *CustomError {
	return &CustomError{
		desc:       desc,
		statusCode: status,
		showToUser: show,
		errConst:   errConst,
	}
}

func (e *CustomError) Error() string {
	return e.desc
}

func (e *CustomError) GetStatusCode() int {
	return e.statusCode
}

func (e *CustomError) IsNeedToBeShown() bool {
	return e.showToUser
}

func (e *CustomError) GetErrConst() string {
	return e.errConst
}

func (e *CustomError) FormatParams(params ...any) *CustomError {
	return &CustomError{
		desc:       fmt.Sprintf(e.Error(), params...),
		statusCode: e.GetStatusCode(),
		showToUser: e.IsNeedToBeShown(),
		errConst:   e.GetErrConst(),
	}
}

func (e *CustomError) ToFiberError(ctx *fiber.Ctx) error {
	if e.IsNeedToBeShown() {
		return ctx.Status(e.GetStatusCode()).JSON(fiber.Map{
			"error": e.Error(),
		})
	}
	return ctx.SendStatus(e.GetStatusCode())
}

func ToCustomError(err error, statusCode int, showToUser bool, errConst string) *CustomError {
	return &CustomError{
		desc:       err.Error(),
		statusCode: statusCode,
		showToUser: showToUser,
		errConst:   errConst,
	}
}
