package entity

import srv "monografia/service"

type Entity struct {
	service srv.Service
}

func New(service srv.Service) *Entity {
	return &Entity{
		service: service,
	}
}
