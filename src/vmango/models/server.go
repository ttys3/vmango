package models

type Server struct {
	Data map[string]interface{}
	Type string
}

type ServerList []*Server
