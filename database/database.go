package database

import (
	supa "github.com/nedpals/supabase-go"
)

type DatabaseInstance interface {
	GetDatabase() *supa.Client
}
