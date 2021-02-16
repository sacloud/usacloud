package query

import "testing"

func TestByGoJQ(t *testing.T) {
	type args struct {
		input   interface{}
		query   string
		printer func(interface{}) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "with map",
			args: args{
				input: map[string]interface{}{
					"foo": "bar",
				},
				query:   ".",
				printer: func(interface{}) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "with slice",
			args: args{
				input:   []interface{}{"1", "2"},
				query:   ".",
				printer: func(interface{}) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "with struct",
			args: args{
				input:   struct{ Foo string }{Foo: "bar"},
				query:   ".",
				printer: func(interface{}) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "with primitive",
			args: args{
				input:   1,
				query:   ".",
				printer: func(interface{}) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "with invalid query",
			args: args{
				input:   map[string]interface{}{"foo": "bar"},
				query:   "...",
				printer: func(interface{}) error { return nil },
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ByGoJQ(tt.args.input, tt.args.query, tt.args.printer); (err != nil) != tt.wantErr {
				t.Errorf("ByGoJQ() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
