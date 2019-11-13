/* eslint-disable @typescript-eslint/no-var-requires */
function getURLFromConf() {
  const yml = require("yaml");
  const fs = require("fs");
  const file = fs.readFileSync("../serverless.yml", "utf8");
  const config = yml.parse(file);
  const url = (config.custom || {}).REGTIC_DATABASE_URL;
  return url.split('"')[1];
}
module.exports = {
  development: {
    client: "pg",
    connection: process.env.REGTIC_DATABASE_URL || getURLFromConf(),
    migrations: {
      directory: "migrations"
    },
    seeds: {
      directory: "seeds"
    }
  }
};
