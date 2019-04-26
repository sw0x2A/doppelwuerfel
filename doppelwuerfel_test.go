package doppelwuerfel

import (
	"reflect"
	"testing"
)

func Test_encrypt(t *testing.T) {
	type args struct {
		message []byte
		psk     []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "NOTEBOOK",
			args: args{
				message: []byte("HALLODASHIERISTEINLANGERBEISPIELTEXTUMDASVERFAHRENZUZEIGEN"),
				psk:     []byte("NOTEBOOK"),
			},
			want: []byte("OINPUFZLRASTRUSERLARGHHIBTSEEAINEEVNNDSGIMAEATEEDHILELIXEZ"),
		},
		{
			name: "DECKEL",
			args: args{
				message: []byte("OINPUFZLRASTRUSERLARGHHIBTSEEAINEEVNNDSGIMAEATEEDHILELIXEZ"),
				psk:     []byte("DECKEL"),
			},
			want: []byte("NRSGSESAIEOZRABINADIILURTNDEHXUSRHEVIEEPAEHEEGTLZFTLIANMEL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encrypt(tt.args.message, tt.args.psk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encrypt() = %v (%s), want %v (%s)", got, got, tt.want, tt.want)
			}
		})
	}
}

func Test_decrypt(t *testing.T) {
	type args struct {
		message []byte
		psk     []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "NOTEBOOK",
			args: args{
				message: []byte("OINPUFZLRASTRUSERLARGHHIBTSEEAINEEVNNDSGIMAEATEEDHILELIXEZ"),
				psk:     []byte("NOTEBOOK"),
			},
			want: []byte("HALLODASHIERISTEINLANGERBEISPIELTEXTUMDASVERFAHRENZUZEIGEN"),
		},
		{
			name: "DECKEL",
			args: args{
				message: []byte("NRSGSESAIEOZRABINADIILURTNDEHXUSRHEVIEEPAEHEEGTLZFTLIANMEL"),
				psk:     []byte("DECKEL"),
			},
			want: []byte("OINPUFZLRASTRUSERLARGHHIBTSEEAINEEVNNDSGIMAEATEEDHILELIXEZ"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.message, tt.args.psk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoppelwürfel_Encrypt(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		r    *Doppelwürfel
		args args
		want []byte
	}{
		{
			name: "NOTEBOOK/DECKEL",
			r:    NewDoppelwürfel([]byte("NOTEBOOK"), []byte("DECKEL")),
			args: args{
				text: []byte("HALLODASHIERISTEINLANGERBEISPIELTEXTUMDASVERFAHRENZUZEIGEN"),
			},
			want: []byte("NRSGSESAIEOZRABINADIILURTNDEHXUSRHEVIEEPAEHEEGTLZFTLIANMEL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Encrypt(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Doppelwürfel.Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoppelwürfel_Decrypt(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		r    *Doppelwürfel
		args args
		want []byte
	}{
		{
			name: "NOTEBOOK/DECKEL",
			r:    NewDoppelwürfel([]byte("NOTEBOOK"), []byte("DECKEL")),
			args: args{
				text: []byte("NRSGSESAIEOZRABINADIILURTNDEHXUSRHEVIEEPAEHEEGTLZFTLIANMEL"),
			},
			want: []byte("HALLODASHIERISTEINLANGERBEISPIELTEXTUMDASVERFAHRENZUZEIGEN"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Decrypt(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Doppelwürfel.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permutation(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "NOTEBOOK",
			args: args{
				s: []byte("NOTEBOOK"),
			},
			want: []int{4, 3, 7, 0, 1, 5, 6, 2},
		},
		{
			name: "DECKEL",
			args: args{
				s: []byte("DECKEL"),
			},
			want: []int{2, 0, 1, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
