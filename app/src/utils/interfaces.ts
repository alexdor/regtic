type EntityType = "PERSON" | "COMPANY";
type CheckStatusType = "OK" | "WARNING" | "ISSUE"; //Should we capitalize this?
type BadType = "PEP" | "SANCTION";
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
  status?: string | null;
  statusNotes?: string | null;
  type?: string;
}
export interface CommonFields {
  id: string;
  countryCode: string;
  name: string;
  checkStatus: CheckStatusType;
  statusType?: BadType | null;
  source?: string | null;
  ownedBy?: Relationship[] | null;
  entityType: EntityType;
}
export type CompanyJson = Omit<
  Company,
  "ownedBy" | "checkStatus" | "entityType"
>;
export type Person = CommonFields;
export interface ValidationResponse {
  info: CompanyJson;
  people: Person[];
  companies: Company[];
  errors?: any[];
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
  statusNotes?: string;
}

export interface ApiResponsePersonInfo extends Person {
  owns: OwnedCompany[];
  detectedLists: DetectedList[];
}

export interface Errors {
  msg: string;
}

export interface FindCompaniesData {
  companies?: Company[];
}
