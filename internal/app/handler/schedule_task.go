package handler

import (
	"github.com/gofiber/fiber/v2"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/parser"
	"github.com/cholazzzb/amaz_corp_be/pkg/response"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

func (h *ScheduleHandler) PostAddTask(ctx *fiber.Ctx) error {
	req := new(ent.TaskWithDetailCommandRequest)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	startTime, err := parser.ParseTime(req.StartTime)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	formattedReq := ent.TaskWithDetailCommand{
		ScheduleID:  req.ScheduleID,
		StartTime:   *startTime,
		DurationDay: req.DurationDay,
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		AssigneeID:  req.AssigneeID,
		StatusID:    req.StatusID,
	}

	err = h.svc.AddTask(ctx.Context(), formattedReq)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
}

func (h *ScheduleHandler) GetListTaskStatus(ctx *fiber.Ctx) error {
	out, err := h.svc.GetListTaskStatus(ctx.Context())
	if err != nil {
		return response.InternalServerError(ctx)
	}
	return response.Ok(ctx, out)
}

func (h *ScheduleHandler) PutEditTask(ctx *fiber.Ctx) error {
	ok, taskID, resFactory := validator.CheckPathParams(ctx, "taskID")
	if !ok {
		return resFactory.Create()
	}

	req := new(ent.TaskWithDetailCommand)
	ok, resFactory = validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	err := h.svc.EditTask(ctx.Context(), taskID, *req)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
}

// TODO: TEST
func (h *ScheduleHandler) CreateTaskDependency(
	ctx *fiber.Ctx,
) error {
	req := new(ent.TaskDependencyCommand)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	err := h.svc.CreateTaskDependency(ctx.Context(), *req)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
}

// TODO: Test
func (h *ScheduleHandler) EditTaskDependency(
	ctx *fiber.Ctx,
) error {
	req := new(ent.TaskDependencyCommand)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	err := h.svc.EditTaskDependency(ctx.Context(), *req)
	if err != nil {
		response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
}

// TODO: Test
func (h *ScheduleHandler) DeleteTaskDependency(
	ctx *fiber.Ctx,
) error {
	req := new(ent.TaskDependencyCommand)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	err := h.svc.DeleteTaskDependency(ctx.Context(), *req)
	if err != nil {
		response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
}

type GetTasksByRoomIDQueryParams struct {
	StartTime string `query:"start-time"`
	EndTime   string `query:"end-time"`
	Page      int32  `query:"page"`
	PageSize  int32  `query:"page-size"`
}

func (h *ScheduleHandler) GetTasksByRoomID(
	ctx *fiber.Ctx,
) error {
	roomID := ctx.Params("roomID")
	queryParams := new(GetTasksByRoomIDQueryParams)
	ok, resFactory := validator.CheckQueryParams(ctx, queryParams)
	if !ok {
		return resFactory.Create()
	}

	startTime, err := parser.ParseTime(queryParams.StartTime)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}
	endTime, err := parser.ParseTime((queryParams.EndTime))
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	tasks, err := h.svc.GetTasksByRoomID(
		ctx.Context(),
		roomID,
		queryParams.Page,
		queryParams.PageSize,
		startTime,
		endTime,
	)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, tasks)

}
