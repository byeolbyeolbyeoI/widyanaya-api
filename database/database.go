package database

import "github.com/supabase-community/supabase-go"

type DatabaseInstance interface {
	GetDatabase() *supabase.Client
}
