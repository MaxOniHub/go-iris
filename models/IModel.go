package models

type IModel interface
{
	Validate(data interface{}) (bool, interface{})
}
