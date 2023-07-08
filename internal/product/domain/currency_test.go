package domain

import "testing"

func TestCurrency_String(t *testing.T) {
	type fields struct {
		Amount   float64
		Currency string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				Amount:   100,
				Currency: "USD",
			},
			want: "100.00 USD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Currency{
				Amount:   tt.fields.Amount,
				Currency: tt.fields.Currency,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
