package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{
		ID: "1",
	}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{
		ID:    "1",
		Price: 100,
	}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{ID: "2", Price: 100, Tax: 5}

	assert.Equal(t, "2", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)

	assert.Nil(t, order.IsValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, order.IsValid())
	assert.Nil(t, err)
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenShouldReturnFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, order.IsValid())
	assert.Nil(t, err)
	err = order.CalculateFinalPrice()
	assert.Nil(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)

}
