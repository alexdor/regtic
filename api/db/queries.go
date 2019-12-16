package db

import (
	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"
)

var (
	badCompanyOrderBy    = models.BadCompanyColumns.Type + " desc"
	badPersonWhereClause = models.BadPersonColumns.NameVector + " @@ plainto_tsquery('simple', ? )"
	badPersonSelect      = models.BadPersonColumns.Type + ", " + models.BadPersonColumns.Source + ", ts_rank(" + models.BadPersonColumns.NameVector + ", plainto_tsquery('simple', $1)) as rank"
	badPersonOrderBy     = "rank desc, " + models.BadPersonColumns.Type + " desc"

	findCompanyWhereClause = models.CompanyColumns.NameVector + " @@ plainto_tsquery('simple', ?)"
	companyToPeopleFilter  = qm.AndIn(
		models.CompanyToPersonColumns.Relations+" && ?",
		types.StringArray{"ultimate beneficial owner", "legal owner", "founder"})

	companyToCompanyFilter = qm.AndIn(
		models.CompanyToCompanyColumns.Relations+" && ?",
		types.StringArray{"ultimate beneficial owner", "legal owner", "founder"})
)
