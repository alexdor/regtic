package db

import "github.com/alexdor/regtic/api/models"

var badCompanyOrderBy = models.BadCompanyColumns.Type + " desc"
var badPersonWhereClause = models.BadPersonColumns.NameVector + " @@ plainto_tsquery('simple', ? )"
var badPersonSelect = models.BadPersonColumns.Type + ", " + models.BadPersonColumns.Source + ", ts_rank(" + models.BadPersonColumns.NameVector + ", plainto_tsquery('simple', $1)) as rank"
var badPersonOrderBy = "rank desc, " + models.BadPersonColumns.Type + " desc"

var findCompanyWhereClause = models.CompanyColumns.NameVector + " @@ plainto_tsquery('simple', ?)"
