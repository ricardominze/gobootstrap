package errord

import "errors"

var (
	CustomerErrorEmptyName = errors.New("não é possível criar o cliente")
)
