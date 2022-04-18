package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_extractDate(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     time.Time
		wantErr  bool
	}{
		{
			name:     "not valid #1",
			filename: "Bild 1.jpg",
			want:     time.Time{},
			wantErr:  true,
		},
		{
			name:     "YYYYMMDD date #1",
			filename: "corr_s-Ge_X_20220404_1_Bild 1.jpg",
			want:     time.Date(2022, 4, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "YYYYMMDD date #2",
			filename: "20201022_X_Deutsch.pdf",
			want:     time.Date(2020, 10, 22, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "YYMMDD date #1",
			filename: "200415_STEST-X_Anlage_Deutsch Anlage.pdf",
			want:     time.Date(2020, 4, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "YYMMDD date #2",
			filename: "091124_OEB-X_Brief.pdf",
			want:     time.Date(2009, 11, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "YYMMDD date #3",
			filename: "140609_FOO-X_Deutsch.pdf",
			want:     time.Date(2014, 6, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "YYYYMM date #1",
			filename: "S-123_s-Ge_X_202109_Deutsch.pdf",
			want:     time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "YY.MM date #1",
			filename: "S-123-20.08-X_Ge_Deutsch.pdf",
			want:     time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractDate(tt.filename)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil, "error not as expected: %v", err)
		})
	}
}
