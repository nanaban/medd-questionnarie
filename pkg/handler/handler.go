package handler

import (
	"context"

	"questionnarie/pkg/model"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Close() error
	CreateQuestion(ctx context.Context, q *model.Question) error
	SelectQuestion(ctx context.Context, id int) (*model.Question, error)
	UpdateQuestion(ctx context.Context, q *model.Question) error
	DeleteQuestion(ctx context.Context, q *model.Question) error
}

type Handler struct {
	ctx  context.Context
	repo Repository
}

func NewHandler(ctx context.Context, repo Repository) *Handler {
	h := &Handler{
		ctx:  ctx,
		repo: repo,
	}

	return h
}

func (h *Handler) Register(g *echo.Group) {
	g.GET("/question", h.GetQuestion)
	g.POST("/question/upload-link", h.CreateUploadLinkQuestion)
}
