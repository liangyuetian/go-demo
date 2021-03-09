package pre_order

import (
	"fmt"
	"github.com/kataras/iris/v12"
	Mysql "study/src/mysql"
	"sync"
	"time"
)

type Controller struct {
	/* dependencies */
}

func URL() string {
	return "pre-order"
}

type PreOrderRow struct {
	id          int `db:"id"`
	status      int `db:"status"`
	createdTime int `db:"created_time"`
}

func (c *Controller) Get(ctx iris.Context) string {
	DB := Mysql.GetDB()
	var info = new(Mysql.UserFarmerLogEntity)
	row := DB.QueryRow("select * from xznz_user_farmer_log where id = ?", 4)
	err := row.Scan(&info.Id, &info.Status, &info.CreatedTime, &info.UpdatedTime, &info.Version, &info.CustomerId, &info.CustomerName, &info.Role, &info.SellerCid, &info.SellerUserName, &info.ShareToken, &info.SupplyId, &info.SupplyName, &info.PageSource)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(ctx)
	return ""
}

func (c *Controller) Post(b PreOrderRow) string {
	fmt.Println(b)
	return "porder controller post"
}

func (c *Controller) Put() string {
	return "porder controller put"
}

func w() {
	var s sync.WaitGroup
	s.Add(2)
	go func() {
		//.....
		time.Sleep(1 * time.Second)
		s.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		s.Done()
	}()
	s.Wait()
}

func Get(ctx iris.Context) {
	DB := Mysql.GetDB()
	var info Mysql.UserFarmerLogEntity
	row := DB.QueryRow("select * from xznz_user_farmer_log where id = ?", 4)
	err := row.Scan(&info.Id, &info.Status, &info.CreatedTime, &info.UpdatedTime, &info.Version, &info.CustomerId, &info.CustomerName, &info.Role, &info.SellerCid, &info.SellerUserName, &info.ShareToken, &info.SupplyId, &info.SupplyName, &info.PageSource)

	if err != nil {
		fmt.Println(err)
	}
	w()
	fmt.Println(info)
	ctx.JSON(info)
}

func Post(ctx iris.Context) {

}
