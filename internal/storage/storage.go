/*
	1. реализовать postgres хранилище
	2. реализовать миграции
*/

package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)
