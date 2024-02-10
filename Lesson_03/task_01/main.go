package main

import (
	"fmt"
)

type Order struct {
	quantity     int
	price        int
	customerName string
	productName  string
}

func (o Order) printOrder() string {

	return fmt.Sprintf("Заказ от %s: %dx %s (Итого: $%d)\n", o.customerName, o.quantity, o.productName, o.getTotalPrice())

}

func (o Order) getTotalPrice() int {
	return o.quantity * o.price
}

func main() {

	order := Order{12, 300, "Иван", "Книга"}

	fmt.Println(order.printOrder())
}
