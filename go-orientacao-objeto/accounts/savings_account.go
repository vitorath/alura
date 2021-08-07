package accounts

import "bank/customer"

type SavingsAccount struct {
	HolderUser                             customer.HolderUser
	BranchNumber, AccountNumber, Operation int
	funds                                  float64
}

func (c *SavingsAccount) Withdraw(values ...float64) string {
	defaultFunds := c.funds
	for _, value := range values {
		canWithdraw := value > 0 && value <= c.funds
		if canWithdraw {
			c.funds -= value
		} else {
			c.funds = defaultFunds
			return "Insufficient funds"
		}
	}

	return "Withdraw with successfuly"
}

func (c *SavingsAccount) Deposit(values ...float64) (string, float64) {
	defaultFunds := c.funds
	for _, value := range values {
		if value > 0 {
			c.funds += value
		} else {
			c.funds = defaultFunds
			return "Deposit failure", c.funds
		}
	}

	return "Deposit with successfully", c.funds
}

func (c *SavingsAccount) Funds() float64 {
	return c.funds
}
