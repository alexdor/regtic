/* eslint-disable */

import axios, { AxiosResponse } from "axios";

function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

export default {
  async getCountries(callback: any) {
    axios.get("https://restcountries.eu/rest/v2/all").then(callback);
  },
  async findCompanies(searchStr: string) {
    await sleep(2000);
    return [
      {
        name: "Demo Regtic",
        id: "1234-56789-01234-5678",
        address: "Demo Street 1",
        vat: "DK-12345678",
        country: "Denmark / DK",
        type: "IVS"
      },
      {
        name: "Regtic Example",
        id: "2345-56789-01234-5678",
        address: "Demo Street 1",
        vat: "DK-23456789",
        country: "Denmark / DK",
        type: "IVS"
      },
      {
        name: "Regtic A/S",
        id: "3456-56789-01234-5678",
        address: "Demo Street 2",
        vat: "DK-34567890",
        country: "Denmark / DK",
        type: "IVS"
      },
      {
        name: "Regtic Finance",
        id: "4567-56789-01234-5678",
        address: "Test Street 100",
        vat: "DK-45678901",
        country: "Denmark / DK",
        type: "IVS"
      },
      {
        name: "Regtic Law",
        id: "5678-56789-01234-5678",
        address: "Test Street 100",
        vat: "DK-56789012",
        country: "Denmark / DK",
        type: "IVS"
      },
      {
        name: "Listed Regtic",
        id: "6789-56789-01234-5678",
        address: "Demo Blv. 5",
        vat: "DK-67890123",
        country: "Denmark / DK",
        type: "IVS"
      },
      {
        name: "Regtic Test Name",
        id: "7890-56789-01234-5678",
        address: "DTU Lyngby 101 A",
        vat: "DK-78901234",
        country: "Denmark / DK",
        type: "IVS"
      }
    ];
  },
  async validateCompany(id: string) {
    await sleep(300);

    return {
      "info": {
        "id": "68cc018c-2ce3-4d48-a69b-3a389cd5319c",
        "address": "7190 Billund Åstvej 1",
        "vat": "DK-54562519",
        "startingDate": "1975-12-19",
        "countryCode": "DK",
        "name": "LEGO A/S",
        type: "A/S",
        entityType: "COMPANY",
        ownedBy: [
          {
            "id": "3410d139-4de2-4ae7-bd46-81ac5d04644f",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "7ced1230-daf4-4a79-abf3-ef797ad48008",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "b5f873e9-6f73-4143-b3e0-5d7c5a056a05",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "2c23fa63-50fc-4241-a976-0abdd0a95ef1",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "5189782d-f34b-48bc-ad24-9a29e65321d5",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "ed5e7d8e-6686-446d-827a-956141af42a7",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "62141fe0-fdad-4644-8749-ec1fe1b0549c",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "2af60e4c-7e4d-492d-ab0b-8a9d6bd70578",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "c7b4b229-d0c9-4427-81e8-dea69a8c1c37",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "2c986b47-88b6-43cd-ab58-9f8f2107266f",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "475556fa-4940-4ae8-a800-9cfe72ad19bc",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "24316a30-9202-4209-afa9-0308b96843b1",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "da668682-db5c-47e4-83db-bb002e7f9190",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "6ada8692-c4ab-42fd-a355-48bd7c6c05ee",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "fabfa59c-8d96-4a9f-a8d9-8b6e663be921",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "da368f65-d189-44ec-a3ce-756d182411d0",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "dbe546ce-8534-474d-be9f-392092cde1a4",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "5765fdd8-b672-4103-9d4e-768f033a3618",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "0ad2aeaf-19b8-42cd-a71c-5d96cb2bb18d",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "362b9e04-2aff-4a0c-a9e8-6be4bbeb6637",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "0d01400e-a65c-46d1-8ba6-d4b4e4ebdd14",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "1656cf7c-c44c-418a-a31b-ba338a856718",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "726999b7-c744-4c61-ba3b-55b134a3c965",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "008addc7-2f3c-4bb0-9391-6a6672a24772",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "f99a2cd9-fb04-4a5f-a911-97cf17493e8f",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "1c895a2f-289e-4056-9b00-7d7993e71f47",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "1fe934dd-23ae-456c-9b70-499e903282d5",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "23c25cdc-7765-49d9-a05d-06b2f140e34e",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "48449e7f-8c95-48c6-bcbe-13222f23e972",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "55ce741b-e169-41d0-9ed0-2fb51d1e98cd",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "98102179-db71-4d10-bef7-e4b8db0095ee",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "eb07181a-8755-43f7-9244-27ceffc18bf5",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "f2369a40-9a54-4ed9-8a4a-978d2117df01",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          },
          {
            "id": "22a67306-4cf4-4a6a-857a-86bb8656684f",
            "relation": [],
            "ownership": 0,
            "votingRights": 0,
            "entityType": "PERSON"
          }
        ],
        checkStatus: "WARNING"
      },
      "people": [
        {
          "id": "3410d139-4de2-4ae7-bd46-81ac5d04644f",
          "countryCode": "DK",
          "name": "Mogens Johansen",
          "entityType": "PERSON"
        },
        {
          "id": "7ced1230-daf4-4a79-abf3-ef797ad48008",
          "countryCode": "ZZ",
          "name": "Lars Erik Kann-Rasmussen",
          "entityType": "PERSON"
        },
        {
          "id": "b5f873e9-6f73-4143-b3e0-5d7c5a056a05",
          "countryCode": "ZZ",
          "name": "Vagn Holck-Andersen",
          "entityType": "PERSON"
        },
        {
          "id": "2c23fa63-50fc-4241-a976-0abdd0a95ef1",
          "countryCode": "DK",
          "name": "Mads Øvlisen",
          "entityType": "PERSON"
        },
        {
          "id": "5189782d-f34b-48bc-ad24-9a29e65321d5",
          "countryCode": "ZZ",
          "name": "John Bøndergaard",
          "entityType": "PERSON"
        },
        {
          "id": "ed5e7d8e-6686-446d-827a-956141af42a7",
          "countryCode": "ZZ",
          "name": "Søren Thorup Sørensen",
          "entityType": "PERSON"
        },
        {
          "id": "62141fe0-fdad-4644-8749-ec1fe1b0549c",
          "countryCode": "DK",
          "name": "Niels Jacobsen",
          "entityType": "PERSON"
        },
        {
          "id": "2af60e4c-7e4d-492d-ab0b-8a9d6bd70578",
          "countryCode": "DK",
          "name": "Christian Majgaard Nielsen",
          "entityType": "PERSON"
        },
        {
          "id": "c7b4b229-d0c9-4427-81e8-dea69a8c1c37",
          "countryCode": "CH",
          "name": "Poul Plougmann",
          "entityType": "PERSON"
        },
        {
          "id": "2c986b47-88b6-43cd-ab58-9f8f2107266f",
          "countryCode": "ZZ",
          "name": "Thomas Kirk Kristiansen",
          "entityType": "PERSON"
        },
        {
          "id": "475556fa-4940-4ae8-a800-9cfe72ad19bc",
          "countryCode": "DK",
          "name": "Eva Merete Søfelde Berneke",
          "entityType": "PERSON"
        },
        {
          "id": "24316a30-9202-4209-afa9-0308b96843b1",
          "countryCode": "DK",
          "name": "Torsten Erik Rasmussen",
          "entityType": "PERSON"
        },
        {
          "id": "da668682-db5c-47e4-83db-bb002e7f9190",
          "countryCode": "ZZ",
          "name": "Bent Skov",
          "entityType": "PERSON"
        },
        {
          "id": "6ada8692-c4ab-42fd-a355-48bd7c6c05ee",
          "countryCode": "ZZ",
          "name": "Kåre Schultz",
          "entityType": "PERSON"
        },
        {
          "id": "fabfa59c-8d96-4a9f-a8d9-8b6e663be921",
          "countryCode": "ZZ",
          "name": "Ove Gundtoft",
          "entityType": "PERSON"
        },
        {
          "id": "da368f65-d189-44ec-a3ce-756d182411d0",
          "countryCode": "ZZ",
          "name": "Kjeld Kirk Kristiansen",
          "entityType": "PERSON"
        },
        {
          "id": "dbe546ce-8534-474d-be9f-392092cde1a4",
          "countryCode": "DK",
          "name": "Stig Christensen",
          "checkStatus": "WARNING",
          "statusType": "PEP",
          "source": "https://www.finanstilsynet.dk/~/media/Tal-og-fakta/2019/PEP/PEPliste-opdateret-den-131219-xlsx.xlsx?la=da",
          "entityType": "PERSON"
        },
        {
          "id": "5765fdd8-b672-4103-9d4e-768f033a3618",
          "countryCode": "ZZ",
          "name": "Arne Christian Johansen",
          "entityType": "PERSON"
        },
        {
          "id": "0ad2aeaf-19b8-42cd-a71c-5d96cb2bb18d",
          "countryCode": "ZZ",
          "name": "Erik Viuff Quistgaard",
          "entityType": "PERSON"
        },
        {
          "id": "362b9e04-2aff-4a0c-a9e8-6be4bbeb6637",
          "countryCode": "DK",
          "name": "Torben Ballegaard Sørensen",
          "entityType": "PERSON"
        },
        {
          "id": "0d01400e-a65c-46d1-8ba6-d4b4e4ebdd14",
          "countryCode": "ZZ",
          "name": "Niels Bjørn Christiansen",
          "entityType": "PERSON"
        },
        {
          "id": "1656cf7c-c44c-418a-a31b-ba338a856718",
          "countryCode": "ZZ",
          "name": "Niels Christian Jensen",
          "entityType": "PERSON"
        },
        {
          "id": "726999b7-c744-4c61-ba3b-55b134a3c965",
          "countryCode": "GB",
          "name": "Karsten Borch",
          "entityType": "PERSON"
        },
        {
          "id": "008addc7-2f3c-4bb0-9391-6a6672a24772",
          "countryCode": "DK",
          "name": "Kjeld Møller Pedersen",
          "entityType": "PERSON"
        },
        {
          "id": "f99a2cd9-fb04-4a5f-a911-97cf17493e8f",
          "countryCode": "NL",
          "name": "Anders Christer Moberg",
          "entityType": "PERSON"
        },
        {
          "id": "1c895a2f-289e-4056-9b00-7d7993e71f47",
          "countryCode": "DK",
          "name": "Jens Jesper Ovesen",
          "entityType": "PERSON"
        },
        {
          "id": "1fe934dd-23ae-456c-9b70-499e903282d5",
          "countryCode": "ZZ",
          "name": "Jens Nordahl Ravnbøl",
          "entityType": "PERSON"
        },
        {
          "id": "23c25cdc-7765-49d9-a05d-06b2f140e34e",
          "countryCode": "DK",
          "name": "Caroline Søeborg Ahlefeldt-Laurvig-Bille",
          "entityType": "PERSON"
        },
        {
          "id": "48449e7f-8c95-48c6-bcbe-13222f23e972",
          "countryCode": "ZZ",
          "name": "Stig Toftgaard",
          "entityType": "PERSON"
        },
        {
          "id": "55ce741b-e169-41d0-9ed0-2fb51d1e98cd",
          "countryCode": "ZZ",
          "name": "Jørgen Vig Knudstorp",
          "entityType": "PERSON"
        },
        {
          "id": "98102179-db71-4d10-bef7-e4b8db0095ee",
          "countryCode": "ZZ",
          "name": "Godtfred Kirk Christiansen",
          "entityType": "PERSON"
        },
        {
          "id": "eb07181a-8755-43f7-9244-27ceffc18bf5",
          "countryCode": "MX",
          "name": "Sten Daugaard",
          "entityType": "PERSON"
        },
        {
          "id": "f2369a40-9a54-4ed9-8a4a-978d2117df01",
          "countryCode": "DK",
          "name": "Jan Thorsgaard Nielsen",
          "entityType": "PERSON"
        },
        {
          "id": "22a67306-4cf4-4a6a-857a-86bb8656684f",
          "countryCode": "DK",
          "name": "Poul Hartvig Nielsen",
          "entityType": "PERSON"
        }
      ],
      "companies": [
        {
          "id": "68cc018c-2ce3-4d48-a69b-3a389cd5319c",
          "address": "7190 Billund Åstvej 1",
          "countryCode": "DK",
          "name": "LEGO A/S",
          "ownedBy": [
            {
              "id": "3410d139-4de2-4ae7-bd46-81ac5d04644f",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "7ced1230-daf4-4a79-abf3-ef797ad48008",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "b5f873e9-6f73-4143-b3e0-5d7c5a056a05",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "2c23fa63-50fc-4241-a976-0abdd0a95ef1",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "5189782d-f34b-48bc-ad24-9a29e65321d5",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "ed5e7d8e-6686-446d-827a-956141af42a7",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "62141fe0-fdad-4644-8749-ec1fe1b0549c",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "2af60e4c-7e4d-492d-ab0b-8a9d6bd70578",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "c7b4b229-d0c9-4427-81e8-dea69a8c1c37",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "2c986b47-88b6-43cd-ab58-9f8f2107266f",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "475556fa-4940-4ae8-a800-9cfe72ad19bc",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "24316a30-9202-4209-afa9-0308b96843b1",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "da668682-db5c-47e4-83db-bb002e7f9190",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "6ada8692-c4ab-42fd-a355-48bd7c6c05ee",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "fabfa59c-8d96-4a9f-a8d9-8b6e663be921",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "da368f65-d189-44ec-a3ce-756d182411d0",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "dbe546ce-8534-474d-be9f-392092cde1a4",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "5765fdd8-b672-4103-9d4e-768f033a3618",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "0ad2aeaf-19b8-42cd-a71c-5d96cb2bb18d",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "362b9e04-2aff-4a0c-a9e8-6be4bbeb6637",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "0d01400e-a65c-46d1-8ba6-d4b4e4ebdd14",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "1656cf7c-c44c-418a-a31b-ba338a856718",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "726999b7-c744-4c61-ba3b-55b134a3c965",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "008addc7-2f3c-4bb0-9391-6a6672a24772",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "f99a2cd9-fb04-4a5f-a911-97cf17493e8f",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "1c895a2f-289e-4056-9b00-7d7993e71f47",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "1fe934dd-23ae-456c-9b70-499e903282d5",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "23c25cdc-7765-49d9-a05d-06b2f140e34e",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "48449e7f-8c95-48c6-bcbe-13222f23e972",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "55ce741b-e169-41d0-9ed0-2fb51d1e98cd",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "98102179-db71-4d10-bef7-e4b8db0095ee",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "eb07181a-8755-43f7-9244-27ceffc18bf5",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "f2369a40-9a54-4ed9-8a4a-978d2117df01",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            },
            {
              "id": "22a67306-4cf4-4a6a-857a-86bb8656684f",
              "relation": [],
              "ownership": 0,
              "votingRights": 0,
              "entityType": "PERSON"
            }
          ],
          "vat": "DK-54562519",
          "startingDate": "1975-12-19",
          "type": "A/S",
          checkStatus: "OK",
          entityType: "COMPANY"
        }
      ],
      "errors": []
    } as ValidationResponse;
  },
  async getPerson(id: string) {
    await sleep(3000);

    return {
      id: "1234-67890-12345-6789",
      name: "Mallory Malicious",
      countryCode: "DK",
      entityType: "PERSON",
      address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
      checkStatus: "ISSUE",
      statusType: "SANCTION",
      owns: [
        {
          id: "3456-56789-01234-5678",
          relation: ["Accountant"],
          ownership: 0.2,
          votingRight: 0.3,
          name: "Demo Company 3 ApS",
          vat: "DK-34567890",
          type: "IVS",
          checkStatus: "ISSUE",
          statusType: "SANCTION",
          source: "Inherited from beneficiaries",
          countryCode: "ZZ",
          entityType: "COMPANY"
        },
        {
          id: "2345-56789-01234-5678",
          relation: ["Ultimate Beneficial Owner"],
          ownership: 0.2,
          votingRight: 0.3,
          name: "Demo Company 2 ApS",
          vat: "DK-23456789",
          type: "IVS",
          checkStatus: "ISSUE",
          statusType: "SANCTION",
          source: "Inherited from beneficiaries",
          countryCode: "ZZ",
          entityType: "COMPANY"
        }
      ],
      detectedLists: [
        {
          status: "ISSUE",
          type: "SANCTION",
          source: "https://www.example.org/link/sanctions-list.pdf",
          updated: "11-11-2019Z05:12:34T"
        },
        {
          status: "WARNING",
          type: "PEP",
          source: "https://www.example.org/link/pep-2019-list.pdf",
          updated: "11-10-2019Z05:12:34T"
        }
      ]
    } as ApiResponsePersonInfo;
  },
  signup: function (email: string) {
    return "";
  }
};

