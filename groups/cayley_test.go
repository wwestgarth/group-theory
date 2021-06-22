package groups

import "testing"

func Test_cayleyTable_add(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		c    *cayleyTable
		args args
	}{
		{
			name: "add and lookup",
			c:    newCayleyTable(),
			args: args{
				a: 1,
				b: 2,
				c: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.add(tt.args.a, tt.args.b, tt.args.c)
			res, err := tt.c.lookup(tt.args.a, tt.args.b)

			if err != nil {
				t.Errorf("CayleyTable.lookup() error = %v", err)
				return
			}

			if value, ok := res.(int); !ok || value != tt.args.c {
				t.Errorf("CayleyTable.lookup() value = %v", value)
				return
			}

		})
	}
}
