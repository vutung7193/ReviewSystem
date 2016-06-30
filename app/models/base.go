package base

type Base interface {
	Insert ()
	Update ()
	Delete ()
	Find ()
	FindAll ()
	FindBy ()
	FindOne ()
}

type A struct {

}

type B struct {
	a A
}