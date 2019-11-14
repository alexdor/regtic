exports.up = function(knex) {
  return knex.schema.raw(`
    CREATE TABLE bad_companies (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name TEXT NOT NULL,
    name_vector tsvector
    generated always as (
      to_tsvector('simple', name)
    ) stored,
    address TEXT,
    source TEXT NOT NULL,
    country_code varchar(2) default 'ZZ' NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    type BAD_PERSON_TYPE NOT NULL
    );

    CREATE TABLE bad_companies_aliases (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    company_id UUID,
    name TEXT NOT NULL,
    name_vector tsvector
    generated always as (
      to_tsvector('simple', name)
    ) stored,
    address TEXT,
    source TEXT NOT NULL,
    country_code varchar(2) default 'ZZ' NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    type BAD_PERSON_TYPE NOT NULL
    );

    ALTER TABLE companies ALTER COLUMN country_code SET DEFAULT 'ZZ';ALTER TABLE companies ALTER COLUMN country_code SET NOT NULL;

    ALTER TABLE persons ALTER COLUMN country_code SET DEFAULT 'ZZ';
    ALTER TABLE persons ALTER COLUMN country_code SET NOT NULL;

    ALTER TABLE bad_companies_aliases ADD FOREIGN KEY (company_id) REFERENCES companies (id);
    ALTER TABLE bad_persons ADD country_code varchar(2) DEFAULT 'ZZ' NOT NULL;

    CREATE UNIQUE INDEX bad_persons_uniqueness ON bad_persons (name_vector, address, country_code, type);
    CREATE UNIQUE INDEX bad_companies_uniqueness ON bad_companies (name_vector, address, country_code, type);
    CREATE UNIQUE INDEX bad_companies_aliases_uniqueness ON bad_companies_aliases (company_id, name);
    CREATE UNIQUE INDEX bad_persons_aliases_uniqueness ON bad_persons_aliases (bad_person_id, full_name);
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    DROP INDEX IF EXISTS bad_companies_aliases_uniqueness;
    DROP INDEX IF EXISTS bad_persons_aliases_uniqueness;
    DROP INDEX IF EXISTS bad_persons_uniqueness;
    DROP INDEX IF EXISTS bad_companies_uniqueness;
    ALTER TABLE bad_persons DROP country_code;
    DROP TABLE bad_companies;
    DROP TABLE bad_companies_aliases;
  `);
};
