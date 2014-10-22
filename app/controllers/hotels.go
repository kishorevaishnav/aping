package controllers

import (
	"log"

	"github.com/revel/revel"
)

type Hotels struct {
	GorpController
}
type Hotel struct {
	HotelId          int
	Name, Address    string
	City, State, Zip string
	Country          string
	Price            int
}

func (c Hotels) Index() revel.Result {
	// var list []Hotel
	count, _ := Dbm.SelectInt("select count(*) from Hotel")
	log.Println("Row count - should be zero:", count)
	// h, _ := c.Txn.Select(models.Hotel{}, "SELECT * FROM hotels")
	// fmt.Println(list[0])
	return c.Render(count)
}

func (c Hotels) List() revel.Result {
	return c.Render()
}
