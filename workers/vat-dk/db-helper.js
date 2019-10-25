// imports
const fs = require("fs");
const path = require("path");
const { Pool } = require("pg");

const pool = new Pool({
  connectionString: process.env.REGTIC_DATABASE_URL
});

// helper function to load external .sql files
function sql(file) {
  const fullPath = path.join(__dirname, "sql", file);
  return fs.readFileSync(fullPath).toString();
}

const insertCompanyQuery = sql("insert-company.sql");
const insertPersonQuery = sql("insert-person.sql");
const insertCompanyToPersonQuery = sql("insert-company-to-person.sql");
const insertCompanyToCompanyQuery = sql("insert-company-to-company.sql");

// functions
async function insertCompany(client, company) {
  return client.query(insertCompanyQuery, [
    company.name,
    company.address,
    company.vat,
    company.startingDate,
    company.countryCode
  ]);
}

async function insertPerson(client, person) {
  return client.query(insertPersonQuery, [
    person.firstName,
    person.lastName,
    person.countryCode
  ]);
}

async function insertCompanyToPerson(client, companyId, personId) {
  return client.query(insertCompanyToPersonQuery, [companyId, personId]);
}

async function insertCompanyToCompany(
  client,
  motherCompanyId,
  daugtherCompanyId
) {
  return client.query(insertCompanyToCompanyQuery, [
    motherCompanyId,
    daugtherCompanyId
  ]);
}

async function insertDataTransactionally(company) {
  const client = await pool.connect();
  try {
    await client.query("BEGIN");

    // insert company
    const companyId = (await insertCompany(client, company)).rows[0].id;

    // insert persons
    const personsInsertPromises = company.persons.map(person => {
      return insertPerson(client, person);
    });
    const personIds = (await Promise.all(personsInsertPromises)).map(
      result => result.rows[0].id
    );

    // insert company to persons relations
    const companyToPersonInsertPromises = personIds.map(personId => {
      return insertCompanyToPerson(client, companyId, personId);
    });
    await Promise.all(companyToPersonInsertPromises);

    // insert mother companies
    const motherCompanyInsertPromises = company.motherCompanies.map(
      motherCompany => {
        return insertCompany(client, motherCompany);
      }
    );
    const motherCompanyIds = (await Promise.all(
      motherCompanyInsertPromises
    )).map(result => result.rows[0].id);

    // insert company to company relations
    const companyToCompanyInsertPromises = motherCompanyIds.map(
      motherCompanyId => {
        return insertCompanyToCompany(client, motherCompanyId, companyId);
      }
    );
    await Promise.all(companyToCompanyInsertPromises);

    await client.query("COMMIT");
  } catch (error) {
    await client.query("ROLLBACK");
    console.log(error);
  } finally {
    client.release();
  }
}

module.exports = {
  insertDataTransactionally
};
