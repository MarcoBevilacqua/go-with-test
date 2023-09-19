package main

type Transaction struct {
	From, To string
	Sum      float64
}

type Account struct {
	Name    string
	Balance float64
}

func Sum(numbers []int) int {

	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}

	return result
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, t Transaction) Account {
	if t.From == a.Name {
		a.Balance -= t.Sum
	}

	if t.To == a.Name {
		a.Balance += t.Sum
	}

	return a
}
