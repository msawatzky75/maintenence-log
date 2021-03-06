// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type DateFilter struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type DistanceValue struct {
	Value *float64      `json:"value"`
	Type  *DistanceUnit `json:"type"`
}

type DistanceValueInput struct {
	Value float64      `json:"value"`
	Type  DistanceUnit `json:"type"`
}

type FluidValue struct {
	Value *float64   `json:"value"`
	Type  *FluidUnit `json:"type"`
}

type FluidValueInput struct {
	Value float64   `json:"value"`
	Type  FluidUnit `json:"type"`
}

type FuelEfficiency struct {
	Kml   float64 `json:"kml"`
	L100k float64 `json:"l100k"`
	Mpg   float64 `json:"mpg"`
}

type FuelLogInput struct {
	Date       time.Time           `json:"date"`
	VehicleID  string              `json:"vehicleId"`
	Odometer   *DistanceValueInput `json:"odometer"`
	Notes      *string             `json:"notes"`
	Trip       *DistanceValueInput `json:"trip"`
	Grade      int                 `json:"grade"`
	FuelAmount *FluidValueInput    `json:"fuelAmount"`
	FuelPrice  *MoneyValueInput    `json:"fuelPrice"`
}

type LogsFilter struct {
	Date   *DateFilter `json:"date"`
	Recent *int        `json:"recent"`
}

type MaintenanceLogInput struct {
	Date      time.Time           `json:"date"`
	VehicleID string              `json:"vehicleId"`
	Odometer  *DistanceValueInput `json:"odometer"`
	Notes     string              `json:"notes"`
}

type MoneyValue struct {
	Value *float64   `json:"value"`
	Type  *MoneyUnit `json:"type"`
}

type MoneyValueInput struct {
	Value float64   `json:"value"`
	Type  MoneyUnit `json:"type"`
}

type OilChangeLogInput struct {
	Date      time.Time           `json:"date"`
	VehicleID string              `json:"vehicleId"`
	Odometer  *DistanceValueInput `json:"odometer"`
	Notes     *string             `json:"notes"`
}

type UserInput struct {
	Email string `json:"email"`
}

type UserPreferenceInput struct {
	Distance  *DistanceUnit `json:"distance"`
	Fluid     *FluidUnit    `json:"fluid"`
	Money     *MoneyUnit    `json:"money"`
	VehicleID *string       `json:"vehicleId"`
}

type VehicleInput struct {
	Make     string              `json:"make"`
	Model    string              `json:"model"`
	Year     int                 `json:"year"`
	Odometer *DistanceValueInput `json:"odometer"`
}

type DistanceUnit string

const (
	DistanceUnitKilometre DistanceUnit = "Kilometre"
	DistanceUnitMile      DistanceUnit = "Mile"
)

var AllDistanceUnit = []DistanceUnit{
	DistanceUnitKilometre,
	DistanceUnitMile,
}

func (e DistanceUnit) IsValid() bool {
	switch e {
	case DistanceUnitKilometre, DistanceUnitMile:
		return true
	}
	return false
}

func (e DistanceUnit) String() string {
	return string(e)
}

func (e *DistanceUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DistanceUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DistanceUnit", str)
	}
	return nil
}

func (e DistanceUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FluidUnit string

const (
	FluidUnitLitre  FluidUnit = "Litre"
	FluidUnitGallon FluidUnit = "Gallon"
)

var AllFluidUnit = []FluidUnit{
	FluidUnitLitre,
	FluidUnitGallon,
}

func (e FluidUnit) IsValid() bool {
	switch e {
	case FluidUnitLitre, FluidUnitGallon:
		return true
	}
	return false
}

func (e FluidUnit) String() string {
	return string(e)
}

func (e *FluidUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FluidUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FluidUnit", str)
	}
	return nil
}

func (e FluidUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MoneyUnit string

const (
	MoneyUnitCad MoneyUnit = "CAD"
	MoneyUnitUsd MoneyUnit = "USD"
)

var AllMoneyUnit = []MoneyUnit{
	MoneyUnitCad,
	MoneyUnitUsd,
}

func (e MoneyUnit) IsValid() bool {
	switch e {
	case MoneyUnitCad, MoneyUnitUsd:
		return true
	}
	return false
}

func (e MoneyUnit) String() string {
	return string(e)
}

func (e *MoneyUnit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MoneyUnit(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MoneyUnit", str)
	}
	return nil
}

func (e MoneyUnit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
