package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// srv1 := Server{NameID: "tx01", CPU: "2"}
	// srv2 := Server{NameID: "tx02", CPU: "4"}
	// srv3 := Server{NameID: "tx03", CPU: "8"}
	// srv1.addnewtodb()
	// srv2.addnewtodb()
	// srv3.addnewtodb()

	// owner1 := Owner{Username: "owner1"}
	// owner2 := Owner{Username: "owner2"}
	// owner1.addnewOwner()
	// owner2.addnewOwner()

	// fmt.Printf("%v 000000000000000", srv3)
	// fmt.Printf("%v", owner1.Servers)

	// srv1.AddServerToOwner(owner2)
	// srv2.AddServerToOwner(owner1)

	ListAllUsers()

	// result, e := Server{NameID: "tx03"}.Searchbyname()
	// if e != nil {
	// 	fmt.Printf("%s", e)
	// 	return
	// }

	// fmt.Printf("%v", result)
	// fmt.Printf("%v", result)

	// create new server
	// srv := Server{
	// 	NameID: "tx03",
	// }
	// err := srv.addnewtodb()
	// if err != nil {
	// 	println(err)
	// }

	//create new owner
	// owner2 := Owner{
	// 	Username: "owner5",
	// }
	// err := srv.AddServerToOwner(owner1)
	// if err != nil {
	// 	println(err)
	// }
	// fmt.Printf("%v", owner1.ID)

}
