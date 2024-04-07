package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"marie": {
		AuthToken: "789GHI", 
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"jason": {
		Coins: 200,
		Username: "jason",
	},
	"marie": {
		Coins: 300, 
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

/* 
// generic function that combines the GetLoginDetails and GetUserCoins functions

func GetUserDetails[D LoginDetails | CoinDetails ](username string) *any {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = D{}
	clientData, ok := mockDetails[username]   // assuming mockCoin.. and mockLogin.. are combined
	if !ok {
		return nil
	}

	return &clientData
}
*/

func (d *mockDB) SetupDatabase() error {
	return nil
}

