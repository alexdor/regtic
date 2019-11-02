const axios = require("axios");
const yml = require("yaml");
const fs = require("fs");

const GREEN = "\x1b[32m";
const RED = "\x1b[31m";

const backendUrl = process.env.BACKEND_URL || "http://localhost:3000";

const file = fs.readFileSync("../serverless.yml", "utf8");
const config = yml.parse(file);
const localApi = axios.create({
  baseURL: `${backendUrl}${(config.custom || {}).version}`,
  timeout: 200000
});

async function cvrWorker() {
  let counter = 0;
  let scrollId;
  let keepGoing = true;

  while (keepGoing) {
    try {
      console.log(
        GREEN,
        "cvr worker:",
        "page:",
        ++counter,
        "more pages:",
        keepGoing
      );
      const { data } = await localApi.post("cvr-parse", { scrollId });
      scrollId = data.scrollId;
      keepGoing = data.hitListLength === 200;
    } catch (error) {
      console.error(RED, error);
      break;
    }
  }
}

function worker(url) {
  return localApi
    .get(url)
    .then(res => console.log(GREEN, `${url}:`, res.data))
    .catch(error => console.error(RED, `${url}:`, error.message));
}

async function seed() {
  await Promise.all([
    cvrWorker(),
    worker("sanctionworker"),
    worker("pepworker")
  ]);
}

seed();
