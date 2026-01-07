package seed

import "github.com/pathway/backend/models"

func courseHTTP() models.Course {
	return models.Course{
		Title:       "HTTP Networking",
		Description: "Deep dive into HTTP protocols, REST APIs, request/response cycles, and modern web communication patterns.",
		Modules: []models.Module{
			moduleHTTP1(),
			moduleHTTP2(),
			moduleHTTP3(),
			moduleHTTP4(),
			moduleHTTP5(),
			moduleHTTP6(),
			moduleHTTP7(),
		},
	}
}
