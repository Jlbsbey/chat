package main

/*
	func main() {
		var db DB

}
*/

type DefinedAction interface {
	GetFromJSON([]byte)
	Process( /*dB *DB*/ )
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
