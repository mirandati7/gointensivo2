package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_Gets_An_Error_IF_ID_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_If_Gets_An_Error_IF_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 10.00}
	assert.Error(t, order.Validate(), "invalid id")
}
