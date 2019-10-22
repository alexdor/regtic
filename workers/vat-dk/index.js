// imports
const axios = require("axios");
const cvrParser = require("./cvr-parser");
const dbHelper = require("./db-helper");

// const companyArray = {}

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
  size: 200,
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

async function scrollRequest(config = {}) {
  const scrollTimeout = config.scrollTimeout ? config.scrollTimeout : "1m";
  const isFirstRequest = !config.scrollId;
  const request = isFirstRequest
    ? cvrApi.post(
        `/cvr-permanent/virksomhed/_search?scroll=${scrollTimeout}`,
        query
      )
    : cvrApi.post("/_search/scroll", {
        scroll: scrollTimeout,
        scroll_id: config.scrollId
      });
  return request;
}

async function asyncForEach(array, callback) {
  for (let index = 0; index < array.length; index++) {
    await callback(array[index], index, array);
  }
}

async function parseAndSaveResponse(response) {
  const hitList = response.data.hits.hits;
  await asyncForEach(hitList, async hit => {
    try {
      const company = cvrParser.parse(hit);

      const isCompanyParsed = !!company;
      if (!isCompanyParsed) return;

      const companyId = await dbHelper.insertCompany(company);

      for (let i = 0; i < company.persons.length; i++) {
        const person = company.persons[i];
        const personId = await dbHelper.insertPerson(person);
        await dbHelper.insertCompanyToPerson(companyId, personId);
      }

      for (let i = 0; i < company.motherCompanies.length; i++) {
        const motherCompany = company.motherCompanies[i];
        const motherCompanyId = await dbHelper.insertCompany(motherCompany);
        if (!motherCompanyId) {
          const error = { error_msg: "no id was returned for company" };
          console.log({ error, motherCompany });
          return;
        }
        await dbHelper.insertCompanyToCompany(motherCompanyId, companyId);
      }
    } catch (error) {
      console.log(error);
    }
  });
}

module.exports.scrollAndParse = async (event, context) => {
  try {
    eventBody = JSON.parse(event.body);
    const scrollRequestResponse = await scrollRequest({
      scrollId: eventBody.scrollId
    });

    const scrollIdFromResponse = scrollRequestResponse.data._scroll_id;
    const hitListLengthFromResponse = scrollRequestResponse.data.hits.hits
      ? scrollRequestResponse.data.hits.hits.length
      : -1;

    await parseAndSaveResponse(scrollRequestResponse);

    return {
      statusCode: 200,
      body: JSON.stringify({
        scrollId: scrollIdFromResponse,
        hitListLength: hitListLengthFromResponse
      })
    };
  } catch (error) {
    console.log(error);
    return {
      statusCode: 500,
      body: JSON.stringify({ error: "unexpected_failure" })
    };
  }
};
