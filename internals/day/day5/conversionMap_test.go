package day5

import "testing"

func TestConversionMap_convert(t *testing.T) {
	type fields struct {
		destinationStart int
		sourceStart      int
		size             int
		sourceEnd        int
	}
	type args struct {
		seed int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"convert test 1", fields{50, 98, 2, 99}, args{99}, 51, false},
		{"convert test 2", fields{52, 50, 48, 97}, args{50}, 52, false},
		{"convert test 3", fields{52, 50, 48, 97}, args{30}, 0, true},
		{"convert test 4", fields{50, 98, 2, 99}, args{100}, 0, true},
		{"convert test 5", fields{50, 50, 2, 51}, args{51}, 51, false},
		{"convert test 6", fields{50, 50, 2, 51}, args{50}, 50, false},
		{"convert test 7", fields{50, 50, 2, 51}, args{52}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConversionMap{
				DestinationStart: tt.fields.destinationStart,
				SourceStart:      tt.fields.sourceStart,
				Size:             tt.fields.size,
				SourceEnd:        tt.fields.sourceEnd,
			}
			got, err := c.Convert(tt.args.seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConversionMap_inRange(t *testing.T) {
	type fields struct {
		destinationStart int
		sourceStart      int
		size             int
		sourceEnd        int
	}
	type args struct {
		seed int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"inRange test 1", fields{50, 98, 2, 99}, args{50}, false},
		{"inRange test 2", fields{50, 98, 2, 99}, args{98}, true},
		{"inRange test 3", fields{50, 98, 2, 99}, args{99}, true},
		{"inRange test 4", fields{50, 98, 2, 99}, args{100}, false},
		{"inRange test 5", fields{50, 50, 2, 51}, args{52}, false},
		{"inRange test 6", fields{50, 50, 2, 51}, args{51}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConversionMap{
				DestinationStart: tt.fields.destinationStart,
				SourceStart:      tt.fields.sourceStart,
				Size:             tt.fields.size,
				SourceEnd:        tt.fields.sourceEnd,
			}
			if got := c.InRange(tt.args.seed); got != tt.want {
				t.Errorf("inRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
