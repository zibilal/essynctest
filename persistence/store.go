package persistence

import (
	"bitbucket.org/kudoindonesia/microservice_order/helpers/uuid"
)

type Storer interface {
	Store(name string, data interface{}) error
}

type Updater interface {
	Update(id uuid.ID, name string, data interface{}) error
}

type Fetcher interface {
	Fetch(interface{}, string, interface{}) error
}

type Persistence interface {
	Storer
	Fetcher
}
