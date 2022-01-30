const knex = require("knex");
const db = knex({
  client: "postgresql",
  connection: {
    database: "mutitenancy-node",
    user: "postgres",
    password: "postgres",
  },
});
module.exports = db;
