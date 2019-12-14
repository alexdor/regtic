exports.up = function(knex) {
  return knex.schema.raw(`
    CREATE TABLE bad_person_to_person(
      bad_person_id UUID NOT NULL,
      person_id UUID NOT NULL,
      FOREIGN KEY (person_id) REFERENCES persons (id),
      FOREIGN KEY (bad_person_id) REFERENCES bad_persons (id),
      PRIMARY KEY (bad_person_id,person_id)
    );

    CREATE TABLE bad_company_to_company(
      bad_company_id UUID NOT NULL,
      company_id UUID NOT NULL,
      FOREIGN KEY (company_id) REFERENCES companies (id),
      FOREIGN KEY (bad_company_id) REFERENCES bad_companies (id),
      PRIMARY KEY (bad_company_id,company_id)
    );

    ALTER TABLE company_to_company ADD voting_rights REAL NOT NULL DEFAULT 0.0;
    ALTER TABLE company_to_person ADD voting_rights REAL NOT NULL DEFAULT 0.0;
    ALTER TABLE company_to_company ADD ownership REAL NOT NULL DEFAULT 0.0;
    ALTER TABLE company_to_person ADD ownership REAL NOT NULL DEFAULT 0.0;
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    DROP TABLE IF EXISTS bad_person_to_person;
    DROP TABLE IF EXISTS bad_company_to_company;
    ALTER TABLE company_to_company DROP COLUMN IF EXISTS voting_rights;
    ALTER TABLE company_to_person DROP COLUMN IF EXISTS voting_rights;
    ALTER TABLE company_to_company DROP COLUMN IF EXISTS ownership;
    ALTER TABLE company_to_person DROP COLUMN IF EXISTS ownership;
  `);
};
