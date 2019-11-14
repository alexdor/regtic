exports.up = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE bad_persons ALTER COLUMN type SET NOT NULL;
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE bad_persons ALTER COLUMN type DROP NOT NULL;
  `);
};
