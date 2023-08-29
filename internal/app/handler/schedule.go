package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	ent "github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
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

	scheduleID, err := h.svc.CreateSchedule(ctx.Context(), req.RoomID)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data": map[string]string{
			"scheduleID": scheduleID,
		},
	})
}

func (h *ScheduleHandler) GetScheduleIDByRoomID(ctx *fiber.Ctx) error {
	roomID := ctx.Params("roomID")
	if len(roomID) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "roomID is missing from the request",
		})
	}

	sid, err := h.svc.GetScheduleIDByRoomID(ctx.Context(), roomID)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "ok",
		"schedule_id": sid.ID,
	})
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
		"message":      "ok",
		"task_details": td,
	})
}

func (h *ScheduleHandler) GetListTaskByScheduleID(ctx *fiber.Ctx) error {
	scheduleID := ctx.Params("scheduleID")
	if len(scheduleID) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "scheduleID is missing from the request",
		})
	}
	tks, err := h.svc.GetListTaskByScheduleID(ctx.Context(), scheduleID)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"tasks":   tks,
	})
}

func (h *ScheduleHandler) PostAddTask(ctx *fiber.Ctx) error {
	req := new(ent.TaskWithDetailCommand)
	if err := ctx.BodyParser(req); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	err := h.svc.AddTask(ctx.Context(), *req)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

func (h *ScheduleHandler) PutEditTask(ctx *fiber.Ctx) error {
	taskID := ctx.Params("taskID")
	req := new(ent.TaskWithDetailCommand)

	if err := ctx.BodyParser(req); err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errs := validator.Validate(req); errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errs)
	}

	err := h.svc.EditTask(ctx.Context(), taskID, *req)
	if err != nil {
		h.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}
