package response

import "github.com/gofiber/fiber/v2"

type res int64

const (
	Noop res = iota
	ERequestNotValid
	EBadRequest
	EInternalServerError
	EOk
)

type ResponseFactory interface {
	Create() error
}

func Make(resType res, ctx *fiber.Ctx, message interface{}) ResponseFactory {
	switch resType {
	case Noop:
		return noop{}
	case ERequestNotValid:
		return requestNotValid{ctx, message}
	case EBadRequest:
		return badRequest{ctx, message}
	}

	panic("resType not valid!")
}

type noop struct{}

func (req noop) Create() error {
	return nil
}

type requestNotValid struct {
	ctx             *fiber.Ctx
	validationError interface{}
}

func (req requestNotValid) Create() error {
	return RequestNotValid(req.ctx, req.validationError)
}

type badRequest struct {
	ctx     *fiber.Ctx
	message interface{}
}

func (req badRequest) Create() error {
	return BadRequest(req.ctx, req.message)
}
