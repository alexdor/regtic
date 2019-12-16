module.exports = `
INSERT INTO
  persons (first_name, last_name, country_code, address)
VALUES
  ($1, $2, $3, $4)
ON CONFLICT
  (first_name, last_name, country_code, address)
DO UPDATE SET
  first_name = $1,
  last_name = $2,
  country_code = $3,
  address = $4
RETURNING
  id
`;
