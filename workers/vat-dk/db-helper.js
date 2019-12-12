// imports
const { Pool } = require("pg");

const insertCompanyQuery = require("./sql/insert-company");
const insertPersonQuery = require("./sql/insert-person");
const insertCompanyToPersonQuery = require("./sql/insert-company-to-person");
const insertCompanyToCompanyQuery = require("./sql/insert-company-to-company");

const pool = new Pool({
  connectionString: process.env.REGTIC_DATABASE_URL
});

// functions
function insertCompany(client, company) {
  return client.query(insertCompanyQuery, [
    company.name,
    company.address,
    company.vat,
    company.startingDate,
    company.countryCode,
    company.type
  ]);
}

function insertPerson(client, person) {
  return client.query(insertPersonQuery, [
    person.firstName,
    person.lastName,
    person.countryCode
  ]);
}

function insertCompanyToPerson(client, companyId, personId, relations) {
  return client.query(insertCompanyToPersonQuery, [
    companyId,
    personId,
    `{${relations.join(",")}}`
  ]);
}

function insertCompanyToCompany(
  client,
  motherCompanyId,
  daugtherCompanyId,
  relations
) {
  return client.query(insertCompanyToCompanyQuery, [
    motherCompanyId,
    daugtherCompanyId,
    `{${relations.join(",")}}`
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
      return insertPerson(client, person).then(result => {
        return {
          id: result.rows[0].id,
          relations: person.relations
        };
      });
    });
    const persons = await Promise.all(personsInsertPromises);

    // insert company to persons relations
    const companyToPersonInsertPromises = persons.map(person => {
      return insertCompanyToPerson(
        client,
        companyId,
        person.id,
        person.relations
      );
    });
    await Promise.all(companyToPersonInsertPromises);

    // insert mother companies
    const motherCompanyInsertPromises = company.motherCompanies.map(
      motherCompany => {
        return insertCompany(client, motherCompany).then(result => {
          return {
            id: result.rows[0].id,
            relations: motherCompany.relations
          };
        });
      }
    );
    const motherCompanies = await Promise.all(motherCompanyInsertPromises);

    // insert company to company relations
    const companyToCompanyInsertPromises = motherCompanies.map(
      motherCompany => {
        return insertCompanyToCompany(
          client,
          motherCompany.id,
          companyId,
          motherCompany.relations
        );
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
