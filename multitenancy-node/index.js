const express = require("express");
const app = express();
const knex = require("./db.js");

app.use(express.json());

app.post("/create-tenant", async (req, res) => {
  const { name, subdomain, adminEmail } = req.body;

  await knex("tenants").insert({
    name,
    subdomain,
    admin_email: adminEmail,
  });

  await knex.schema.createTable(`${subdomain}_users`, (table) => {
    table.uuid("id").defaultTo(knex.raw("uuid_generate_v4()"));
    table.string("first_name");
    table.string("last_name");
    table.string("email").unique();
  });

  res.send("Tenant Created !!");
});

app.post("/create-user", async (req, res) => {
  // const subdomain = req.subdomains[0];

  const { firstName, lastName, email, subdomain } = req.body;

  await knex(`${subdomain}_users`).insert({
    first_name: firstName,
    lastName: lastName,
    email,
  });

  res.send("User Created !!");
});

app.listen(4000, () => {
  console.log("Server listening to Port 4000");
});
