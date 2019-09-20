const yml = require("yaml");
const fs = require("fs");
const file = fs.readFileSync("../serverless.yml", "utf8");
const config = yml.parse(file);

module.exports = {
  development: {
    client: "pg",
    connection: (config.environment || {}).REGTIC_DATABASE_URL,
    migrations: {
      directory: "migrations"
    },
    seeds: {
      directory: "seeds"
    }
  }
};
