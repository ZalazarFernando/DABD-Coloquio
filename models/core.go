package models

import (
	"time"

	"gorm.io/gorm"
)

type AdministradorServicios struct {
	gorm.Model
	DNI        int
	Email      string
	Contrasena string
}

type Servicio struct {
	gorm.Model
	Disponibilidad bool      `gorm:"column:Disponibilidad; not null"`
	FechaPartida   time.Time `gorm:"column:Fecha_Partida; not null"`
	FechaLlegada   time.Time `gorm:"column:Fecha_Llegada; not null"`
	CostoServicio  float64   `gorm:"column:Costo_Servicio; not null"`
}

// TableName overrides the table name used by User to `profiles`
func (Servicio) TableName() string {
	return "ViajaPlus.dbo.Servicio"
}

type CalidadServicio struct {
	ID         int `gorm:"primaryKey"`
	IDServicio int
	Calidad    string
	Servicio   Servicio `gorm:"foreignKey:IDServicio"`
}

type Transporte struct {
	gorm.Model
	NroUnidad       int
	Pisos           int
	Situacion       bool
	CostoTransporte float64
}

// TableName overrides the table name used by User to `profiles`
func (Transporte) TableName() string {
	return "ViajaPlus.dbo.Transporte"
}

type Asiento struct {
	ID             int `gorm:"primaryKey"`
	IDTransporte   int
	Disponibilidad bool
	Transporte     Transporte `gorm:"foreignKey:IDTransporte"`
}

type CategoriaTransporte struct {
	ID           int `gorm:"primaryKey"`
	IDTransporte int
	Categoria    string
	Transporte   Transporte `gorm:"foreignKey:IDTransporte"`
}

type TipoAtencionTransporte struct {
	ID           int `gorm:"primaryKey"`
	IDTransporte int
	Tipo         string
	Transporte   Transporte `gorm:"foreignKey:IDTransporte"`
}

type Tramo struct {
	gorm.Model
	Distancia    int       `gorm:"column:Distancia; not null"`
	FechaPartida time.Time `gorm:"column:Fecha_Partida; not null"`
	FechaLlegada time.Time `gorm:"column:Fecha_Llegada; not null"`
	CostoTramo   float64   `gorm:"column:Costo_Tramo; not null"`
}

func (Tramo) TableName() string {
	return "ViajaPlus.dbo.Tramo"
}

type Itinerario struct {
	gorm.Model
	Distancia int `gorm:"column:Distancia; not null"`
}

type Ciudad struct {
	gorm.Model
	Nombre string
}

func (Ciudad) TableName() string {
	return "ViajaPlus.dbo.Ciudad"
}

type Pasaje struct {
	gorm.Model
	IDServicio int
	Costo      float64
	Servicio   Servicio `gorm:"foreignKey:IDServicio"`
}

type Reserva struct {
	gorm.Model
	Nombre   string
	Apellido string
	DNI      int
}

type EstadoReserva struct {
	ID        int `gorm:"primaryKey"`
	IDReserva int
	Estado    string
	Reserva   Reserva `gorm:"foreignKey:IDReserva"`
}

type ServicioItinerario struct {
	IDServicio   int        `gorm:"primaryKey"`
	IDItinerario int        `gorm:"primaryKey"`
	Servicio     Servicio   `gorm:"foreignKey:IDServicio"`
	Itinerario   Itinerario `gorm:"foreignKey:IDItinerario"`
}

type ItinerarioTramo struct {
	IDTramo      int        `gorm:"primaryKey"`
	IDItinerario int        `gorm:"primaryKey"`
	Tramo        Tramo      `gorm:"foreignKey:IDTramo"`
	Itinerario   Itinerario `gorm:"foreignKey:IDItinerario"`
	Orden        int        `gorm:"column:Orden"`
}

type ServicioTransporte struct {
	IDServicio   int        `gorm:"primaryKey"`
	IDTransporte int        `gorm:"primaryKey"`
	Servicio     Servicio   `gorm:"foreignKey:IDServicio"`
	Transporte   Transporte `gorm:"foreignKey:IDTransporte"`
}

type TramoCiudad struct {
	IDTramo  int    `gorm:"primaryKey"`
	IDCiudad int    `gorm:"primaryKey"`
	Tramo    Tramo  `gorm:"foreignKey:IDTramo"`
	Ciudad   Ciudad `gorm:"foreignKey:IDCiudad"`
}

type ReservaCiudad struct {
	IDReserva int `gorm:"primaryKey"`
	IDCiudad  int `gorm:"primaryKey"`
	EsOrigen  bool
	Reserva   Reserva `gorm:"foreignKey:IDReserva"`
	Ciudad    Ciudad  `gorm:"foreignKey:IDCiudad"`
}

type AdministradorTransporte struct {
	IDAdministrador int                    `gorm:"primaryKey"`
	IDTransporte    int                    `gorm:"primaryKey"`
	Administrador   AdministradorServicios `gorm:"foreignKey:IDAdministrador"`
	Transporte      Transporte             `gorm:"foreignKey:IDTransporte"`
}

type AdministradorServicio struct {
	IDAdministrador int                    `gorm:"primaryKey"`
	IDServicio      int                    `gorm:"primaryKey"`
	Administrador   AdministradorServicios `gorm:"foreignKey:IDAdministrador"`
	Servicio        Servicio               `gorm:"foreignKey:IDServicio"`
}

type AdministradorItinerario struct {
	IDAdministrador int                    `gorm:"primaryKey"`
	IDItinerario    int                    `gorm:"primaryKey"`
	Administrador   AdministradorServicios `gorm:"foreignKey:IDAdministrador"`
	Itinerario      Itinerario             `gorm:"foreignKey:IDItinerario"`
}

type AdministradorTramo struct {
	IDAdministrador int                    `gorm:"primaryKey"`
	IDTramo         int                    `gorm:"primaryKey"`
	Administrador   AdministradorServicios `gorm:"foreignKey:IDAdministrador"`
	Tramo           Tramo                  `gorm:"foreignKey:IDTramo"`
}

type AdministradorCiudad struct {
	IDAdministrador int                    `gorm:"primaryKey"`
	IDCiudad        int                    `gorm:"primaryKey"`
	Administrador   AdministradorServicios `gorm:"foreignKey:IDAdministrador"`
	Ciudad          Ciudad                 `gorm:"foreignKey:IDCiudad"`
}
