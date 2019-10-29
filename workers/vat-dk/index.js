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
  }
});

cvrApi.interceptors.request.use(function(config) {
  config.headers.Authorization = process.env.CVR_AUTH_TOKEN;

  return config;
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
  const scrollTimeout = config.scrollTimeout ? config.scrollTimeout : "2m";
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

async function parseAndSaveResponse(responseData) {
  const hitList = responseData.hits.hits;
  await asyncForEach(hitList, async hit => {
    try {
      const company = cvrParser.parse(hit);
      if (!company) return;
      await dbHelper.insertDataTransactionally(company);
    } catch (error) {
      // TODO: Maybe use sentry here...
      console.log(error);
    }
  });
}

module.exports.scrollAndParse = async (event, context) => {
  try {
    const scrollRequestResponse = await scrollRequest({
      scrollId: JSON.parse(event.body).scrollId
    });

    const data = scrollRequestResponse.data || {};
    await parseAndSaveResponse(data);

    const hitListLengthFromResponse = data.hits.hits
      ? data.hits.hits.length
      : -1;

    return {
      statusCode: 200,
      body: JSON.stringify({
        scrollId: data._scroll_id,
        hitListLength: hitListLengthFromResponse
      })
    };
  } catch (error) {
    // TODO: Use sentry here and don't retun the error to the consumer
    return {
      statusCode: 500,
      body: JSON.stringify({ error, message: "unexpected_failure" })
    };
  }
};
