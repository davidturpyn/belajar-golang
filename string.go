package main

import "fmt"

func main() {
	fmt.Println("David")
	fmt.Println("David Apricio")
	fmt.Println("David Apricio Turpyn")

	fmt.Println(len("David"))
	fmt.Println("David Apricio"[0])        // mengembalikan byte index ke 0 harus di convert kalau mau dapetin index ke 0 jadi string
	fmt.Println("David Apricio Turpyn"[1]) // mengembalikan byte index ke 1 harus di convert kalau mau dapetin index ke 1 jadi string
}
