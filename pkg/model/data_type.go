package model

type DataType int

const (
	DataUnknown DataType = iota
	DataMedical
	DataNonMedical
	DataClarifying
)
