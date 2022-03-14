package clients

import "os"

type Options struct {
	PostgreSQLConnString string `json:"postgresql_connstring"`
}

func (opts *Options) Init() {
	if opts.PostgreSQLConnString == "" {
		opts.PostgreSQLConnString = os.Getenv("POSTGRESQL_URL")
	}
}
