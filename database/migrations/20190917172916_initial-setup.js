require("load-environment");

const isDevelopment = process.env.NODE_ENV === "development";

exports.up = function (knex) {
  return knex.schema.raw(`
    -- CREATE DATABASE regtic
    -- OWNER =  postgres
    -- ENCODING = 'UTF8'
    -- LC_COLLATE = 'en_US.UTF-8'
    -- LC_CTYPE = 'en_US.UTF-8'
    -- CONNECTION LIMIT = -1

    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    CREATE TABLE persons (
      id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
      first_name TEXT,
      last_name TEXT,
      country_code varchar(2),
      updated_at timestamp NOT NULL,
      created_at timestamp NOT NULL
    );

    CREATE TABLE companies (
      id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
      address TEXT,
      vat TEXT UNIQUE NOT NULL,
      starting_date TEXT,
      country_code varchar(2),
      updated_at timestamp DEFAULT now() NOT NULL,
      created_at timestamp DEFAULT now() NOT NULL
    );

    CREATE TABLE company_to_company (
      id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
      mother_company UUID NOT NULL,
      daughter_company UUID NOT NULL,
      updated_at timestamp DEFAULT now() NOT NULL,
      created_at timestamp DEFAULT now() NOT NULL
    );

    CREATE TABLE company_to_person (
      id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
      company UUID NOT NULL,
      person UUID NOT NULL,
      updated_at timestamp DEFAULT now() NOT NULL,
      created_at timestamp DEFAULT now() NOT NULL
    );

    CREATE TABLE bad_persons (
      id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
      full_name TEXT,
      type TEXT,
      source TEXT,
      address TEXT,
      updated_at timestamp DEFAULT now() NOT NULL,
      created_at timestamp DEFAULT now() NOT NULL
    );

    CREATE TABLE bad_persons_aliases (
      id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
      full_name TEXT,
      bad_person_id UUID NOT NULL,
      updated_at timestamp DEFAULT now() NOT NULL,
      created_at timestamp DEFAULT now() NOT NULL
    );

    ALTER TABLE company_to_company ADD FOREIGN KEY (mother_company) REFERENCES companies (id);

    ALTER TABLE company_to_company ADD FOREIGN KEY (daughter_company) REFERENCES companies (id);

    ALTER TABLE company_to_person ADD FOREIGN KEY (company) REFERENCES companies (id);

    ALTER TABLE company_to_person ADD FOREIGN KEY (person) REFERENCES persons (id);

    ALTER TABLE bad_persons_aliases ADD FOREIGN KEY (bad_person_id) REFERENCES bad_persons (id);
  `);
};

exports.down = function (knex) {
  // make sure that this rollback is not able to run in production
  if (!isDevelopment) return;

  return knex.schema.raw(`
    DROP TABLE bad_persons_aliases;
    DROP TABLE bad_persons;
    DROP TABLE company_to_person;
    DROP TABLE company_to_company;
    DROP TABLE companies;
    DROP TABLE persons;

    DROP EXTENSION "uuid-ossp";
  `);
};
