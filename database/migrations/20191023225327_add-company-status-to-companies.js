exports.up = function(knex) {
  return knex.schema.raw(`
    CREATE TYPE COMPANY_STATUS AS ENUM ('active', 'liquidation', 'dissolved', 'closed');
    ALTER TABLE companies ADD status COMPANY_STATUS;
    ALTER TABLE companies ADD status_notes TEXT; /* we can use this field to store additional info like 'dissolved after bankruptcy', 'dissolved after merger' or 'forced dissolved' */
  `);
};

exports.down = function(knex) {
  return knex.schema.raw(`
    DROP TYPE IF EXISTS COMPANY_STATUS;
    ALTER TABLE companies DROP COLUMN IF EXISTS status;
     ALTER TABLE companies DROP COLUMN IF EXISTS status_notes;
  `);
};
