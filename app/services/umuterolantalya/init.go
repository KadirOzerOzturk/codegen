package umuterolantalya

import (
	"github.com/KadirOzerOzturk/deneme/internal/database"
	"github.com/KadirOzerOzturk/deneme/internal/service_collection"
	"github.com/KadirOzerOzturk/deneme/pkg/di"
)

// Inject dependencies
func init() {
	di.AddSingleton[Repository](service_collection.Collection(), func(s *di.Scope) any {
		return NewRepo(database.Connection())
	})
	di.AddScoped[Service](service_collection.Collection(), func(s *di.Scope) any {
		repo, _ := di.GetService[Repository](s)
		return NewService(repo)
	})
}