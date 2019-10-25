INSERT INTO
  company_to_company (mother_company_id, daughter_company_id)
VALUES
  ($1, $2)
ON CONFLICT DO NOTHING
