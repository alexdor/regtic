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
	Name        null.String     `boil:"name" json:"name" toml:"name" yaml:"name"`
	CheckStatus CheckStatusEnum `json:"checkStatus" toml:"checkStatus" yaml:"checkStatus"`
	BadType     null.String     `json:"statusType" toml:"statusType" yaml:"statusType"`
	Source      null.String     `json:"source" toml:"source" yaml:"source"`
	OwnedBy     []Relationship  `json:"ownedBy" toml:"ownedBy" yaml:"ownedBy"`
	EntityType  EntityTypeEnum  `json:"entityType" toml:"entityType" yaml:"entityType"`
}
type People []Person
type Company struct {
	ID           string          `boil:"id" json:"id" toml:"id" yaml:"id"`
	Address      null.String     `boil:"address" json:"address" toml:"address" yaml:"address"`
	CountryCode  string          `boil:"countryCode" json:"countryCode" toml:"countryCode" yaml:"countryCode"`
	Name         null.String     `boil:"name" json:"name" toml:"name" yaml:"name"`
	CheckStatus  CheckStatusEnum `json:"checkStatus" toml:"checkStatus" yaml:"checkStatus"`
	BadType      null.String     `json:"statusType" toml:"statusType" yaml:"statusType"`
	Source       null.String     `json:"source" toml:"source" yaml:"source"`
	OwnedBy      []Relationship  `json:"ownedBy" toml:"ownedBy" yaml:"ownedBy"`
	EntityType   EntityTypeEnum  `json:"entityType" toml:"entityType" yaml:"entityType"`
	Vat          string          `json:"vat" toml:"vat" yaml:"vat"`
	StartingDate null.String     `json:"startingDate" toml:"startingDate" yaml:"startingDate"`
	Status       null.String     `json:"status" toml:"status" yaml:"status"`
	StatusNotes  null.String     `json:"statusNotes" toml:"statusNotes" yaml:"statusNotes"`
	Type         null.String     `json:"type" toml:"type" yaml:"type"`
}
type Companies []Company
type ValidationResponse struct {
	Info      CompanyJson `json:"info" toml:"info" yaml:"info"`
	People    People      `json:"people" toml:"people" yaml:"people"`
	Companies Companies   `json:"companies" toml:"companies" yaml:"companies"`
	Errors    []error     `json:"errors"`
}

type Relationship struct {
	EntityID     string            `json:"id" toml:"id" yaml:"id"`
	Relation     types.StringArray `json:"relation" toml:"relation" yaml:"relation"`
	Ownership    float32           `json:"ownership" toml:"ownership" yaml:"ownership"`
	VotingRights float32           `json:"votingRights" toml:"votingRights" yaml:"votingRights"`
	EntityType   EntityTypeEnum    `json:"entityType" toml:"entityType" yaml:"entityType"`
}
