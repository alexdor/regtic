INSERT INTO
  persons (first_name, country_code)
VALUES
  ($1, $2)
ON CONFLICT
  (id)
DO UPDATE SET
  first_name = $1
RETURNING
  id