type EntityType = "PERSON" | "COMPANY";
type CheckStatusType = "OK" | "WARNING" | "ISSUE"; //Should we capitalize this?
type BadType = "PEP" | "SANCTION"
export interface Relationship {
  id: string;
  relation: string[];
  ownership: number;
  votingRights: number;
  entityType: EntityType;
}
export interface Company extends CommonFields {
  address?: string;
  vat: string;
  startingDate?: string;
  status?: string;
  statusNotes?: string;
  type?: string;
}
export interface CommonFields {
  id: string;
  countryCode: string;
  name: string;
  checkStatus: CheckStatusType;
  statusType?: BadType;
  source?: string;
  ownedBy?: Relationship[];
  entityType: EntityType;
}
export interface Person extends CommonFields { }
export interface ValidationResponse {
  info: Company;
  people: Person[];
  companies: Company[];
  errors: any[];
}



export interface DetectedList {
  status: CheckStatusType;
  type?: BadType;
  source: string;
  updated: string;
}

export interface OwnedCompany extends CommonFields {
  relation: string[];
  ownership: number;
  votingRight: number;
  vat: string;
  type: string;
}

export interface ApiResponsePersonInfo extends Person {
  owns: OwnedCompany[];
  detectedLists: DetectedList[];
}




/*export interface Company {
  id: string;
  address: string;
  vat: string;
  starting_date: string;
  country_code: string;
  updated_at: Date;
  created_at: Date;
  name: string;
  status?: string;
  status_notes?: string;
  rank?: number;
  name_vector?: string;
}*/

export interface BadPerson extends CommonPersonFields {
  type: string;
  source?: string;
  address?: string;
}

/*export interface Person extends CommonPersonFields {
  first_name: string;
  last_name: string;
  country_code: string;
}*/
export interface CommonPersonFields {
  id: string;
  full_name: string;
  updated_at: Date;
  created_at: Date;
  name_vector?: string;
}

export interface People {
  bad?: BadPerson[];
  warning: BadPerson[];
  good: Person[];
}
export interface Errors {
  msg: string;
}
export interface CheckCompanyData {
  info: Company;
  companies: Company[];
  people: People;
  errors?: Errors[];
}

export interface FindCompaniesData {
  companies: Company[];
}
