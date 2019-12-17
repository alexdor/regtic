module.exports = `
INSERT INTO
  companies(name, vat, starting_date, country_code, type, zip_code, region, street, city)
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT
  (vat)
DO UPDATE SET
  name = $1,
  type = $5,
  zip_code = $6,
  region = $7,
  street = $8,
  city = $9
RETURNING
  id
`;
