package db

import (
	"context"

	"github.com/alexdor/regtic/api/interfaces"

	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func FindCompany(ctx context.Context, name string) (res []interfaces.CompanyJson, err error) {
	err = models.Companies(
		qm.Select(
			models.CompanyColumns.Vat,
			models.CompanyColumns.ID,
			models.CompanyColumns.Address,
			models.CompanyColumns.CountryCode,
			models.CompanyColumns.Name,
			models.CompanyColumns.Status,
			models.CompanyColumns.StatusNotes,
			models.CompanyColumns.StartingDate,
			models.CompanyColumns.CreatedAt,
			models.CompanyColumns.UpdatedAt,
			"ts_rank(name_vector, plainto_tsquery('simple', $1)) as rank",
		),
		qm.Where(
			"name_vector @@ plainto_tsquery('simple', ?)", name,
		),
		qm.OrderBy("rank desc"),
		qm.Limit(10),
	).Bind(ctx, DB, &res)
	return
}
