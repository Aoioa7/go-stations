package model

type ErrNotFound struct{
	Message string
}

func (e *ErrNotFound) Error()string{
	return "error not found"
}