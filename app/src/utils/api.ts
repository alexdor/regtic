import axios from "axios";
import { Notification } from "element-ui";
import { FindCompaniesData, ValidationResponse } from "./interfaces";
import mockapi from "./mockapi";
const UseMockApi = process.env.VUE_APP_USE_MOCK_API;
const api = axios.create({
  baseURL: process.env.VUE_APP_API_URL || "http://localhost:8090/",
  timeout: 200000,
  headers: {
    "Content-Type": "application/json",
  },
});

api.interceptors.response.use(undefined, (error) => {
  const isNetworkError = !(error.response || {}).status;
  const errorTitle = isNetworkError ? "Network error" : "There was an error 😅";
  const errorMessage = isNetworkError
    ? "There was an unexpected network error, please verify that you are connected to the internet and refresh the page"
    : ((error.response || {}).data || {}).error || error.message;

  Notification.error({
    duration: 4000,
    title: errorTitle,
    message: errorMessage,
    position: "bottom-right",
  });

  return Promise.reject(error);
});

export default Object.freeze({
  apiVersion: "v2",
  findCompanies(searchStr: string) {
    return UseMockApi
      ? mockapi.findCompanies(searchStr)
      : api
          .get(`${this.apiVersion}/find_companies?name=${searchStr}`)
          .then(
            ({ data }: { data: FindCompaniesData }) => data && data.companies
          );
  },
  validateCompany(id: string) {
    return UseMockApi
      ? mockapi.validateCompany(id)
      : api
          .get(`${this.apiVersion}/validate_company?id=${id}`)
          .then(({ data }: { data: ValidationResponse }) => data);
  },

  signup(email: string) {
    return UseMockApi
      ? mockapi.signup(email)
      : api.get(`${this.apiVersion}/signup?email=${email}`);
  },
  getPerson: mockapi.getPerson,
});
