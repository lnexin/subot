package subscription

import (
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/decimal"
	"strings"
)

type SubscriptionPlan interface {
	Id() string
	Name() string
	Price() decimal.Decimal
}

type ClubSubscriptionPlan struct {
}

func (c ClubSubscriptionPlan) Id() string {
	return "id: " + uuid.New().String()
}

func (c ClubSubscriptionPlan) Name() string {
	return strings.Join([]string{"plan", c.Name()}, ":")
}

func (c ClubSubscriptionPlan) Price() decimal.Decimal {
	return c.Price()
}
