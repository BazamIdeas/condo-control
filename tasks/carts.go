package tasks

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/vjeantet/jodaTime"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

// the toolbox package
func init() {
	CartsTask := toolbox.NewTask("CartsTask", "0 0 0 25 * *", func() error {

		limitDate := time.Now().AddDate(0, 0, -30)
		date := jodaTime.Format("YYYY-MM-dd HH:mm:ss", limitDate)

		//beego.Debug("Exect Carts tasks")

		//TODO: LOG

		o := orm.NewOrm()

		_, err := o.QueryTable("carts").Filter("created_at__lt", date).Delete()

		if err != nil {
			beego.Debug(err.Error())
			return err
		}

		//beego.Debug("Limit Date: ", date)
		//beego.Debug("Deleted Carts: ", num)

		return nil
	})

	toolbox.AddTask("CartsTask", CartsTask)
	toolbox.StartTask()
	defer toolbox.StopTask()
}
