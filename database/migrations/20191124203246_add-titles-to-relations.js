exports.up = function(knex) {
  return knex.schema.raw(`
    CREATE TYPE ENTITY_RELATION AS ENUM (
      'ultimate beneficial owner',
      'legal owner',
      'accountant',
      'direction',
      'board of directors',
      'founder'
    );
    ALTER TABLE company_to_company ADD relations ENTITY_RELATION[] NOT NULL;
    ALTER TABLE company_to_person ADD relations ENTITY_RELATION[] NOT NULL;
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    ALTER TABLE company_to_company DROP COLUMN IF EXISTS relations;
    ALTER TABLE company_to_person DROP COLUMN IF EXISTS relations;
    DROP TYPE IF EXISTS ENTITY_RELATION;
  `);
};
