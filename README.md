# Regtic

This is the base repo for the anti-money laundering project which is/was a part of the HardTech Entrepreneurship course at the Technical University of Denmark anno 2019.

[![Netlify Status](https://api.netlify.com/api/v1/badges/bdb923b9-fd7d-46f0-91b5-796690da8448/deploy-status)](https://app.netlify.com/sites/affectionate-swartz-baae10/deploys)

## Structure

The system is using the serverless framework, and has the following structure

```md
└── regtic
├── api
├── app
├── database
└── workers
```

### Api

Serverless functions (mainly written in [Go](https://golang.org/)) in order to facilitate the requirements of the web app.

### App

The app directory includes a Vue based web app that communicates with the API in order to search for companies and display their structure (while checking every person across/company with the sanctions list and with the "Politically Exposed Persons" list).

### Database

Postgres migrations and helper scripts, using the knex.js library

### Workers

This directory includes serverless functions that can parse and aggregate data. Namely this projects uses the following workers:

- `pepworker` retrieves the excel file from the [Danish "Politically Exposed Persons" list](https://www.finanstilsynet.dk), processes/cleanup the data and inserts the people to the database (or updates the existing records if the person is already in the database). Written in [Python](https://www.python.org).
- `sanctionworker` retrieves the xml file from [EU's sanction list](webgate.ec.europa.eu/europeaid/fsd/fsf/public/files/xmlFullSanctionsList_1_1/), processes/cleanup the data and inserts the people to the database (or updates the existing records if the person is already in the database). Written in [Python](https://www.python.org).
- `vat-dk` scraps the entire [Danish VAT registry](http://distribution.virk.dk/cvr-permanent) using elasticsearch, cleans up the data and inserts them to the database. This activity is done sequential instead of in parallel in order to minimize the load on the cvr registry, unfortunately this also means that the worker requires multiple hours to complete it's task. Written in [Node.js](https://nodejs.org/).

## Initial setup

1. Install [Docker](https://www.docker.com/), [Docker Compose](https://docs.docker.com/compose/install/), [NodeJS 12.x](https://nodejs.org/en/download/), [Serverless](https://serverless.com/), [Sam local](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) and [Go](https://golang.org/dl/)

2. `$ git clone git@github.com:alexdor/regtic.git`

3. `$ npm i`

4. `$ cd database && npm i`

5. In base repo run `$ npm run start`

6. Then to get the databse structure `$ cd database && npm run migrate-latest`

7. To stop the app press `ctrl+c`

## Formats

- Country codes [ISO 3166-1](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes), e.g. DK for Denmark
- Use the full VAT number, for each company, as specified [here](https://en.wikipedia.org/wiki/VAT_identification_number)

## Linters

Please install the [black](https://pypi.org/project/black/) linter for python.
