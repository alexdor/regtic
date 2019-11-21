import axios from "axios";

const api = axios.create({
  baseURL: process.env.API_URL || "http://localhost:3030/",
  timeout: 200000,
  headers: {
    "Content-Type": "application/json"
  }
});

export default {
  findCompanies(searchStr: string) {
    return api
      .get(`v1/find_companies?name=${searchStr}`)
      .then((res: { data: FindCompaniesData }) => res.data.companies);
  },
  validateCompany(id: string) {
    return api
      .get(`v1/validate_company?id=${id}`)
      .then((res: { data: CheckCompanyData }) => res.data);
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
