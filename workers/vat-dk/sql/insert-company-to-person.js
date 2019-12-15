module.exports = `
INSERT INTO
  company_to_person (company_id, person_id, relations)
VALUES
  ($1, $2, $3)
ON CONFLICT DO NOTHING
`;
