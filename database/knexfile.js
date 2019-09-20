require("load-environment");

module.exports = {
  development: {
    client: "pg",
    connection: process.env.REGTIC_DATABASE_URL,
    migrations: {
      directory: "migrations"
    },
    seeds: {
      directory: "seeds"
    }
  }
};
