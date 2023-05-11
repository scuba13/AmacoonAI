package model

import (
	"time"


	"gorm.io/gorm"
)

type Cat struct {
	gorm.Model
	Name             string                 `gorm:"column:name"`
	Registration     string                 `gorm:"column:registration"`
	RegistrationType string                 `gorm:"column:registration_type"`
	Microchip        string                 `gorm:"column:microchip"`
	Gender           string                 `gorm:"column:gender"`
	Birthdate        time.Time              `gorm:"column:birthdate"`
	Neutered         bool                   `gorm:"column:neutered"`
	Validated        bool                   `gorm:"column:validated"`
	Observation      string                 `gorm:"column:observation"`
	Fifecat          bool                   `gorm:"column:fifecat"`
	MotherID         *uint                  `gorm:"column:mother_id"`
	MotherName       string                 `gorm:"-"`
	FatherID         *uint                  `gorm:"column:father_id"`
	FatherName       string                 `gorm:"-"`
	FederationID     *uint                  `gorm:"column:federation_id"`
	Federation       *Federation `gorm:"foreignKey:FederationID"`
	BreedID          *uint                  `gorm:"column:breed_id"`
	Breed            *Breed           `gorm:"foreignKey:BreedID"`
	ColorID          *uint                  `gorm:"column:color_id"`
	Color            *Color           `gorm:"foreignKey:ColorID"`
	CatteryID        *uint                  `gorm:"column:cattery_id"`
	Cattery          *Cattery       `gorm:"foreignKey:CatteryID"`
	OwnerID          *uint                  `gorm:"column:owner_id"`
	Owner            *Owner           `gorm:"foreignKey:OwnerID"`
	CountryID        *uint                  `gorm:"column:country_id"`
	Country          *Country       `gorm:"foreignKey:CountryID"`
	Titles           []TitlesCat            `gorm:"foreignKey:CatID"`
	FatherNameTemp   string                 
	MotherNameTemp   string
	//Files            []utils.Files
}

func (Cat) TableName() string {
	return "cats"
}

type TitlesCat struct {
	gorm.Model
	CatID        uint
	TitleID      uint
	Titles       *Title `gorm:"foreignkey:TitleID"`
	Date         time.Time
	FederationID uint `gorm:"foreignkey:FederationID"`
}

func (TitlesCat) TableName() string {
	return "titles_cats"
}


type Breed struct {
	gorm.Model
	BreedCode     string `gorm:"type:varchar(191);unique"`
	BreedName     string
	BreedCategory int
	BreedByGroup  string
}

func (Breed) TableName() string {
	return "breeds"
}

type Cattery struct {
	gorm.Model
	Name        string
	BreederName string
	OwnerID          *uint                  `gorm:"foreignKey:OwnerID"`
	CountryID        *uint                  `gorm:"foreignKey:CountryID"`
}

func (Cattery) TableName() string {
    return "catteries"
}

type Color struct {
    gorm.Model
    BreedCode string
    EmsCode   string
    Name      string
    Group     int
    SubGroup  int
}

func (Color) TableName() string {
    return "colors"
}

type Country struct {
    gorm.Model
    Code        string
    Name        string
    IsActivated bool
}

func (Country) TableName() string {
    return "countries"
}

type Federation struct {
	gorm.Model
	Name           string          
	FederationCode string          
	CountryID        *uint                  `gorm:"foreignKey:CountryID"`    
	
}

func (Federation) TableName() string {
	return "federations"
}

type Owner struct {
	gorm.Model
	Email        string
	PasswordHash string
	Name         string
	CPF          string
	Address      string
	City         string
	State        string
	ZipCode      string
	CountryID    *uint 
	Country      *Country `gorm:"foreignKey:CountryID"`
	Phone        string
	Valid        bool
	ValidId      string
	Observation  string
}

func (Owner) TableName() string {
	return "owners"
}

type Title struct {
    gorm.Model
    Name        string
	Code        string `gorm:"type:varchar(191);unique"`
    Type        string
    Certificate string
    Amount      int
    Observation string
}

func (Title) TableName() string {
    return "titles"
}