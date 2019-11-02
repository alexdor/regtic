exports.up = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE companies ADD name_vector TSVECTOR;
    UPDATE companies SET name_vector = to_tsvector(name);
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE companies DROP IF EXISTS name_vector;
  `);
};
