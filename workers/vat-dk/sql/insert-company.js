module.exports = `
INSERT INTO
  companies(name, address, vat, starting_date, country_code)
VALUES
  ($1, $2, $3, $4, $5)
ON CONFLICT
  (vat)
DO UPDATE SET
  vat = $3
RETURNING
  id
`;
