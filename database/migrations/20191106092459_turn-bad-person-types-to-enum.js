exports.up = function (knex) {
  return knex.schema.raw(`
    CREATE TYPE BAD_PERSON_TYPE AS ENUM ('PEP', 'SANCTION');
    ALTER TABLE bad_persons ALTER COLUMN type TYPE BAD_PERSON_TYPE using type::BAD_PERSON_TYPE;
  `);
};

exports.down = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE bad_persons ALTER COLUMN type TYPE TEXT;
    DROP TYPE IF EXISTS BAD_PERSON_TYPE;
  `);
};
