[![Netlify Status](https://api.netlify.com/api/v1/badges/bdb923b9-fd7d-46f0-91b5-796690da8448/deploy-status)](https://app.netlify.com/sites/affectionate-swartz-baae10/deploys)

# Regtic

This is the base repo for the anti-money laundering project which is/was a part of the HardTech Entrepreneurship course at the Technical University of Denmark anno 2019.

The system is using the serverless framework, and has the following structure

```
└── regtic
    ├── app
    ├── database
    ├── gateway
    └── workers
```

## Initial setup

1. Install [Docker](https://www.docker.com/), [Docker Compose](https://docs.docker.com/compose/install/), [NodeJS](https://nodejs.org/en/download/), [Serverless](https://serverless.com/), [Sam local](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) and [GO](https://golang.org/dl/)

2. `$ git clone git@github.com:alexdor/regtic.git`

3. `$ npm i`

4. `$ cd database && npm i`

5. In base repo run `$ npm run start`

6. Then to get the databse structure `$ cd database && npm run migrate-latest`

7. To kill the main process press `ctrl+c`

## Formats

- Country codes [ISO 3166-1](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes), e.g. DK for Denmark
- Use the full VAT number, for each company, as specified [here](https://en.wikipedia.org/wiki/VAT_identification_number)

## Linters

Please install the [black](https://pypi.org/project/black/) linter for python.
