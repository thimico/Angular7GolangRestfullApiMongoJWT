package idao

import (
	"../../entities"
)

type IDAO interface {
	Create(*entities.IEntity)
	FindAll()
}
