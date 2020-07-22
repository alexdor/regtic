exports.up = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE bad_companies_all_names DROP COLUMN type;
  `);
};

exports.down = function (knex) {
  return knex.schema.raw(`
    ALTER TABLE bad_companies_all_names ADD type BAD_PERSON_TYPE NOT NULL;
  `);
};
