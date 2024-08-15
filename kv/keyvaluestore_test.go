package keyvaluestore

import (
	"testing"
)

func TestSet(t *testing.T) {
	keystr := "Gladys"
	valuestr := "Meyer"
	want := valuestr
	store := NewKeyValueStore()
	store.Set(keystr, valuestr)
	got := store.data[keystr]
	if got == want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSet2(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{ //test cases
		{ //test case
			name: "should return the correct value",
			args: args{"gunjan", "sinha"},
			want: "sinha",
		}, { //test case
			name: "should return the error ",
			args: args{"gunjan", "sinha"},
			want: "sinhar",
		},
	}
	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) {
			store := NewKeyValueStore()
			store.Set(tt.args.a, tt.args.b)
			got := store.data[tt.args.a]
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
