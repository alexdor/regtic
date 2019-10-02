// dotenv handler
require('load-environment')

// imports
const Pool = require('pg').Pool

const pool = new Pool({
  connectionString: process.env.REGTIC_DATABASE_URL
})

// functions
async function insertCompany (company) {
  return pool.query(`
    INSERT INTO companies (name, address, vat, starting_date, country_code)
    VALUES ($1, $2, $3, $4, $5)
    ON CONFLICT (vat)
    DO UPDATE SET address = $1, country_code = $5
    RETURNING id`,
    [company.name, company.address, company.vat, company.startingDate, company.countryCode]
  )
}

async function insertPerson (person) {
  return pool.query(`
    INSERT INTO persons (first_name, country_code)
    VALUES ($1, $2)
    RETURNING id`,
    [person.name, person.countryCode]
  )
}

async function insertCompanyToPerson (companyId, personId) {
  return pool.query(`
    INSERT INTO company_to_person (company_id, person_id)
    VALUES ($1, $2)`,
    [companyId, personId]
  )
}

async function insertCompanyToCompany (motherCompanyId, daugtherCompanyId) {
  return pool.query(`
    INSERT INTO company_to_company (mother_company_id, daughter_company_id)
    VALUES ($1, $2)`,
    [motherCompanyId, daugtherCompanyId]
  )
}

module.exports = { insertCompany, insertPerson, insertCompanyToPerson, insertCompanyToCompany }
