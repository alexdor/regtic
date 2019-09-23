exports.up = function(knex) {
  return knex.schema.raw(`
  ALTER TABLE company_to_company drop constraint if exists company_to_company_pkey;
  ALTER TABLE company_to_company RENAME mother_company TO mother_company_id;
  ALTER TABLE company_to_company RENAME daughter_company TO daughter_company_id;
  ALTER TABLE company_to_company add primary key(mother_company_id,daughter_company_id);
  ALTER TABLE company_to_company drop id, drop updated_at, drop created_at;
  ALTER TABLE company_to_person drop constraint if exists company_to_person_pkey;
  ALTER TABLE company_to_person RENAME company TO company_id;
  ALTER TABLE company_to_person RENAME person TO person_id;
  ALTER TABLE company_to_person add primary key(company_id,person_id);
  ALTER TABLE company_to_person drop id, drop updated_at, drop created_at;
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
  ALTER TABLE company_to_company drop constraint if exists company_to_company_pkey;
  ALTER TABLE company_to_person drop constraint if exists company_to_person_pkey;
  `);
};
