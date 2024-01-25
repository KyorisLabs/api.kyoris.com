package trainerDomain

import "time"

type TrainerModel struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"notNull"`
	Plan int    `gorm:"default:0"`

	// Meta
	CreateAt time.Time `gorm:"autoCreateTime"`
}

type TrainerUserModel struct {
	Id        uint      `gorm:"primaryKey;autoIncrement"`
	TrainerID int       `gorm:"notNull"`
	UserID    int       `gorm:"notNull"`
	RolID     int       `gorm:"notNull"`
	CreateAt  time.Time `gorm:"autoCreateTime"`
}

type TrainerTax struct {
	TrainerID    string `gorm:"primaryKey;autoIncrement"`
	TaxNumber    uint   `gorm:"notNull"` // Identificación fiscal
	NameTax      string `gorm:"notNull"` // Nombre fisca
	NameCommerce string // Nombre comercio

	// Domicilio (Direccion de facturación)
	Address  string `gorm:"notNull"` // Dirección de domicilio
	Zip      string `gorm:"notNull"` // Código postal, es un string ya que hay paises que contienen letras.
	Village  string `gorm:"notNull"` // Población
	Province string `gorm:"notNull"` // Provincia
	Country  string `gorm:"notNull"` // Pais

	// Contacto
	Phone  string // Teléfono
	Mobile string // Móvil
	Fax    string
	Mail   string // Correo electrónico

	// Meta
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
