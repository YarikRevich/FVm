package exceptions

type IException interface{
	GetCode()int
}

type Exception struct {
	code int
}

func (e *Exception) GetCode()int{
	return 0;
};
