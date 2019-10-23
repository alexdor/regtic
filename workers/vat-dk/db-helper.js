// imports
const { Pool } = require("pg");

const pool = new Pool({
  connectionString: process.env.REGTIC_DATABASE_URL
});

// functions
async function insertCompany(company) {
  const client = await pool.connect();
  try {
    const result = await client.query(
      `
      INSERT INTO companies(name, name_vector, address, vat, starting_date, country_code)
      VALUES ($1, to_tsvector($1) , $2, $3, $4, $5)
      ON CONFLICT (vat) DO UPDATE SET vat = $3
      RETURNING id`,
      [
        company.name,
        company.address,
        company.vat,
        company.startingDate,
        company.countryCode
      ]
    );

    return result.rows[0].id;
  } catch (error) {
    console.log(error, company);
  } finally {
    client.release();
  }
}

async function insertPerson(person) {
  const client = await pool.connect();
  try {
    const result = await client.query(
      `
      INSERT INTO persons (first_name, country_code)
      VALUES ($1, $2)
      ON CONFLICT (id) DO UPDATE SET first_name = $1
      RETURNING id`,
      [person.name, person.countryCode]
    );

    return result.rows[0].id;
  } catch (error) {
    console.log(error, person);
  } finally {
    client.release();
  }
}

async function insertCompanyToPerson(companyId, personId) {
  const client = await pool.connect();
  try {
    await client.query(
      `
      INSERT INTO company_to_person (company_id, person_id)
      VALUES ($1, $2)
      ON CONFLICT DO NOTHING`,
      [companyId, personId]
    );
  } catch (error) {
    console.log("companyToPerson error", companyId, personId);
  } finally {
    client.release();
  }
}

async function insertCompanyToCompany(motherCompanyId, daugtherCompanyId) {
  const client = await pool.connect();
  try {
    await client.query(
      `
      INSERT INTO company_to_company (mother_company_id, daughter_company_id)
      VALUES ($1, $2)
      ON CONFLICT DO NOTHING`,
      [motherCompanyId, daugtherCompanyId]
    );
  } catch (error) {
    console.log("companyToCompany error", motherCompanyId, daugtherCompanyId);
  } finally {
    client.release();
  }
}

module.exports = {
  insertCompany,
  insertPerson,
  insertCompanyToPerson,
  insertCompanyToCompany
};
