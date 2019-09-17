// dotenv handler
require('load-environment')

// imports
const axios = require('axios')

// configurations
const cvrApi = axios.create({
  baseURL: process.env.CVR_BASE_URL,
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  },
  auth: {
    username: process.env.CVR_USER,
    password: process.env.CVR_PASS
  }
})

async function scrollRequest (body, config) {
  const scrollTimeout = config.scrollTimeout
    ? config.scrollTimeout
    : '1m'
  const endpoint = config.initialRequest
    ? `cvr-permanent/_search?scroll=${scrollTimeout}`
    : '_search/scroll'
  return cvrApi.get(endpoint, body)
}

const testQuery = {
  size: 1,
  _source: ["Vrdeltagerperson.navne.navn"],
  query: {
    match_all: {}
  }
}

async function scrollTest () {
  try {
    const intialRequestResult = await scrollRequest(testQuery, { initialRequest: true })
    // const subsequentRequestResult = await scrollRequest({ scroll: '1m', scroll_id: intialRequestResult.data._scroll_id })

    console.log(JSON.stringify(intialRequestResult.data))
  }
  catch (error) {
    console.log('### ERROR INCOMING ###')
    console.log(error)
  }
}

scrollTest()
