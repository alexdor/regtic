/* eslint-disable */

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
    await sleep(3000);

    return {
      info: {
        id: "1234-56789-01234-5678",
        address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
        vat: "DK-12345678",
        starting_date: "11-11-2019Z05:12:34T",
        country_code: "DK",
        updated_at: new Date(),
        created_at: new Date(),
        name: "Demo Company ApS",
        status: "warning",
        status_notes: "",
        rank: 2,
        name_vector: "Demo Company ApS"
      },
      companies: [
        {
          id: "2345-56789-01234-5678",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
          vat: "DK-23456789",
          country: "Danmark / DK",
          type: "IVS",
          starting_date: "11-11-2019Z05:12:34T",
          country_code: "DK",
          updated_at: new Date(),
          created_at: new Date(),
          name: "Demo Company 2 ApS",
          status: "warning",
          status_notes: "",
          rank: 2,
          name_vector: "Demo Company ApS"
        },
        {
          id: "3456-56789-01234-5678",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
          country: "Danmark / DK",
          type: "IVS",
          vat: "DK-34567890",
          starting_date: "11-11-2019Z05:12:34T",
          country_code: "DK",
          updated_at: new Date(),
          created_at: new Date(),
          name: "Demo Company 3 ApS",
          status: "warning",
          status_notes: "",
          rank: 2,
          name_vector: "Demo Company ApS"
        },
        {
          id: "4567-56789-01234-5678",
          address: "Anker Engelundsvej 1, 2800 Kgs. Lyngby",
          country: "Danmark / DK",
          type: "IVS",
          vat: "DK-45678901",
          starting_date: "11-11-2019Z05:12:34T",
          country_code: "DK",
          updated_at: new Date(),
          created_at: new Date(),
          name: "Demo Company 4 ApS",
          status: "good",
          status_notes: "",
          rank: 2,
          name_vector: "Demo Company ApS"
        }
      ],
      people: {
        bad: [
          {
            id: "1234-67890-12345-6789",
            full_name: "Mallory Malicious",
            relation: "Ultimate Beneficial Owner",
            country: "Denmark / DK",
            ownership: 0.2,
            voting_rights: 0.3,
            updated_at: new Date(),
            created_at: new Date(),
            name_vector: "Mallory Malicious",
            type: "sanctions",
            source: "EU Sanctions List 2019",
            address: "https://www.example.org/link/sanctions-list.pdf"
          }
        ],
        warning: [
          {
            id: "2345-67890-12345-6789",
            full_name: "Phil Pepper",
            relation: "Ultimate Beneficial Owner",
            country: "Denmark / DK",
            ownership: 0.6,
            voting_rights: 0,
            updated_at: new Date(),
            created_at: new Date(),
            name_vector: "Phil Pepper",
            type: "pep",
            source: "UN PEP List 2019, Q4",
            address: "https://www.example.org/link/pep-list.pdf"
          }
        ],
        good: [
          {
            id: "3456-67890-12345-6789",
            full_name: "Gary Goodguy",
            relation: "Ultimate Beneficial Owner",
            country: "Denmark / DK",
            ownership: 20,
            voting_rights: 70,
            updated_at: new Date(),
            created_at: new Date(),
            name_vector: "Gary Goodguy",
            first_name: "Gary",
            last_name: "Goodguy",
            country_code: "DK"
          },
          {
            id: "4567-67890-12345-6789",
            full_name: "Gary Goodguy2",
            relation: "Ultimate Beneficial Owner",
            country: "Denmark / DK",
            ownership: 0.2,
            voting_rights: 0.3,
            updated_at: new Date(),
            created_at: new Date(),
            name_vector: "Gary Goodguy2",
            first_name: "Gary",
            last_name: "Goodguy2",
            country_code: "DK"
          }
        ]
      },
      errors: []
    };
  },
  signup: function (email: string) {
    return "";
  }
};



export interface Company {
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
}

export interface BadPerson extends CommonPersonFields {
  type: string;
  source?: string;
  address?: string;
}

export interface Person extends CommonPersonFields {
  first_name: string;
  last_name: string;
  country_code: string;
}
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
