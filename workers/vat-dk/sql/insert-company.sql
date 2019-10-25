INSERT INTO
  companies(name, name_vector, address, vat, starting_date, country_code)
VALUES
  ($1, to_tsvector($1) , $2, $3, $4, $5)
ON CONFLICT
  (vat)
DO UPDATE SET
  vat = $3
RETURNING
  id
