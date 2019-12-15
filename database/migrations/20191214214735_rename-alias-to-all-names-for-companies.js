exports.up = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE bad_companies_aliases
    RENAME TO bad_companies_all_names;
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE IF EXISTS bad_companies_all_names
    RENAME TO bad_companies_aliases;
  `);
};
