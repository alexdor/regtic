exports.up = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE persons ADD full_name text
      generated always as (
        first_name || ' ' || last_name
      ) stored;
  `);
};

exports.down = function (knex) {
  return knex.schema.raw(`ALTER TABLE persons DROP IF EXISTS full_name;`);
};
