package accounts

import "bank/customer"

type CheckingAccount struct {
	HolderUser    customer.HolderUser
	BranchNumber  int
	AccountNumber int
	funds         float64
}

func (c *CheckingAccount) Withdraw(values ...float64) string {
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

func (c *CheckingAccount) Deposit(values ...float64) (string, float64) {
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

func (c *CheckingAccount) Transfer(value float64, destineAccount *CheckingAccount) bool {
	canNotTranfer := value > c.funds || value < 0
	if canNotTranfer {
		return false
	}

	c.funds -= value
	destineAccount.Deposit(value)
	return true
}

func (c *CheckingAccount) Funds() float64 {
	return c.funds
}
