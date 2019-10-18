package routes

func GetRoutes() Routes {
	return Routes{
		Route{
			Name:    "SendRequest",
			Method:  "GET",
			Path:    "/send",
			Handler: Send(),
		},
	}

}
