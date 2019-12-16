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
    person.countryCode,
    person.address
  ]);
}

function insertCompanyToPerson(
  client,
  companyId,
  personId,
  relations,
  ownership,
  votingRights
) {
  return client.query(insertCompanyToPersonQuery, [
    companyId,
    personId,
    `{${relations.join(",")}}`,
    ownership,
    votingRights
  ]);
}

function insertCompanyToCompany(
  client,
  motherCompanyId,
  daugtherCompanyId,
  relations,
  ownership,
  votingRights
) {
  return client.query(insertCompanyToCompanyQuery, [
    motherCompanyId,
    daugtherCompanyId,
    `{${relations.join(",")}}`,
    ownership,
    votingRights
  ]);
}

async function insertDataTransactionally(company) {
  const client = await pool.connect();
  try {
    await client.query("BEGIN");

    // insert company
    const companyId = (await insertCompany(client, company)).rows[0].id;

    // make persons promises
    const personsInsertPromises = company.persons.map(person => {
      return insertPerson(client, person).then(result => {
        return {
          id: result.rows[0].id,
          relations: person.relations,
          ownership: person.ownershipPercentage,
          votingRights: person.votingsRightsPercentage
        };
      });
    });

    // make mother companies promises
    const motherCompanyInsertPromises = company.motherCompanies.map(
      motherCompany => {
        return insertCompany(client, motherCompany).then(result => {
          return {
            id: result.rows[0].id,
            relations: motherCompany.relations,
            ownership: motherCompany.ownershipPercentage,
            votingRights: motherCompany.votingsRightsPercentage
          };
        });
      }
    );

    // insert companies and persons
    const [persons, motherCompanies] = await Promise.all([
      Promise.all(personsInsertPromises),
      Promise.all(motherCompanyInsertPromises)
    ]);

    // insert company to persons relations
    const companyToPersonInsertPromises = persons.map(person => {
      return insertCompanyToPerson(
        client,
        companyId,
        person.id,
        person.relations,
        person.ownership,
        person.votingRights
      );
    });

    // insert company to company relations
    const companyToCompanyInsertPromises = motherCompanies.map(
      motherCompany => {
        return insertCompanyToCompany(
          client,
          motherCompany.id,
          companyId,
          motherCompany.relations,
          motherCompany.ownership,
          motherCompany.votingRights
        );
      }
    );

    await Promise.all([
      ...companyToPersonInsertPromises,
      ...companyToCompanyInsertPromises
    ]);

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
