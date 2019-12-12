module.exports = `
INSERT INTO
  company_to_company (mother_company_id, daughter_company_id, relations)
VALUES
  ($1, $2, $3)
ON CONFLICT DO NOTHING
`;
