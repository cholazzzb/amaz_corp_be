package handler

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/parser"
	"github.com/cholazzzb/amaz_corp_be/pkg/response"
	"github.com/cholazzzb/amaz_corp_be/pkg/validator"
)

type ScheduleHandler struct {
	svc    *service.ScheduleService
	logger *slog.Logger
}

func NewScheduleHandler(svc *service.ScheduleService) *ScheduleHandler {
	sublogger := logger.Get().With(slog.String("domain", "schedule"), slog.String("layer", "handler"))

	return &ScheduleHandler{svc: svc, logger: sublogger}
}

func (h *ScheduleHandler) CreateSchedule(ctx *fiber.Ctx) error {
	req := new(ent.ScheduleCommand)

	if err := ctx.BodyParser(req); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	scheduleID, err := h.svc.CreateSchedule(ctx.Context(), *req)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, ent.ScheduleCommandRes{
		ScheduleID: scheduleID,
	})
}

func (h *ScheduleHandler) GetListScheduleByRoomID(ctx *fiber.Ctx) error {
	roomID := ctx.Params("roomID")
	if len(roomID) == 0 {
		return response.BadRequest(ctx, "roomID is missing from the request")
	}

	schs, err := h.svc.GetListScheduleByRoomID(ctx.Context(), roomID)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, schs)
}

func (h *ScheduleHandler) GetTaskDetail(ctx *fiber.Ctx) error {
	taskID := ctx.Params("taskID")
	if len(taskID) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "taskID is missing from the request",
		})
	}

	td, err := h.svc.GetTaskDetail(ctx.Context(), taskID)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "ok",
		"taskDetail": td,
	})
}

// TODO:
// ## Filter
// assignee=string
// dependency=Array<taskID>
// ## Sort
// sort-by=Array<assignee|owner|startDate|endDate|duration>
// sort-dir=asc|dsc
func (h *ScheduleHandler) GetListTaskByScheduleID(ctx *fiber.Ctx) error {
	scheduleID := ctx.Params("scheduleID")
	if len(scheduleID) == 0 {
		return response.BadRequest(ctx, "scheduleID is missing from the request")
	}

	queryFilterParams := new(ent.TaskQueryFilterParams)

	if err := ctx.QueryParser(queryFilterParams); err != nil {
		return response.BadRequest(ctx, "query params is in wrong format")
	}

	startTime, err := parser.ParseTime(queryFilterParams.StartTime)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	endTime, err := parser.ParseTime(queryFilterParams.EndTime)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	tks, err := h.svc.GetListTaskByScheduleID(ctx.Context(), scheduleID, ent.TaskQueryFilter{
		StartTime: startTime,
		EndTime:   endTime,
	})
	if err != nil {
		return response.InternalServerError(ctx)
	}
	return response.Ok(ctx, tks)
}

func (h *ScheduleHandler) GetListTaskWithDetailByScheduleID(ctx *fiber.Ctx) error {
	ok, scheduleID, resFactory := validator.CheckPathParams(ctx, "scheduleID")
	if !ok {
		return resFactory.Create()
	}

	queryParams := new(ent.AutoSchedulePreviewQueryParams)
	ok, resFactory = validator.CheckQueryParams(ctx, queryParams)
	if !ok {
		return resFactory.Create()
	}

	twds, err := h.svc.GetListTaskWithDetailByScheduleID(ctx.Context(), scheduleID)

	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, twds)
}

func (h *ScheduleHandler) AutoSchedulePreview(ctx *fiber.Ctx) error {
	ok, scheduleID, resFactory := validator.CheckPathParams(ctx, "scheduleID")
	if !ok {
		return resFactory.Create()
	}

	queryParams := new(ent.AutoSchedulePreviewQueryParams)
	ok, resFactory = validator.CheckQueryParams(ctx, queryParams)
	if !ok {
		return resFactory.Create()
	}

	sorted, err := h.svc.AutoSchedulePreview(ctx.Context(), scheduleID)
	if err != nil {
		return response.InternalServerError(ctx)
	}

	return response.Ok(ctx, sorted)
}

func (h *ScheduleHandler) AutoScheduleSave(ctx *fiber.Ctx) error {
	return errors.New("not implemented")
}
