package repository

import (
	"context"
	"fmt"

	"questionnarie/pkg/model"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Repository struct {
	db *pg.DB
}

func New(conn string) (*Repository, error) {
	opts, err := pg.ParseURL(conn)
	if err != nil {
		return nil, fmt.Errorf("pg.ParseURL err: %w", err)
	}

	db := pg.Connect(opts)
	err = db.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("db.Ping err: %w", err)
	}

	r := &Repository{db: db}

	return r, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) createSchema() error {
	models := []interface{}{
		&model.Question{},
	}

	for _, m := range models {
		opts := &orm.CreateTableOptions{
			IfNotExists: true,
		}

		if err := r.db.Model(m).CreateTable(opts); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) CreateQuestion(ctx context.Context, q *model.Question) error {
	_, err := r.db.ModelContext(ctx, q).Insert()
	return err
}

func (r *Repository) SelectQuestion(ctx context.Context, id int) (*model.Question, error) {
	q := &model.Question{ID: id}
	if err := r.db.ModelContext(ctx, q).WherePK().Select(); err != nil {
		return nil, err
	}
	return q, nil
}

func (r *Repository) UpdateQuestion(ctx context.Context, q *model.Question) error {
	_, err := r.db.ModelContext(ctx, q).WherePK().Update()
	return err
}

func (r *Repository) DeleteQuestion(ctx context.Context, q *model.Question) error {
	_, err := r.db.ModelContext(ctx, q).WherePK().Delete()
	return err
}
