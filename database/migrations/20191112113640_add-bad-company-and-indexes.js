const defaultFieldsForTable = `id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  updated_at timestamp DEFAULT now() NOT NULL,
  created_at timestamp DEFAULT now() NOT NULL`;

const uniqueRel = (tableName, columns) =>
  `CREATE UNIQUE INDEX ${tableName}_${columns
    .split(", ")
    .join("_")}_uniqueness ON ${tableName} (${columns});`;

const countryCodeType = `varchar(2) default 'ZZ' NOT NULL`;
const addressInfo = `country_code ${countryCodeType},
  street TEXT,
  city TEXT,
  zip_code TEXT,
  region TEXT,
  place TEXT,
  po_box  TEXT`;

const foreignKeyToID = (keyName, table) =>
  `${keyName} UUID,
 FOREIGN KEY (${keyName}) REFERENCES ${table} (id)`;

const citizenShip = `citizenship_country_code varchar(2)[] default '{ZZ}' NOT NULL`;

exports.up = function(knex) {
  return knex.schema.raw(`
    CREATE TYPE ADDRESS_TYPE AS ENUM ('birthplace', 'address');

    CREATE TABLE bad_companies (
      ${defaultFieldsForTable},
      name TEXT NOT NULL,
      name_vector tsvector
      generated always as (
        to_tsvector('simple', name)
      ) stored,
      address TEXT,
      source TEXT NOT NULL,
      citizenship_region TEXT,
      type BAD_PERSON_TYPE NOT NULL,
      ${citizenShip}
    );

    CREATE TABLE bad_companies_aliases (
      ${defaultFieldsForTable},
      ${foreignKeyToID("bad_company_id", "bad_companies")},
      name TEXT NOT NULL,
      name_vector tsvector
      generated always as (
        to_tsvector('simple', name)
        ) stored,
      type BAD_PERSON_TYPE NOT NULL
    );

    CREATE TABLE bad_companies_addresses (
      ${defaultFieldsForTable},
      ${foreignKeyToID("bad_company_id", "bad_companies")},
      ${addressInfo}
    );

    CREATE TABLE bad_persons_addresses (
      ${defaultFieldsForTable},
      ${foreignKeyToID("bad_person_id", "bad_persons")},
      ${addressInfo},
      type ADDRESS_TYPE NOT NULL
    );

    ALTER TABLE companies ALTER COLUMN country_code SET DEFAULT 'ZZ';ALTER TABLE companies ALTER COLUMN country_code SET NOT NULL;

    ALTER TABLE persons ALTER COLUMN country_code SET DEFAULT 'ZZ';
    ALTER TABLE persons ALTER COLUMN country_code SET NOT NULL;

    ALTER TABLE bad_persons DROP IF EXISTS country_code;
    ALTER TABLE bad_persons ADD ${citizenShip};

    ${uniqueRel(
      "bad_persons",
      "name_vector, address, citizenship_country_code, type"
    )}
    ${uniqueRel(
      "bad_persons_addresses",
      "bad_person_id, country_code, street, city, zip_code, region, place, po_box"
    )}
    ${uniqueRel("bad_persons_aliases", "bad_person_id, full_name")}

    ${uniqueRel("bad_companies", "name_vector, address, type")}
    ${uniqueRel("bad_companies_aliases", "bad_company_id, name")}
    ${uniqueRel(
      "bad_companies_addresses",
      "bad_company_id, country_code, street, city, zip_code, region, place, po_box"
    )}
     `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    DROP INDEX IF EXISTS bad_companies_aliases_uniqueness;
    DROP INDEX IF EXISTS bad_persons_aliases_uniqueness;
    DROP INDEX IF EXISTS bad_persons_uniqueness;
    DROP INDEX IF EXISTS bad_companies_uniqueness;
    ALTER TABLE bad_persons DROP IF EXISTS citizenship_country_code;
    DROP TABLE bad_companies_aliases;
    DROP TABLE bad_companies_addresses;
    DROP TABLE bad_persons_addresses;
    DROP TABLE bad_companies;
    DROP INDEX IF EXISTS bad_persons_aliases_bad_person_id_full_name_uniqueness;
    DROP TYPE IF EXISTS ADDRESS_TYPE;
  `);
};
