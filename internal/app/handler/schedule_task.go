package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"

	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/response"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

func (h *ScheduleHandler) PostAddTask(ctx *fiber.Ctx) error {
	req := new(ent.TaskWithDetailCommandRequest)
	ok, resFactory := validator.CheckReqBodySchema(ctx, req)
	if !ok {
		return resFactory.Create()
	}

	var startTime *time.Time
	startTimeParsed, err := time.Parse(time.RFC1123, req.StartTime)
	if err == nil {
		startTime = &startTimeParsed
	}

	formattedReq := ent.TaskWithDetailCommand{
		ScheduleID:  req.ScheduleID,
		StartTime:   *startTime,
		DurationDay: req.DurationDay,
		Name:        req.Name,
		OwnerID:     req.OwnerID,
		AssigneeID:  req.AssigneeID,
		Status:      req.Status,
	}

	err = h.svc.AddTask(ctx.Context(), formattedReq)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, nil)
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
