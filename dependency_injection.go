package pagination

import "go.uber.org/dig"

func SetupMongoDbPaginationProvider(container *dig.Container) {
	container.Provide(NewMongoDbPaginationProvider)
}
