package main

type DefinedAction interface {
	GetFromJSON([]byte)
	Process()
}

type GeneralObject interface {
	Create() DefinedAction
	Edit() DefinedAction
	Delete() DefinedAction
	Read() DefinedAction
	Login() DefinedAction
	Print()
	GetID() string
}
