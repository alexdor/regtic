const yml = require("yaml");
const fs = require("fs");
const file = fs.readFileSync("../serverless.yml", "utf8");
const config = yml.parse(file);

module.exports = {
  development: {
    client: "pg",
    connection: process.env.REGTIC_DATABASE_URL || (config.custom || {}).dbURL,
    migrations: {
      directory: "migrations"
    },
    seeds: {
      directory: "seeds"
    }
  }
};
