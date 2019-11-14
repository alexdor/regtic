module.exports = `
INSERT INTO
  companies(name, address, vat, starting_date, country_code, type)
VALUES
  ($1, $2, $3, $4, $5, $6)
ON CONFLICT
  (vat) WHERE type IS NOT NULL
DO UPDATE SET
  type = $6
RETURNING
  id
`;
