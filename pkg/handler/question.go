package handler

import (
	"net/http"

	"questionnarie/pkg/model"

	"github.com/labstack/echo/v4"
)

type GetQuestionRequest struct {
	ID int `json:"id" query:"id"`
}

func (h *Handler) GetQuestion(c echo.Context) error {
	req := &GetQuestionRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	q, err := h.repo.SelectQuestion(h.ctx, req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, q)
}

type CreateUploadLinkQuestionRequest struct {
	QuestionValue string `json:"question_value"`
	QuestionLabel string `json:"question_label"`
	//todo other fields
}

func (r *CreateUploadLinkQuestionRequest) Validate() error {
	// todo validate fields
	return nil
}

func (r *CreateUploadLinkQuestionRequest) toQuestion() *model.Question {
	return &model.Question{
		QuestionValue: r.QuestionValue,
		QuestionLabel: r.QuestionLabel,
	}
}

func (h *Handler) CreateUploadLinkQuestion(c echo.Context) error {
	req := &CreateUploadLinkQuestionRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	q := req.toQuestion()
	if err := h.repo.CreateQuestion(h.ctx, req.toQuestion()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, q)
}
