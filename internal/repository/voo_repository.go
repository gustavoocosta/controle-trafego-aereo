package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/gustavocosta/controle-trafego-aereo/internal/model"
)

type VooRepository struct {
    db *sqlx.DB
}

func NewVooRepository(db *sqlx.DB) *VooRepository {
    return &VooRepository{db: db}
}

func (r *VooRepository) ListarTodos() ([]model.Voo, error) {
    var voos []model.Voo
    err := r.db.Select(&voos, "SELECT * FROM voos")
    return voos, err
}