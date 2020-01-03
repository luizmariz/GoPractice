package core

// SetupRouter is an important method of Core
// and set up all api routes
func (c *Core) SetupRouter() {
	c.Router.
		Methods("POST").
		Path("/api/mandrake").
		HandlerFunc(c.handleMandrakeSearch)

	c.Router.
		Methods("GET").
		Path("/api/mandrake").
		HandlerFunc(c.handleSearchesByUser)

	c.Router.
		Methods("GET").
		Path("/api/mandrake/download").
		HandlerFunc(c.handleCsvDownload)

	c.Router.
		Methods("POST").
		Path("/api/user").
		HandlerFunc(c.handleUser)
}
