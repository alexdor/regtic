exports.up = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE companies ADD name text;
    ALTER TABLE persons ALTER COLUMN created_at SET DEFAULT now();
    ALTER TABLE persons ALTER COLUMN updated_at SET DEFAULT now();
  `);
};

exports.down = function(knex) {};
