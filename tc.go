package client

func (c *Client) TCs() ([]TC, error) {
	var tcs []TC
	return tcs, c.get("/api/tcs", &tcs)
}

func (c *Client) TC(id string) (TC, error) {
	var tc TC
	return tc, c.get("/api/tcs/"+id, &tc)
}

func (c *Client) CreateTC(o TC) error {
	return c.put("/api/tcs", &o)
}

func (c *Client) DeleteTC(id string) error {
	return c.delete("/api/tcs/" + id)
}

func (c *Client) UpdateTC(id string, o TC) error {
	return c.post("/api/tcs/"+id, &o)
}

func (c *Client) TCUsage(id string) (StatsResponse, error) {
	var s StatsResponse
	return s, c.get("/api/tcs/"+id+"/usage", &s)
}

func (c *Client) TempSensors() ([]string, error) {
	var sensors []string
	return sensors, c.get("/api/tcs/sensors", &sensors)
}
