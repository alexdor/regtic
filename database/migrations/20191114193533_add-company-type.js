exports.up = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE companies ADD type TEXT;
  `);
};

exports.down = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE companies DROP IF EXISTS type;
  `);
};
