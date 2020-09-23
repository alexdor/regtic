exports.up = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE companies DROP IF EXISTS name_vector;
    ALTER TABLE companies ADD name_vector tsvector
      generated always as (
        to_tsvector('simple', name)
      ) stored;
    ALTER TABLE persons ADD name_vector tsvector
      generated always as (
        to_tsvector('simple', first_name || ' ' || last_name)
      ) stored;
    ALTER TABLE bad_persons ADD name_vector tsvector
      generated always as (
        to_tsvector('simple', full_name)
      ) stored;
    ALTER TABLE bad_persons_aliases ADD name_vector tsvector
      generated always as (
        to_tsvector('simple', full_name)
      ) stored;
  `);
};

exports.down = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE companies DROP IF EXISTS name_vector;
    ALTER TABLE persons DROP IF EXISTS name_vector;
    ALTER TABLE bad_persons DROP IF EXISTS name_vector;
    ALTER TABLE bad_persons_aliases DROP IF EXISTS name_vector;
  `);
};
