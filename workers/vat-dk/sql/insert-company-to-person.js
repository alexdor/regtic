module.exports = `
INSERT INTO
  company_to_person (company_id, person_id)
VALUES
  ($1, $2)
ON CONFLICT DO NOTHING
`;
