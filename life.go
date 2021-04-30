package life

type Lifer interface {
	Live()
	Die()
	Code()
	Sleep()
}

type Life struct {
}

//Newlife returns a pointer to life. Life cannot be copied, make sure no one alters your life,
//careful when sharing your life. Anyone can chaneg your life if you are not carefull.
//You get only one,
//TODO: should it return Lifer?
func Newlife() *Life {
	return &Life{}
}

//TODO: yet to be implemented, just born bro. Take a break before I hit your life with granade.
func (l *Life) Live()  {}
func (l *Life) Die()   {}
func (l *Life) Code()  {}
func (l *Life) Sleep() {}
