package main

//
//import "fmt"
//
//type Observer interface {
//	Update(string)
//}
//
//type GuitarType struct {
//	model string
//	price int
//}
//
//type GuitarSellingShop struct {
//	name   string
//	guitar GuitarType
//}
//
//func (gss GuitarSellingShop) Update(msg string) {
//	fmt.Printf("New Guitar Is Out %s and the price is %d, %s\n", gss.guitar.model, gss.guitar.price, msg)
//}
//
//type RegisterShop struct {
//	observers []Observer
//}
//
//func (rs *RegisterShop) Attach(observer Observer) {
//	rs.observers = append(rs.observers, observer)
//
//
//// notify
//func (rs *RegisterShop) NotifyAviability() {
//	for _, observer := range rs.observers {
//		observer.Update("The new model is coming")
//	}
//}
//
////func main() {
////	Fender := GuitarSellingShop{
////		name:   "WE ARE GUITAR",
////		guitar: GuitarType{model: "Fender Player Stratocaster", price: 125},
////	}
////	Gibson := GuitarSellingShop{
////		name:   "WE ARE GUITAR 2",
////		guitar: GuitarType{model: "Gibson Les Paul Studio", price: 135},
////	}
////	Yamaha := GuitarSellingShop{
////		name:   "WE ARE GUITAR 3",
////		guitar: GuitarType{model: "Yamaha Pacifica 112V", price: 155},
////	}
////	RegisterShop := RegisterShop{}
////
////	RegisterShop.Attach(Fender)
////	RegisterShop.Attach(Gibson)
////	RegisterShop.Attach(Yamaha)
////
////	RegisterShop.NotifyAviability()
////}
