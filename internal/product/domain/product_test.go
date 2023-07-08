package domain

import (
	"testing"
	"time"
)

func TestProduct_ToString(t *testing.T) {
	type fields struct {
		ProductId   string
		ProductName string
		ProductType ProductType
		Price       Currency
		Created     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				ProductId:   "1",
				ProductName: "Product 1",
				ProductType: ProductTypeEquity,
				Price: Currency{
					Amount:   100,
					Currency: "USD",
				},
				Created: time.Now(),
			},
			want: "{1,Equity,Product 1,100.00 USD}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				ProductId:   tt.fields.ProductId,
				ProductName: tt.fields.ProductName,
				Price:       tt.fields.Price,
				ProductType: tt.fields.ProductType,
				Created:     tt.fields.Created,
			}
			if got := p.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
