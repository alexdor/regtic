module.exports = `
INSERT INTO
  persons (first_name, last_name, country_code, zip_code, region, street, city)
VALUES
  ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT ON CONSTRAINT persons_unique_name_address
DO UPDATE SET
  first_name = $1,
  last_name = $2,
  country_code = $3,
  zip_code = $4,
  region = $5,
  street = $6,
  city = $7
RETURNING
  id
`;
