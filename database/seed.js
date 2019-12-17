/* eslint-disable @typescript-eslint/no-var-requires */
const axios = require("axios");
const yml = require("yaml");
const fs = require("fs");

const GREEN = "\x1b[32m";
const RED = "\x1b[31m";
const ADD_DELAY = process.env.ADD_DELAY;
const backendUrl = process.env.BACKEND_URL || "http://localhost:3000";

const file = fs.readFileSync("../serverless.yml", "utf8");
const config = yml.parse(file);
const localApi = axios.create({
  baseURL: `${backendUrl}${(config.custom || {}).version}`,
  timeout: 200000
});

function sleep(ms) {
  return new Promise(resolve => {
    setTimeout(resolve, ms);
  });
}
let breakCounter = 0;
async function cvrWorker() {
  let counter = 0;
  let scrollId;
  let keepGoing = true;

  while (keepGoing) {
    try {
      console.log(
        GREEN,
        "cvr worker:",
        "parsing page:",
        ++counter,
        "more pages:",
        keepGoing
      );
      const { data } = await localApi.post("cvr-parse", { scrollId });
      scrollId = data.scrollId;
      keepGoing = data.hitListLength === 200;
      if (ADD_DELAY) sleep(1000);
    } catch (error) {
      console.error(RED, error);
      if (breakCounter > 4) break;
      breakCounter++;
      counter--;
      sleep(10000);
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
