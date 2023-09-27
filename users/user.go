package users

import (
	"Github/Ch1nolas/Curso-Go/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Chinolas", time.Now(), true)
	fmt.Println(u)
}
