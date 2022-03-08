package modul4_Context

import (
	"context"
	"fmt"
	"testing"
)

func TestWithValue(t *testing.T) {
	ctxA := context.Background()

	ctxB := context.WithValue(ctxA, 2, "B")
	ctxC := context.WithValue(ctxA, 3, "C")

	ctxD := context.WithValue(ctxB, 4, "D")
	ctxE := context.WithValue(ctxB, 5, "E")

	ctxF := context.WithValue(ctxC, 6, "F")
	ctxG := context.WithValue(ctxE, 6, "G")

	fmt.Println("A.", ctxA)
	fmt.Println("B.", ctxB)
	fmt.Println("C.", ctxC)
	fmt.Println("D.", ctxD)
	fmt.Println("E.", ctxE)
	fmt.Println("F.", ctxF)
	fmt.Println("G.", ctxG)

	fmt.Println("------------------")
	fmt.Println(ctxF.Value(6))
	fmt.Println(ctxG.Value(2)) //akan terus mencari data dg key tersebut sampai ke parent yg paling tinggi
	fmt.Println(ctxG.Value(1)) //return nil jika tidak ada
	fmt.Println(ctxE.Value(6)) //tidak bisa mengambil data dari child
}
