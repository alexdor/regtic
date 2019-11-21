package interfaces

import (
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/volatiletech/null"
)

type Response events.APIGatewayProxyResponse

type PersonJson struct {
	ID          string       `boil:"id" json:"id" toml:"id" yaml:"id"`
	FirstName   null.String  `boil:"first_name" json:"first_name,omitempty" toml:"first_name" yaml:"first_name,omitempty"`
	LastName    null.String  `boil:"last_name" json:"last_name,omitempty" toml:"last_name" yaml:"last_name,omitempty"`
	CountryCode null.String  `boil:"country_code" json:"country_code,omitempty" toml:"country_code" yaml:"country_code,omitempty"`
	UpdatedAt   time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt   time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Rank        null.Float64 `boil:"rank" json:"rank,omitempty" toml:"rank" yaml:"rank"`
}

type BadPersonJson struct {
	ID        string       `boil:"id" json:"id" toml:"id" yaml:"id"`
	FullName  null.String  `boil:"full_name" json:"full_name,omitempty" toml:"full_name" yaml:"full_name,omitempty"`
	Type      null.String  `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	Source    null.String  `boil:"source" json:"source,omitempty" toml:"source" yaml:"source,omitempty"`
	Address   null.String  `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	UpdatedAt time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Rank      null.Float64 `boil:"rank" json:"rank,omitempty" toml:"rank" yaml:"rank"`
}

type CompanyJson struct {
	ID           string       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Address      null.String  `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Vat          string       `boil:"vat" json:"vat" toml:"vat" yaml:"vat"`
	StartingDate null.String  `boil:"starting_date" json:"starting_date,omitempty" toml:"starting_date" yaml:"starting_date,omitempty"`
	CountryCode  null.String  `boil:"country_code" json:"country_code,omitempty" toml:"country_code" yaml:"country_code,omitempty"`
	UpdatedAt    time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt    time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Name         null.String  `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Status       null.String  `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	StatusNotes  null.String  `boil:"status_notes" json:"status_notes,omitempty" toml:"status_notes" yaml:"status_notes,omitempty"`
	Rank         null.Float64 `boil:"rank" json:"rank,omitempty" toml:"rank" yaml:"rank"`
}
type CompaniesJson []CompanyJson
