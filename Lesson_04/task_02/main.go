package main

import (
	"errors"
	"fmt"
	"time"
)

type Operator interface {
	Balance() int64
	Withdraw(amount int64) error
	Deposit(amount int64) error
	Transactions() []Tx
}

var ErrNegativeValue = errors.New("negative value")

var _ error = (*withdrawError)(nil)

type withdrawError struct{}

func (b withdrawError) Error() string {
	// TODO implement me
	//panic("implement me")
	return fmt.Sprintf("Ошибка списания %s", b.Error())
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
	return fmt.Sprintf("Сумма: %s%0.1f, time: %s", t.action, t.value, t.createdAt.Format(time.DateTime))
}

var _ Operator = (*txOperator)(nil)

type txOperator struct {
	transactions []Tx
}

func (t *txOperator) Balance() int64 {
	var total int64
	for _, tx := range t.transactions {
		switch tx.action {
		case ActionKindDecr:
			total -= tx.value
		case ActionKindIncr:
			total += tx.value
		default:
		}
	}

	return total
}

func (t *txOperator) Withdraw(amount int64) error {
	balance := t.Balance()
	if balance-amount < 0 {
		return errors.New("insufficient funds")
	}
	t.transactions = append(
		t.transactions, Tx{
			value:     amount,
			action:    ActionKindDecr,
			createdAt: time.Now(),
		},
	)
	return nil
}

func (t *txOperator) Deposit(amount int64) error {
	if amount < 0 {
		return errors.New("negative value")
	}

	t.transactions = append(
		t.transactions, Tx{
			value:     amount,
			action:    ActionKindIncr,
			createdAt: time.Now(),
		},
	)

	return nil
}

func (t *txOperator) Transactions() []Tx {
	return t.transactions
}

func main() {
	var op Operator = &txOperator{}
	_ = op
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", errors.As(withdrawError{}))
	}
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Deposit(100); err != nil {
		fmt.Println("Отрицательное значение", err)
	}
	fmt.Println("Текущий баланс", op.Balance())
	if err := op.Withdraw(100); err != nil {
		fmt.Println("Ошибка списания", err)
	}
	fmt.Println("Текущий баланс", op.Balance())
}
