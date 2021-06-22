package groups

import "testing"

type Z4 struct{}

func (z *Z4) Operate(a, b Element) Element {
	bValue, _ := b.(int) // proper handle if type assertion fails
	aValue, _ := a.(int)
	return (aValue + bValue) % 4
}

func (z *Z4) Equals(a, b Element) bool {
	aValue, aOk := a.(int)
	bValue, bOk := b.(int)
	return aOk && bOk && aValue == bValue
}

func TestGroup_isClosed(t *testing.T) {
	tests := []struct {
		name    string
		g       *Group
		wantErr bool
	}{
		{
			name: "closed valid group",
			g: &Group{
				op:       &Z4{},
				elements: map[Element]bool{0: true, 1: true, 2: true, 3: true},
				table:    newCayleyTable(),
			},
			wantErr: false,
		},
		{
			name: "not closed valid group",
			g: &Group{
				op:       &Z4{},
				elements: map[Element]bool{1: true},
				table:    newCayleyTable(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.isClosed(); (err != nil) != tt.wantErr {
				t.Errorf("Group.isClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGroup_findIdentity(t *testing.T) {
	tests := []struct {
		name             string
		g                *Group
		wantErr          bool
		expectedIdentity int
	}{
		{
			name: "group with identity",
			g: &Group{
				op:       &Z4{},
				elements: map[Element]bool{0: true, 1: true, 2: true, 3: true},
				table:    newCayleyTable(),
			},
			wantErr:          false,
			expectedIdentity: 0,
		},
		{
			name: "group without identity",
			g: &Group{
				op:       &Z4{},
				elements: map[Element]bool{1: true, 2: true, 3: true},
				table:    newCayleyTable(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var err error
			var identity Element
			if identity, err = tt.g.findIdentity(); (err != nil) != tt.wantErr {
				t.Errorf("Group.isClosed() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if value, ok := identity.(int); !ok || value != tt.expectedIdentity {
				t.Errorf("Found unexpected identity")
			}

		})
	}
}

func TestGroup_hasInverses(t *testing.T) {
	tests := []struct {
		name             string
		g                *Group
		wantErr          bool
		expectedIdentity int
	}{
		{
			name: "group with inverses",
			g: &Group{
				op:       &Z4{},
				elements: map[Element]bool{0: true, 1: true, 2: true, 3: true},
				table:    newCayleyTable(),
			},
			wantErr: false,
		},
		{
			name: "group without inverses",
			g: &Group{
				op:       &Z4{},
				elements: map[Element]bool{0: true, 2: true, 3: true},
				table:    newCayleyTable(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.g.hasInverses(); (err != nil) != tt.wantErr {
				t.Errorf("Group.isClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
