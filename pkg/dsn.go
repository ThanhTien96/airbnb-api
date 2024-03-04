package dsn

import "fmt"

type DSN struct {
	Scheme string
	Host   string
	Port   int
	User   string
	Pass   string
	DB     string
	Params string
}

func New(scheme string, host string, port int, user string, pass string, db string, params string) DSN {
	return DSN{
		Scheme: scheme,
		Host:   host,
		Port:   port,
		User:   user,
		Pass:   pass,
		DB:     db,
		Params: params,
	}
}

func (i DSN) String() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?%s", i.Scheme, i.User, i.Pass, i.Host, i.Port, i.DB, i.Params)
}
