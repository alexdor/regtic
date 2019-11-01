package db

import (
	"context"
	"strings"

	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func FindCompany(ctx context.Context, name string) ([]*models.Company, error) {
	var nameQuery strings.Builder
	for i, value := range strings.Split(name, " ") {
		if i != 0 {
			nameQuery.WriteString(" & ")
		}
		nameQuery.WriteString(value)
	}
	query := qm.Where("name_vector @@ to_tsquery(?)", nameQuery.String())
	return models.Companies(query).All(ctx, DB)
}
