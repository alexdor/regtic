package interfaces

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/types"
)

type Response events.APIGatewayProxyResponse

type CompanyJson struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Address      null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Vat          string      `boil:"vat" json:"vat" toml:"vat" yaml:"vat"`
	StartingDate null.String `boil:"startingDate" json:"startingDate,omitempty" toml:"startingDate" yaml:"startingDate,omitempty"`
	CountryCode  string      `boil:"countryCode" json:"countryCode,omitempty" toml:"countryCode" yaml:"countryCode,omitempty"`
	Name         null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Status       null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	StatusNotes  null.String `boil:"statusNotes" json:"statusNotes,omitempty" toml:"statusNotes" yaml:"statusNotes,omitempty"`
	Type         null.String `json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
}

type CompaniesJson []CompanyJson

type Person struct {
	ID          string          `boil:"id" json:"id" toml:"id" yaml:"id"`
	CountryCode string          `boil:"countryCode" json:"countryCode" toml:"countryCode" yaml:"countryCode"`
	Name        null.String     `boil:"name,omitempty" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	CheckStatus CheckStatusEnum `json:"checkStatus,omitempty" toml:"checkStatus" yaml:"checkStatus,omitempty"`
	BadType     null.String     `json:"statusType,omitempty" toml:"statusType" yaml:"statusType,omitempty"`
	Source      null.String     `json:"source,omitempty" toml:"source" yaml:"source,omitempty"`
	OwnedBy     []Relationship  `json:"ownedBy,omitempty" toml:"ownedBy" yaml:"ownedBy,omitempty"`
	EntityType  EntityTypeEnum  `json:"entityType,omitempty" toml:"entityType" yaml:"entityType,omitempty"`
}
type People []Person
type Company struct {
	ID           string          `boil:"id" json:"id" toml:"id" yaml:"id"`
	Address      null.String     `boil:"address,omitempty" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	CountryCode  string          `boil:"countryCode" json:"countryCode" toml:"countryCode" yaml:"countryCode"`
	Name         null.String     `boil:"name,omitempty" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	CheckStatus  CheckStatusEnum `json:"checkStatus,omitempty" toml:"checkStatus" yaml:"checkStatus,omitempty"`
	BadType      null.String     `json:"statusType,omitempty" toml:"statusType" yaml:"statusType,omitempty"`
	Source       null.String     `json:"source,omitempty" toml:"source" yaml:"source,omitempty"`
	OwnedBy      []Relationship  `json:"ownedBy,omitempty" toml:"ownedBy" yaml:"ownedBy,omitempty"`
	EntityType   EntityTypeEnum  `json:"entityType,omitempty" toml:"entityType" yaml:"entityType,omitempty"`
	Vat          string          `json:"vat" toml:"vat" yaml:"vat"`
	StartingDate null.String     `json:"startingDate,omitempty" toml:"startingDate" yaml:"startingDate,omitempty"`
	Status       null.String     `json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	StatusNotes  null.String     `json:"statusNotes,omitempty" toml:"statusNotes" yaml:"statusNotes,omitempty"`
	Type         null.String     `json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
}
type Companies []Company
type ValidationResponse struct {
	Info      CompanyJson `json:"info,omitempty" toml:"info" yaml:"info,omitempty"`
	People    People      `json:"people,omitempty" toml:"people" yaml:"people,omitempty"`
	Companies Companies   `json:"companies,omitempty" toml:"companies" yaml:"companies,omitempty"`
	Errors    []error     `json:"errors"`
}

type Relationship struct {
	EntityID     string            `json:"id" toml:"id" yaml:"id"`
	Relation     types.StringArray `json:"relation" toml:"relation" yaml:"relation"`
	Ownership    float32           `json:"ownership" toml:"ownership" yaml:"ownership"`
	VotingRights float32           `json:"votingRights" toml:"votingRights" yaml:"votingRights"`
	EntityType   EntityTypeEnum    `json:"entityType" toml:"entityType" yaml:"entityType"`
}
