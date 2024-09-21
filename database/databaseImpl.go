package database

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/config"
	supa "github.com/nedpals/supabase-go"
)

type Database struct {
	client *supa.Client
}

func NewDatabase(conf *config.Config) DatabaseInstance {
	client := supa.CreateClient(conf.Database.URL, conf.Database.Key)

	return &Database{client: client}
}

func (d *Database) GetDatabase() *supa.Client {
	return d.client
}
