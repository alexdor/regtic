module.exports = `
INSERT INTO
  company_to_person (company_id, person_id, relations, ownership, voting_rights)
VALUES
  ($1, $2, $3, $4, $5)
ON CONFLICT DO NOTHING
`;
