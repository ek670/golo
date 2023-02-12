package golo

import (
	"testing"
)

// func TestMap(t *testing.T) {
// 	type args struct {
// 		input []I
// 		fn    func(elem I) O
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []O
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Map(tt.args.input, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Map() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestDelete(t *testing.T) {
	tt := struct {
		name string
		args struct {
			values *[]string
			index  int
		}
		want    []string
		wantErr bool
	}{
		name: "",
		args: struct {
			values *[]string
			index  int
		}{
			values: &[]string{"0", "1", "2"},
			index:  1,
		},
		want:    []string{"0", "2"},
		wantErr: false,
	}

	t.Run(tt.name, func(t *testing.T) {
		if err := Delete(tt.args.values, tt.args.index); (err != nil) != tt.wantErr {
			t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			t.Errorf("Delete() want = %v, result %v", tt.want, *(tt.args.values))
		}
	})
}
