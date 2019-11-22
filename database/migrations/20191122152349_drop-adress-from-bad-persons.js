exports.up = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE bad_persons DROP if exists address;
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
      ALTER TABLE bad_persons ADD address TEXT;
  `);
};
