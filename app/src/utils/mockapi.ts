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
      info: {
        id: "1234-56789-01234-5678",
        address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
        vat: "DK-12345678",
        type: "IVS",
        startingDate: "11-11-2019Z05:12:34T",
        countryCode: "DK",
        name: "Demo Company ApS",
        status: "warning",
        statusType: "pep",
        statusNotes: "Inherited from beneficiaries",
        owns: [
          { id: "2345-56789-01234-5678", relation: "Legal Owner", ownership: 0.3, votingRight: 0.2 },
          { id: "3456-56789-01234-5678", relation: "Ultimate Beneficial Owner", ownership: 0.14, votingRight: 0.25 }
        ],
        entityType: "company"
      },
      companies: [
        {
          id: "2345-56789-01234-5678",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
          vat: "DK-23456789",
          type: "IVS",
          startingDate: "11-11-2019Z05:12:34T",
          countryCode: "DK",
          name: "Demo Company 2 ApS",
          status: "warning",
          statusType: "pep",
          statusNotes: "Inherited from beneficiaries",
          owns: [
            { id: "1234-67890-12345-6789", relation: "Ultimate Beneficial Owner", ownership: 0.2, votingRight: 0.3 }
          ],
          entityType: "company"
        },
        {
          id: "3456-56789-01234-5678",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
          type: "IVS",
          vat: "DK-34567890",
          startingDate: "11-11-2019Z05:12:34T",
          countryCode: "DK",
          name: "Demo Company 3 ApS",
          status: "warning",
          statusType: "pep",
          source: "https://www.example.org/link/pep-2019-list.pdf",
          owns: [
            { id: "4567-56789-01234-5678", relation: "Legal Owner", ownership: 0.2, votingRight: 0.3 },
            { id: "1234-67890-12345-6789", relation: "Accountant", ownership: 0.2, votingRight: 0.3 },
            { id: "2345-67890-12345-6789", relation: "Ultimate Beneficial Owner", ownership: 0.6, votingRight: 0 }
          ],
          entityType: "company"
        },
        {
          id: "4567-56789-01234-5678",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
          type: "IVS",
          vat: "DK-45678901",
          startingDate: "11-11-2019Z05:12:34T",
          countryCode: "DK",
          name: "Demo Company 4 ApS",
          status: "ok",
          owns: [
            { id: "4567-67890-12345-6789", relation: "Ultimate Beneficial Owner", ownership: 0.25, votingRight: 0.16 }
          ],
          entityType: "company"
        }
      ],
      people: [
        {
          id: "1234-67890-12345-6789",
          name: "Mallory Malicious",
          countryCode: "DK",
          status: "issue",
          statusType: "saction",
          source: "https://www.example.org/link/sanctions-list.pdf",
          entityType: "person",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby"
        },
        {
          id: "2345-67890-12345-6789",
          name: "Phil Pepper",
          status: "warning",
          statusType: "pep",
          source: "https://www.example.org/link/pep-list.pdf",
          countryCode: "DK",
          entityType: "person",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby"
        },
        {
          id: "3456-67890-12345-6789",
          name: "Gary Goodguy",
          status: "ok",
          countryCode: "DK",
          entityType: "person",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby"
        },
        {
          id: "4567-67890-12345-6789",
          name: "Gary Goodguy2",
          status: "ok",
          countryCode: "DK",
          entityType: "person",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby"
        }
      ]
    } as ApiResponse;
  },
  signup: function (email: string) {
    return "";
  }
};


export interface Relationship {
  id: string;
  relation: string;
  ownership: number;
  votingRight: number;
}
export interface CommonFields {
  id: string;
  address?: string;
  countryCode: string; //DK || ZZ
  name: string;
  status: 'ok' | 'warning' | 'issue';
  statusType?: "saction" | "pep";
  source?: string;
  statusNotes?: string;
  owns: Relationship[];
  entityType: 'person' | 'company';
}

export interface Person extends CommonFields { }

export interface Company extends CommonFields {
  type: string;
  vat: string;
  startingDate: string;
  /*adress: string;*/
}
export interface ApiResponse {
  info: Company;
  people: Person[];
  companies: Company[];
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
