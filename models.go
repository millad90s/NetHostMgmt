package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Server ...
type Server struct {
	gorm.Model
	NameID  string `gorm:"unique;not null"`
	RAM     string
	CPU     string
	OwnerID uint `gorm:"not null"`
}

//Owner ...
type Owner struct {
	gorm.Model
	Name     string
	Email    string
	Username string `gorm:"unique;not null"`
	Password string
	Servers  []Server
}

//connectdb ...
func connectdb() (*gorm.DB, error) {
	ConnStr := db_user + ":" + db_pass + "@tcp(" + db_addr + ")/" + db_dbname + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Printf("%s", ConnStr)

	db, err := gorm.Open("mysql", ConnStr)
	if err != nil {
		fmt.Printf("%v", db)
		return gorm.Open("mysql", ConnStr)
	}
	return gorm.Open("mysql", ConnStr)
}

//ListMyServers ...
func (o Owner) ListMyServers() string {
	db, err := connectdb()

	if err != nil {
		panic(err.Error())
	}
	// db.AutoMigrate(&Server{})
	// defer db.Close()
	// if !db.HasTable(&Server{}) {
	// 	db.CreateTable(&Server{})
	// }
	// server := Server{NameID: s.NameID}
	println("lets search the servers owned to mine")
	db.Where(&o).Find(&o)
	if o.ID == 0 {
		// println("The owner dose not exist ...")
		return ""
	}
	fmt.Printf("%v", o.Servers)
	// return o.servers
	return ""
}

//ListAllUsers ...
func ListAllUsers() string {
	db, err := connectdb()
	if err != nil {
		panic(err.Error())
	}
	// db.AutoMigrate(&Server{})
	// defer db.Close()
	// if !db.HasTable(&Server{}) {
	// 	db.CreateTable(&Server{})
	// }
	// server := Server{NameID: s.NameID}
	println("lets search the servers owned to mine")
	// Get all records
	o := Owner{}
	rows, _ := db.Model(&Owner{}).Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &o)
		fmt.Printf("%v \n", o.Name)
	}
	fmt.Printf("%v", o.Servers)
	// return o.servers
	return ""
}

//AddServerToOwner ...
func (s Server) AddServerToOwner(o Owner) error {
	db, err := connectdb()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Server{})
	defer db.Close()
	if !db.HasTable(&Server{}) {
		db.CreateTable(&Server{})
	}
	// server := Server{NameID: s.NameID}
	println("lets find an server")
	db.Where(&s).Find(&s)
	db.Where(&o).Find(&o)

	if s.ID == 0 {
		println("no server fonded")
		return fmt.Errorf("no server found")
	}

	fmt.Printf("%v: \n ", o)

	// Update single attribute if it is changed
	db.Model(&s).Update("owner_id", o.ID)

	return nil
}
func (s Server) addnewtodb() error {
	db, err := connectdb()
	fmt.Printf("%v", db)
	if err != nil {
		panic(err.Error())
	}
	// db.AutoMigrate(&Server{})
	defer db.Close()
	if !db.Debug().HasTable(&Server{}) {
		db.Debug().CreateTable(&Server{})
	}
	d := db.Debug().Create(&s)
	// fmt.Printf("%v", d.Error)
	return d.Error

}

//Searchbyname ...
func (s Server) Searchbyname() (*Server, error) {
	db, err := connectdb()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	server := Server{NameID: s.NameID}
	db.Where(&server).Debug().Find(&server)
	if server.ID == 0 {
		return &Server{}, fmt.Errorf("no Item matched to keyword")
	}
	return &server, nil
}

func (o Owner) addnewOwner() error {
	db, err := connectdb()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(Owner{})

	defer db.Close()
	if !db.Debug().HasTable(&Server{}) {
		db.Debug().CreateTable(&Server{})
	}
	if !db.HasTable(&Owner{}) {
		db.CreateTable(&Owner{})
	}
	d := db.Create(&o)
	// fmt.Printf("%v", d.Error)
	return d.Error

}
