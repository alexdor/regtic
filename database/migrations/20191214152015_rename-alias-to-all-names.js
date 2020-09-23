exports.up = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE bad_persons_aliases
    RENAME TO bad_persons_all_names;
  `);
};

exports.down = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE IF EXISTS bad_persons_all_names
    RENAME TO bad_persons_aliases;
  `);
};
