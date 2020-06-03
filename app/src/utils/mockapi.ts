/* eslint-disable */

import {
  ApiResponsePersonInfo,
  Company,
  ValidationResponse
} from "./interfaces";

function sleep(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

export default {
  async findCompanies(searchStr: string) {
    await sleep(2000);
    return [
      {
        name: "Demo Regtic",
        id: "1234-56789-01234-5678",
        address: "Demo Street 1",
        vat: "DK-12345678",
        countryCode: "DK",
        type: "IVS"
      },
      {
        name: "Regtic Example",
        id: "2345-56789-01234-5678",
        address: "Demo Street 1",
        vat: "DK-23456789",
        countryCode: "DK",
        type: "IVS"
      },
      {
        name: "Regtic A/S",
        id: "3456-56789-01234-5678",
        address: "Demo Street 2",
        vat: "DK-34567890",
        countryCode: "DK",
        type: "IVS"
      },
      {
        name: "Regtic Finance",
        id: "4567-56789-01234-5678",
        address: "Test Street 100",
        vat: "DK-45678901",
        countryCode: "DK",
        type: "IVS"
      },
      {
        name: "Regtic Law",
        id: "5678-56789-01234-5678",
        address: "Test Street 100",
        vat: "DK-56789012",
        countryCode: "DK",
        type: "IVS"
      },
      {
        name: "Listed Regtic",
        id: "6789-56789-01234-5678",
        address: "Demo Blv. 5",
        vat: "DK-67890123",
        countryCode: "DK",
        type: "IVS"
      },
      {
        name: "Regtic Test Name",
        id: "7890-56789-01234-5678",
        address: "DTU Lyngby 101 A",
        vat: "DK-78901234",
        countryCode: "DK",
        type: "IVS"
      }
    ] as Company[];
  },
  async validateCompany(id: string) {
    await sleep(300);

    return {
      info: {
        id: "a52701c3-801b-4d5d-9caa-2f15f81c4972",
        address: "7190 Billund Åstvej 1",
        vat: "DK-54562519",
        startingDate: "1975-12-19",
        countryCode: "DK",
        name: "LEGO A/S",
        type: "A/S"
      },
      people: [
        {
          id: "93f69023-1c70-406b-bca4-29007bb9b1c4",
          countryCode: "DK",
          name: "Mogens Johansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "cbe6e57b-9e3a-49a5-b166-5d9a6d5c1f66",
          countryCode: "ZZ",
          name: "Lars Erik Kann-Rasmussen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "36e27352-1e26-412d-a943-a079aab96563",
          countryCode: "ZZ",
          name: "Vagn Holck-Andersen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "d931ec83-2cb1-4631-8117-aef1c1e3b8ab",
          countryCode: "DK",
          name: "Mads Øvlisen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "e17a658e-3553-4fbf-b4bf-43bbc55a56d2",
          countryCode: "ZZ",
          name: "John Bøndergaard",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "987b30c3-b64d-4352-a944-cb67308f54ce",
          countryCode: "ZZ",
          name: "Søren Thorup Sørensen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "5f209468-771d-4185-bfc0-8fee7193168e",
          countryCode: "DK",
          name: "Niels Jacobsen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "029637ad-6bd5-40a6-9b7e-9cab20426a84",
          countryCode: "DK",
          name: "Christian Majgaard Nielsen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "cc1cdc81-0c2b-405b-a659-332f942c7103",
          countryCode: "CH",
          name: "Poul Plougmann",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "bbe1d8e7-ada4-46a3-9113-d84b7ad1321d",
          countryCode: "ZZ",
          name: "Thomas Kirk Kristiansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "7d82c2ab-89c9-4b5d-858f-2a9bec0ce6e8",
          countryCode: "DK",
          name: "Eva Merete Søfelde Berneke",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "534572c4-f9f8-4945-ab46-70e75104ab52",
          countryCode: "DK",
          name: "Torsten Erik Rasmussen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "f1906da5-bfea-4196-9adc-66435789f8f4",
          countryCode: "ZZ",
          name: "Bent Skov",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "da34d424-c72b-4913-ad81-d21cb2457d5f",
          countryCode: "ZZ",
          name: "Kåre Schultz",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "a4768602-68cb-448e-9466-647084a1481c",
          countryCode: "ZZ",
          name: "Ove Gundtoft",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "a5485757-574d-4f94-952d-daf3d1cba18c",
          countryCode: "ZZ",
          name: "Kjeld Kirk Kristiansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "004e67d0-dcd4-4846-8bb3-05af3e4cafa9",
          countryCode: "DK",
          name: "Stig Christensen",
          checkStatus: "WARNING",
          statusType: "PEP",
          source:
            "https://www.finanstilsynet.dk/~/media/Tal-og-fakta/2019/PEP/PEPliste-opdateret-den-131219-xlsx.xlsx?la=da",
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "f56b44e2-8e06-47cd-b82f-23614fe310f3",
          countryCode: "ZZ",
          name: "Arne Christian Johansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "eaffe7e9-e8fd-45a6-bde0-f29274790e35",
          countryCode: "ZZ",
          name: "Erik Viuff Quistgaard",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "d79c11a2-6a94-4511-884c-fd3e4898c82d",
          countryCode: "DK",
          name: "Torben Ballegaard Sørensen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "884938db-7413-48f6-add7-5099916a0b4e",
          countryCode: "ZZ",
          name: "Niels Bjørn Christiansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "a8b298b4-a224-49c2-8d18-de7c64da57c9",
          countryCode: "ZZ",
          name: "Niels Christian Jensen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "783be98f-5b77-4c91-81fe-10c75cf9f985",
          countryCode: "GB",
          name: "Karsten Borch",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "f7954aae-825b-4c8a-b3dc-981e99f66bda",
          countryCode: "DK",
          name: "Kjeld Møller Pedersen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "3ff62ed1-cd42-44bf-8a2b-96ebb4eb0011",
          countryCode: "NL",
          name: "Anders Christer Moberg",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "691aa937-8106-4c88-8cf5-137bef29690a",
          countryCode: "DK",
          name: "Jens Jesper Ovesen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "35eee2f6-7cd7-4a2a-a215-94ea9053d4b0",
          countryCode: "ZZ",
          name: "Jens Nordahl Ravnbøl",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "5ef22ae8-255f-4d2b-8049-7b6d8cdca85e",
          countryCode: "DK",
          name: "Caroline Søeborg Ahlefeldt-Laurvig-Bille",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "f642c60c-9032-401e-a2e1-96612406deea",
          countryCode: "ZZ",
          name: "Stig Toftgaard",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "6d0d5b9b-a732-4d56-8501-5e614b34ee24",
          countryCode: "ZZ",
          name: "Jørgen Vig Knudstorp",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "ca4352de-5ef9-4ead-b63f-6f5969b6a48c",
          countryCode: "ZZ",
          name: "Godtfred Kirk Christiansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "ce1965d2-a2b6-4ca3-ba6f-48eb0b9436a3",
          countryCode: "MX",
          name: "Sten Daugaard",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "ed95e040-7bf6-42cb-bb3a-68bcb46a1852",
          countryCode: "DK",
          name: "Jan Thorsgaard Nielsen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "49a9f062-b431-4d9a-92f7-6844fc9b2f95",
          countryCode: "DK",
          name: "Poul Hartvig Nielsen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "ae857e78-eacb-4bac-b877-4c40574ced90",
          countryCode: "DK",
          name: "Mogens Johansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "b5d2ecf1-2823-4bd0-ab3c-797108191884",
          countryCode: "ZZ",
          name: "Lars Erik Kann-Rasmussen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "cf205f0f-f6f0-4a0a-9cf0-d84921b7c9e0",
          countryCode: "ZZ",
          name: "Vagn Holck-Andersen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "622c1357-c818-4f8c-94bf-3a13ea079885",
          countryCode: "DK",
          name: "Mads Øvlisen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "237d6ba2-fd6d-494f-b9ba-ef989330f293",
          countryCode: "ZZ",
          name: "John Bøndergaard",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "6be1878d-36c6-4f72-8eaa-173d49ef9ca9",
          countryCode: "ZZ",
          name: "Søren Thorup Sørensen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "c60bf885-f669-4e50-8e3d-79c6af574b78",
          countryCode: "DK",
          name: "Niels Jacobsen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "ba058341-c96b-48e4-a035-4761d842322f",
          countryCode: "DK",
          name: "Christian Majgaard Nielsen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "d4f497ca-9dde-497f-a976-34da1eb794dd",
          countryCode: "CH",
          name: "Poul Plougmann",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "d1b994f5-813b-494d-8c5b-c807fe1df999",
          countryCode: "ZZ",
          name: "Thomas Kirk Kristiansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "9eef85b7-44cb-4454-9ad1-88ece7570523",
          countryCode: "DK",
          name: "Eva Merete Søfelde Berneke",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "000c8897-8101-46ac-871f-15ce0118363a",
          countryCode: "DK",
          name: "Torsten Erik Rasmussen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "a2012eb8-4bd7-4106-93c2-3c61ce4168b6",
          countryCode: "ZZ",
          name: "Bent Skov",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "076803d1-2c51-4211-b8c7-175ea9fafe9d",
          countryCode: "ZZ",
          name: "Kåre Schultz",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "13a0e8cf-e8c0-49e8-96e7-5de24acd07c8",
          countryCode: "ZZ",
          name: "Ove Gundtoft",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "d1730892-2df1-4740-86bd-85418a884701",
          countryCode: "ZZ",
          name: "Kjeld Kirk Kristiansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "f5139934-b528-49d5-a3bf-4601de28a514",
          countryCode: "DK",
          name: "Stig Christensen",
          checkStatus: "WARNING",
          statusType: "PEP",
          source:
            "https://www.finanstilsynet.dk/~/media/Tal-og-fakta/2019/PEP/PEPliste-opdateret-den-131219-xlsx.xlsx?la=da",
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "8853826d-cbc8-4b8f-aaff-94500dacca9d",
          countryCode: "ZZ",
          name: "Arne Christian Johansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "607ae407-4dec-4eb9-91d6-fec70046c934",
          countryCode: "ZZ",
          name: "Erik Viuff Quistgaard",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "66c46746-a1fb-498d-81cb-b7942420ad46",
          countryCode: "DK",
          name: "Torben Ballegaard Sørensen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "7be93498-303b-4a81-95b7-db0f89f199b1",
          countryCode: "ZZ",
          name: "Niels Bjørn Christiansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "aa1920af-d063-4680-8e9e-d199161a8e21",
          countryCode: "ZZ",
          name: "Niels Christian Jensen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "b0d87f40-7a98-4055-bbd2-9e023e459b1a",
          countryCode: "GB",
          name: "Karsten Borch",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "0576c3b8-9586-4a90-956a-228055737c93",
          countryCode: "DK",
          name: "Kjeld Møller Pedersen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "ed58ab8e-9a8f-4105-8583-4918c9553a6a",
          countryCode: "NL",
          name: "Anders Christer Moberg",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "232b53f2-3674-4b9d-a495-6964e8af5750",
          countryCode: "DK",
          name: "Jens Jesper Ovesen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "b13d3a71-074f-4fc4-af14-0c21a88b552d",
          countryCode: "ZZ",
          name: "Jens Nordahl Ravnbøl",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "b8f1c330-4f5a-4bff-93ab-71fe0a3b1760",
          countryCode: "DK",
          name: "Caroline Søeborg Ahlefeldt-Laurvig-Bille",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "55e0f874-a43a-455c-ab09-df9e20bab675",
          countryCode: "ZZ",
          name: "Stig Toftgaard",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "6239be49-52d6-47a9-867b-75e74e047348",
          countryCode: "ZZ",
          name: "Jørgen Vig Knudstorp",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "00dcc2e7-3d1c-4387-bcaa-43dc239c189a",
          countryCode: "ZZ",
          name: "Godtfred Kirk Christiansen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "743e37a2-83a5-4622-b429-1866d8c6be1b",
          countryCode: "MX",
          name: "Sten Daugaard",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "b26e7343-1937-4f08-9bc0-f456250b0d0d",
          countryCode: "DK",
          name: "Jan Thorsgaard Nielsen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        },
        {
          id: "319e9e3e-7b76-44be-a481-89059c76a87b",
          countryCode: "DK",
          name: "Poul Hartvig Nielsen",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: null,
          entityType: "PERSON"
        }
      ],
      companies: [
        {
          id: "a52701c3-801b-4d5d-9caa-2f15f81c4972",
          address: "7190 Billund Åstvej 1",
          countryCode: "DK",
          name: "LEGO A/S",
          checkStatus: "OK",
          statusType: null,
          source: null,
          ownedBy: [
            {
              id: "93f69023-1c70-406b-bca4-29007bb9b1c4",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "cbe6e57b-9e3a-49a5-b166-5d9a6d5c1f66",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "36e27352-1e26-412d-a943-a079aab96563",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "d931ec83-2cb1-4631-8117-aef1c1e3b8ab",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "e17a658e-3553-4fbf-b4bf-43bbc55a56d2",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "987b30c3-b64d-4352-a944-cb67308f54ce",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "5f209468-771d-4185-bfc0-8fee7193168e",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "029637ad-6bd5-40a6-9b7e-9cab20426a84",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "cc1cdc81-0c2b-405b-a659-332f942c7103",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "bbe1d8e7-ada4-46a3-9113-d84b7ad1321d",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "7d82c2ab-89c9-4b5d-858f-2a9bec0ce6e8",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "534572c4-f9f8-4945-ab46-70e75104ab52",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "f1906da5-bfea-4196-9adc-66435789f8f4",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "da34d424-c72b-4913-ad81-d21cb2457d5f",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "a4768602-68cb-448e-9466-647084a1481c",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "a5485757-574d-4f94-952d-daf3d1cba18c",
              relation: [
                "board of directors",
                "direction",
                "ultimate beneficial owner"
              ],
              ownership: 0.2027,
              votingRights: 0.465,
              entityType: "PERSON"
            },
            {
              id: "004e67d0-dcd4-4846-8bb3-05af3e4cafa9",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "f56b44e2-8e06-47cd-b82f-23614fe310f3",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "eaffe7e9-e8fd-45a6-bde0-f29274790e35",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "d79c11a2-6a94-4511-884c-fd3e4898c82d",
              relation: ["board of directors", "direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "884938db-7413-48f6-add7-5099916a0b4e",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "a8b298b4-a224-49c2-8d18-de7c64da57c9",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "783be98f-5b77-4c91-81fe-10c75cf9f985",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "f7954aae-825b-4c8a-b3dc-981e99f66bda",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "3ff62ed1-cd42-44bf-8a2b-96ebb4eb0011",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "691aa937-8106-4c88-8cf5-137bef29690a",
              relation: ["board of directors", "direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "35eee2f6-7cd7-4a2a-a215-94ea9053d4b0",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "5ef22ae8-255f-4d2b-8049-7b6d8cdca85e",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "f642c60c-9032-401e-a2e1-96612406deea",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "6d0d5b9b-a732-4d56-8501-5e614b34ee24",
              relation: ["board of directors", "direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "ca4352de-5ef9-4ead-b63f-6f5969b6a48c",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "ce1965d2-a2b6-4ca3-ba6f-48eb0b9436a3",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "ed95e040-7bf6-42cb-bb3a-68bcb46a1852",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "49a9f062-b431-4d9a-92f7-6844fc9b2f95",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "ae857e78-eacb-4bac-b877-4c40574ced90",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "b5d2ecf1-2823-4bd0-ab3c-797108191884",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "cf205f0f-f6f0-4a0a-9cf0-d84921b7c9e0",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "622c1357-c818-4f8c-94bf-3a13ea079885",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "237d6ba2-fd6d-494f-b9ba-ef989330f293",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "6be1878d-36c6-4f72-8eaa-173d49ef9ca9",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "c60bf885-f669-4e50-8e3d-79c6af574b78",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "ba058341-c96b-48e4-a035-4761d842322f",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "d4f497ca-9dde-497f-a976-34da1eb794dd",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "d1b994f5-813b-494d-8c5b-c807fe1df999",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "9eef85b7-44cb-4454-9ad1-88ece7570523",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "000c8897-8101-46ac-871f-15ce0118363a",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "a2012eb8-4bd7-4106-93c2-3c61ce4168b6",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "076803d1-2c51-4211-b8c7-175ea9fafe9d",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "13a0e8cf-e8c0-49e8-96e7-5de24acd07c8",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "d1730892-2df1-4740-86bd-85418a884701",
              relation: [
                "board of directors",
                "direction",
                "ultimate beneficial owner"
              ],
              ownership: 0.2027,
              votingRights: 0.465,
              entityType: "PERSON"
            },
            {
              id: "f5139934-b528-49d5-a3bf-4601de28a514",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "8853826d-cbc8-4b8f-aaff-94500dacca9d",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "607ae407-4dec-4eb9-91d6-fec70046c934",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "66c46746-a1fb-498d-81cb-b7942420ad46",
              relation: ["board of directors", "direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "7be93498-303b-4a81-95b7-db0f89f199b1",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "aa1920af-d063-4680-8e9e-d199161a8e21",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "b0d87f40-7a98-4055-bbd2-9e023e459b1a",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "0576c3b8-9586-4a90-956a-228055737c93",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "ed58ab8e-9a8f-4105-8583-4918c9553a6a",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "232b53f2-3674-4b9d-a495-6964e8af5750",
              relation: ["board of directors", "direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "b13d3a71-074f-4fc4-af14-0c21a88b552d",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "b8f1c330-4f5a-4bff-93ab-71fe0a3b1760",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "55e0f874-a43a-455c-ab09-df9e20bab675",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "6239be49-52d6-47a9-867b-75e74e047348",
              relation: ["board of directors", "direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "00dcc2e7-3d1c-4387-bcaa-43dc239c189a",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "743e37a2-83a5-4622-b429-1866d8c6be1b",
              relation: ["direction"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "b26e7343-1937-4f08-9bc0-f456250b0d0d",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            },
            {
              id: "319e9e3e-7b76-44be-a481-89059c76a87b",
              relation: ["board of directors"],
              ownership: 0,
              votingRights: 0,
              entityType: "PERSON"
            }
          ],
          entityType: "COMPANY",
          vat: "DK-54562519",
          startingDate: "1975-12-19",
          status: null,
          statusNotes: null,
          type: "A/S"
        }
      ]
    } as ValidationResponse;
  },
  async getPerson(id: string) {
    await sleep(300);

    return {
      id: "1234-67890-12345-6789",
      name: "Mallory Malicious",
      countryCode: "DK",
      entityType: "PERSON",
      address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
      checkStatus: "ISSUE",
      statusType: "SANCTION",
      source: "https://www.example.org/lists/sanction-list.pdf",
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
          statusNotes: "Inherited from beneficiaries",
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
          statusNotes: "Inherited from beneficiaries",
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
  signup: function(email: string) {
    return "";
  }
};
