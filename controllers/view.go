package controllers

// ViewController operations for Admin
type ViewController struct {
	BaseController
}

// URLMapping ...
func (c *ViewController) URLMapping() {
	c.Mapping("Home", c.Main)
}

// Main ...
// @Title Main
// @router / [get]
func (c *ViewController) Main() {
	c.TplName = "index.html"
}
