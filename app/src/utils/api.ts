const axios = require("axios");

const api = axios.create({
  baseURL: "http://localhost:3030/",
  timeout: 200000,
  headers: {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*"
  }
});

export default {
  async findCompanies(searchStr: string) {
    return api
      .get(`v1/find_companies?name=${searchStr}`)
      .then((res: any) => res.data.companies);
  },
  async validateCompany(id: string) {
    return api.get(`v1/validate_company?id=${id}`).then((res: any) => res.data);
  }
};
