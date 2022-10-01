package main

import (
	"fmt"
	"unsafe"

	"github.com/jackc/pgtype"
)

func main() {
	fmt.Println("pgtype.ArrayType{}", unsafe.Sizeof(pgtype.ArrayType{}))
	fmt.Println("pgtype.Date{}", unsafe.Sizeof(pgtype.Date{}))
	fmt.Println("pgtype.Daterange{}", unsafe.Sizeof(pgtype.Daterange{}))
	fmt.Println("pgtype.Int4range{}", unsafe.Sizeof(pgtype.Int4range{}))
	fmt.Println("pgtype.Numeric{}", unsafe.Sizeof(pgtype.Numeric{}))
	fmt.Println("pgtype.Numrange{}", unsafe.Sizeof(pgtype.Numrange{}))
}
