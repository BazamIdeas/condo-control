package controllers

// CondosController operations for Condos
type CondosController struct {
	BaseController
}

// URLMapping ...
func (c *CondosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *CondosController) Post() {

}

func (c *CondosController) GetOne() {

}

func (c *CondosController) GetAll() {

}

func (c *CondosController) Put() {

}

func (c *CondosController) Delete() {

}
