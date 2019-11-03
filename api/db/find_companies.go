package db

import (
	"context"

	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func FindCompany(ctx context.Context, name string) ([]*models.Company, error) {
	return models.Companies(
		qm.Where(
			"name_vector @@ plainto_tsquery('simple', ?)", name,
		),
	).All(ctx, DB)
}
