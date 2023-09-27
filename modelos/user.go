package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (usuario *User) AddUser(id int, name string, createdAdd time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAdd
	usuario.Status = status
}
