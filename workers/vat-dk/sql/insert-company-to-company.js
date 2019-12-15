module.exports = `
INSERT INTO
  company_to_company (mother_company_id, daughter_company_id, relations, ownership, voting_rights)
VALUES
  ($1, $2, $3, $4, $5)
ON CONFLICT DO NOTHING
`;
