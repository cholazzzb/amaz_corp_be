package handler

import (
	"log/slog"

	"github.com/cholazzzb/amaz_corp_be/internal/app/service"
	"github.com/cholazzzb/amaz_corp_be/pkg/logger"
	"github.com/cholazzzb/amaz_corp_be/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	svc    *service.ReportService
	logger *slog.Logger
}

func NewReportHandler(svc *service.ReportService) *ReportHandler {
	sublogger := logger.Get().With(slog.String("domain", "report"), slog.String("layer", "handler"))

	return &ReportHandler{svc: svc, logger: sublogger}
}

func (h *ReportHandler) GetReportBySchedule(ctx *fiber.Ctx) error {
	scheduleID := ctx.Params("scheduleID")
	reports, err := h.svc.GetReportBySchedule(ctx.Context(), scheduleID)
	if err != nil {
		return response.InternalServerError(ctx)
	}
	return response.Ok(ctx, reports)
}
