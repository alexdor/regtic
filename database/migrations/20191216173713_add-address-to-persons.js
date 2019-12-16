exports.up = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE persons ADD address TEXT;
    ALTER TABLE persons ADD CONSTRAINT unique_name_address_country UNIQUE (first_name, last_name, country_code, address);
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE persons DROP COLUMN IF EXISTS address;
    ALTER TABLE persons DROP CONSTRAINT IF EXISTS unique_name_address_country;
  `);
};
