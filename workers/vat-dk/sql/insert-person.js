module.exports = `
INSERT INTO
  persons (first_name, last_name, country_code)
VALUES
  ($1, $2, $3)
ON CONFLICT
  (id)
DO UPDATE SET
  first_name = $1,
  last_name = $2,
  country_code = $3
RETURNING
  id
`;
