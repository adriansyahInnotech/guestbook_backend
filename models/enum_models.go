package models

type IDCardType string

const (
	IDCardKTP IDCardType = "KTP"
	IDCardSIM IDCardType = "SIM"
	IDCardPAS IDCardType = "PASSPORT"
)
