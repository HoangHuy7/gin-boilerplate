package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"monorepo/apps/gas/graph/model"
	"monorepo/shares/entities/mekyra_db"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
	"github.com/vektah/gqlparser/v2/ast"
)

// Decimal scalar marshaler/unmarshaler for gqlgen
// These methods are called by generated.go

func (ec *executionContext) unmarshalInputDecimal(ctx context.Context, v any) (decimal.Decimal, error) {
	return UnmarshalDecimal(v)
}

func (ec *executionContext) _Decimal(ctx context.Context, sel ast.SelectionSet, v *decimal.Decimal) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return MarshalDecimal(*v)
}

// MarshalDecimal marshals decimal.Decimal to GraphQL Decimal scalar
func MarshalDecimal(d decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, d.String())
	})
}

// UnmarshalDecimal unmarshals GraphQL Decimal scalar to decimal.Decimal
func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	switch v := v.(type) {
	case string:
		return decimal.NewFromString(v)
	case float64:
		return decimal.NewFromFloat(v), nil
	case int:
		return decimal.NewFromInt(int64(v)), nil
	case int64:
		return decimal.NewFromInt(v), nil
	case decimal.Decimal:
		return v, nil
	case json.Number:
		return decimal.NewFromString(v.String())
	default:
		fmt.Printf("❌ Invalid decimal type: %T, value: %+v\n", v, v)
		return decimal.Zero, fmt.Errorf("%T is not a valid Decimal", v)
	}
}

func (r *Resolver) mapDebtToModel(d *mekyra_db.Mkrtb_DebtTransaction) *model.DebtTransaction {
	var orderID *string
	if d.OrderId != nil {
		oid := d.OrderId.String()
		orderID = &oid
	}
	return &model.DebtTransaction{
		ID:         d.Id.String(),
		CustomerID: d.CustomerId.String(),
		OrderID:    orderID,
		Amount:     d.Amount,
		Type:       d.Type,
		Note:       &d.Note,
		CreatedAt:  &d.CreatedAt,
	}
}
