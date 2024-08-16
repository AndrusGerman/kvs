package keyvaluestore

import (
	"testing"
)

func TestSet(t *testing.T) {
	store := NewKeyValueStore()
	store.Set("Gladys", "Meyer")
	store.Set("Harry", "Bosch")
	store.Set("John", "Sanford")
	store.Set("Michael", "Connelly")
	want := "Sanford"
	got := store.data["John"]
	if got != want {
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
			name: "incorrect value",
			args: args{"gunjan", "sinha"},
			want: "sinhar",
		}, {
			name: "change value",
			args: args{"gunjan", "sinhar"},
			want: "sinhar",
		}, {
			name: "incorrect value - 2",
			args: args{"gunjan", "sinhar"},
			want: "sinha",
		},
	}
	store := NewKeyValueStore()
	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) {
			store.Set(tt.args.a, tt.args.b)
			got := store.data[tt.args.a]
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestSet3(t *testing.T) {
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
			name: "incorrect value",
			args: args{"gunjan", "sinha"},
			want: "sinhar",
		}, {
			name: "change value",
			args: args{"gunjan", "sinhar"},
			want: "sinhar",
		}, {
			name: "incorrect value - 2",
			args: args{"gunjan", "sinhar"},
			want: "sinha",
		},
	}
	store := NewKeyValueStore()
	for _, tt := range tests { // test cases executions
		t.Run(tt.name, func(t *testing.T) {
			store.Set(tt.args.a, tt.args.b)
			got := store.data[tt.args.a]
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
