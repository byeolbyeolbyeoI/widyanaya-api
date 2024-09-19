package database

import (
	"fmt"
	"github.com/byeolbyeolbyeoI/icom/config"
	"github.com/supabase-community/supabase-go"
)

type Database struct {
	client *supabase.Client
}

func NewDatabase(conf *config.Config) DatabaseInstance {
	client, err := supabase.NewClient(conf.Database.URL, conf.Database.Key, nil)
	if err != nil {
		fmt.Println("cannot initialize client", err)
		return nil
	}

	return &Database{client: client}
}

func (d *Database) GetDatabase() *supabase.Client {
	return d.client
}
