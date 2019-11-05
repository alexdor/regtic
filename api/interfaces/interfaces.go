package interfaces

import (
	"time"

	"github.com/alexdor/regtic/api/models"
	"github.com/volatiletech/null"
)

type PersonJson struct {
	models.Person
	Rank null.Float64 `boil:"rank" json:"rank,omitempty" toml:"rank" yaml:"rank"`
}

type BadPersonJson struct {
	models.BadPerson
	Rank null.Float64 `boil:"rank" json:"rank,omitempty" toml:"rank" yaml:"rank"`
}

type CompanyJson struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Address      null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Vat          string      `boil:"vat" json:"vat" toml:"vat" yaml:"vat"`
	StartingDate null.String `boil:"starting_date" json:"starting_date,omitempty" toml:"starting_date" yaml:"starting_date,omitempty"`
	CountryCode  null.String `boil:"country_code" json:"country_code,omitempty" toml:"country_code" yaml:"country_code,omitempty"`
	UpdatedAt    time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Name         null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Status       null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	StatusNotes  null.String `boil:"status_notes" json:"status_notes,omitempty" toml:"status_notes" yaml:"status_notes,omitempty"`
	//models.Company
	Rank null.Float64 `boil:"rank" json:"rank,omitempty" toml:"rank" yaml:"rank"`
}
