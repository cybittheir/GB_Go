package main

/*
В этой задаче мы практикуемся использовать интерфейсы и возвращать ошибки,а также повторяем все предыдущие знания о
слайсах. Нам необходимо реализовать интерфейс оператора. У него
есть набор методов, которые нам нужно реализовать в своей структуре txOperator.
Но мы несколько усложняем задачу. Мы будем создавать транзакцию на каждое изменение баланса. Необходимо в методе
Withdraw и Deposit создавать транзакцию, которая либо прибавляет значение либо отнимает.
Как считать баланс. Нужно пройтись по всем транзакциям и сложить или вычесть соответствующее значение.
В методе Withdraw нужно сначала вызвать функцию Balance() а потом проверить Balance() - current value > 0  и если
нет вывести ошибку.
Метод Transactions должен выводить список созданных транзакций.
Задание нужно предварительно разобрать со студентами, чтобы было понятно, что нужно сделать. Наш оператор должен
хранить список транзакций. Он должен проверять допустимость использования Withdraw, чтобы баланс при отнимание был
больше 0
Подсказка 1: Для вывода даты в строку используйте методов Format(time.DateTime)
Подсказка 2: Для форматирования числа с плавающей точкой испольуйте %0.1f в функции fmt.Sprintf()
Подсказка 3: Если мы передаем в метод Withdraw и Deposit отрицательные значения проверьте это и выдайте ошибку если
отрицательное, 0 допустимое значение
*/

import (
	"fmt"
	"time"
)

type Operator interface {
	Balance() int64
	Withdraw(amount int64) error
	Deposit(amount int64) error
	Transactions() []Tx
}

type ActionKind string

const (
	ActionKindIncr ActionKind = "+"
	ActionKindDecr ActionKind = "-"
)

type Tx struct {
	value     int64      // значение на которое изменилось
	action    ActionKind // действие, прибавляем или отнимаем
	createdAt time.Time
}

// Нужно вывести данные транзакции в формате сумма: +-value, time: время создания транзакции
func (t Tx) Print() string {
	return fmt.Sprintf("%s %d %v", t.action, t.value, t.createdAt.GoString())
}

var _ Operator = (*txOperator)(nil)

type txOperator struct{}

func (t *txOperator) Balance() int64 {
	// TODO implement me
	// panic("implement me")
	return t.Balance()
}

func (t *txOperator) Withdraw(amount int64) error {
	// TODO implement me
	panic("implement me")
}

func (t *txOperator) Deposit(amount int64) error {
	// TODO implement me
	panic("implement me")
}

func (t *txOperator) Transactions() []Tx {
	// TODO implement me
	panic("implement me")
}

func main() {
	var op Operator = &txOperator{}
	_ = op
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", err)
	}
	fmt.Println("Текущий баланс", op.Balance())
	op.Deposit(100)
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", err)
	}
	fmt.Println("Текущий баланс", op.Balance())
}
