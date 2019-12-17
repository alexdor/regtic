exports.up = function(knex) {
  return knex.schema.raw(
    `
    CREATE
    OR REPLACE FUNCTION immutable_concat_ws(text, VARIADIC text []) RETURNS TEXT AS 'text_concat_ws' LANGUAGE internal immutable;
    -- Add person fields
    ALTER TABLE persons
    ADD
      street TEXT,
    ADD
      region TEXT,
    ADD
      zip_code TEXT,
    ADD
      city TEXT,
    ADD
      address TEXT generated always as (
        immutable_concat_ws(', ', street, region, zip_code, city)
      ) stored,
    ADD
      CONSTRAINT persons_unique_name_address UNIQUE (
        first_name,
        last_name,
        country_code,
        street,
        region,
        zip_code,
        city
      );
    -- Add company fields
    ALTER TABLE companies
    ADD
      street TEXT,
    ADD
      region TEXT,
    ADD
      zip_code TEXT,
    ADD
      city TEXT,
      DROP COLUMN address,
    ADD
      COLUMN address TEXT generated always as (
        immutable_concat_ws(', ', street, region, zip_code, city)
      ) stored;
  `
  );
};
exports.down = function(knex) {
  return knex.schema.raw(
    `
ALTER TABLE persons DROP COLUMN IF EXISTS street;
ALTER TABLE persons DROP COLUMN IF EXISTS region;
ALTER TABLE persons DROP COLUMN IF EXISTS zip_code;
ALTER TABLE persons DROP COLUMN IF EXISTS city;
ALTER TABLE persons DROP COLUMN IF EXISTS address;
ALTER TABLE persons DROP CONSTRAINT IF EXISTS unique_name_address_country;
ALTER TABLE companies DROP COLUMN IF EXISTS street;
ALTER TABLE companies DROP COLUMN IF EXISTS region;
ALTER TABLE companies DROP COLUMN IF EXISTS zip_code;
ALTER TABLE companies DROP COLUMN IF EXISTS city;
ALTER TABLE companies DROP COLUMN address,
ADD
  COLUMN address TEXT;
`
  );
};
