package main

type Asset struct {
	Id string
	Name string
	Symbol string
	Updated string
	ValueInUsd string
	ValueInBtc string
}

type AssetService interface {
	Retrieve() ([]Asset, error)
}
