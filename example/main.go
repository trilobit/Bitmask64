package main

import (
	"fmt"
	"strings"

	"github.com/trilobit/Bitmask64/pkg/bitmask"
)

const (
	AccessCreate = iota
	AccessRead
	AccessUpdate
	AccessDelete
)

type User struct {
	accessLevel bitmask.Bitmask64
	// other fields for defining a user
}

func main() {
	admin := User{
		// default access level – restrict access to all
		accessLevel: bitmask.NewBitmask64(0),
	}

	checkPermissions(admin.accessLevel)

	// allow user to read
	err := admin.accessLevel.SetBit(AccessRead)
	if err != nil {
		panic(err)
	}

	checkPermissions(admin.accessLevel)
}

func checkPermissions(mask bitmask.Bitmask64) {
	fmt.Println("Check user for all access levels:")
	fmt.Println("Create: ", mask.IsSetBit(AccessCreate))
	fmt.Println("  Read: ", mask.IsSetBit(AccessRead))
	fmt.Println("Update: ", mask.IsSetBit(AccessUpdate))
	fmt.Println("Delete: ", mask.IsSetBit(AccessDelete))
	fmt.Println(strings.Repeat("-", 35))
}
