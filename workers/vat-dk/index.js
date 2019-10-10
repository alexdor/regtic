// imports
const axios = require("axios");
const cvrParser = require("./cvr-parser");
const dbHelper = require("./db-helper");

// configurations
const cvrApi = axios.create({
  baseURL: process.env.CVR_BASE_URL,
  timeout: 200000,
  headers: {
    "Content-Type": "application/json"
  },
  auth: {
    username: process.env.CVR_USER,
    password: process.env.CVR_PASS
  }
});

const query = {
  size: 100,
  _source: [
    "Vrvirksomhed.cvrNummer",
    "Vrvirksomhed.virksomhedMetadata.nyesteNavn.navn",
    "Vrvirksomhed.virksomhedMetadata.nyesteBeliggenhedsadresse",
    "Vrvirksomhed.virksomhedMetadata.stiftelsesDato",
    "Vrvirksomhed.deltagerRelation.deltager",
    "Vrvirksomhed.deltagerRelation.organisationer"
  ],
  query: {
    bool: {
      must: [{ exists: { field: "Vrvirksomhed.cvrNummer" } }],
      must_not: [
        { exists: { field: "Vrvirksomhed.livsforloeb.periode.gyldigTil" } }
      ]
    }
  }
};

async function scrollRequest(body, config = {}) {
  const scrollTimeout = config.scrollTimeout ? config.scrollTimeout : "1m";
  const endpoint = config.scrollId
    ? "_search/scroll"
    : `cvr-permanent/_search?scroll=${scrollTimeout}`;
  return cvrApi.post(endpoint, body);
}

async function parseAndSaveResponse(response) {
  const hitList = response.data.hits.hits;
  hitList.forEach(async hit => {
    const company = cvrParser.parse(hit);

    const isCompanyParsed = !!company;
    if (!isCompanyParsed) return;

    const insertCompanyResult = await dbHelper.insertCompany(company);
    const companyId = insertCompanyResult.rows[0].id;

    company.persons.forEach(async person => {
      const insertPersonResult = await dbHelper.insertPerson(person);
      const personId = insertPersonResult.rows[0].id;
      await dbHelper.insertCompanyToPerson(companyId, personId);
    });

    company.motherCompanies.forEach(async motherCompany => {
      const motherCompanyInsert = await dbHelper.insertCompany(motherCompany);
      const motherCompanyId = motherCompanyInsert.rows[0].id;
      await dbHelper.insertCompanyToCompany(motherCompanyId, companyId);
    });
  });
}

module.exports.scrollAndParse = async (event, context) => {
  try {
    const scrollRequestResponse = await scrollRequest(query, {
      scrollId: event.body.scrollId
    });
    const scrollIdFromResponse = scrollRequestResponse.data._scroll_id;

    await parseAndSaveResponse(scrollRequestResponse);

    return JSON.stringify({
      statusCode: 200,
      body: { scrollId: scrollIdFromResponse }
    });
  } catch (error) {
    console.log(error);
    return JSON.stringify({
      statusCode: 500,
      body: { error: "unexpected_failure" }
    });
  }
};

/*

module.exports.scrollAndParse = async (event, context) => {
  return { statusCode: 200, body: "Hello world!" };
};

*/
