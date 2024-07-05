package test

import (
	"strings"
	"testing"

	rmmerge "github.com/mikeschinkel/readme-merge"
)

func TestReadme_Merge(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
		readme  rmmerge.Merger
	}{
		{
			name: "Recursive includes",
			want: "\n# Read Me\n\n## File 1\n\n### File 2\n\n#### File 3\n",
			readme: rmmerge.NewReadme("readme/_index.md", "\n# Read Me\n[merge](./file1.md)\n").
				AddChild(rmmerge.NewReadme("readme/file1.md", "\n# File 1\n[merge](./file2.md)\n")).
				AddChild(rmmerge.NewReadme("readme/file2.md", "\n# File 2\n[merge](./file3.md)\n")).
				AddChild(rmmerge.NewReadme("readme/file3.md", "\n# File 3\n")).
				Root(),
		},
		{
			name:   "Hashes in Code",
			want:   "\n# Read Me\n\nSome text   ```py\n   # This is a comment\n   ```\nSome more text\n",
			readme: rmmerge.NewReadme("readme/_index.md", "\n# Read Me\n\nSome text   ```py\n   # This is a comment\n   ```\nSome more text\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.readme.MergeWithLevel(0)
			if (err != nil) != tt.wantErr {
				t.Errorf("MergeWithLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MergeWithLevel():\n\tgot =  %v\n\twant = %v", collapseNewLines(got), collapseNewLines(tt.want))
			}
		})
	}
}

func collapseNewLines(s string) string {
	return strings.Replace(s, "\n", "\\n", -1)
}

//func TestNewReadme(t *testing.T) {
//	type args struct {
//		name    string
//		content string
//	}
//	tests := []struct {
//		name string
//		args args
//		want *rmmerge.Readme
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := rmmerge.NewReadme(tt.args.name, tt.args.content); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewReadme() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestReadme_AddChild(t *testing.T) {
//	type fields struct {
//		Name     string
//		reader   io.Reader
//		builder  *strings.Builder
//		children map[string]rmmerge.Merger
//	}
//	type args struct {
//		m rmmerge.Merger
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rm := &rmmerge.Readme{
//				Name:     tt.fields.Name,
//				reader:   tt.fields.reader,
//				builder:  tt.fields.builder,
//				children: tt.fields.children,
//			}
//			rm.AddChild(tt.args.m)
//		})
//	}
//}
//
//func TestReadme_Builder(t *testing.T) {
//	type fields struct {
//		Name     string
//		reader   io.Reader
//		builder  *strings.Builder
//		children map[string]rmmerge.Merger
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   *strings.Builder
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rm := &rmmerge.Readme{
//				Name:     tt.fields.Name,
//				reader:   tt.fields.reader,
//				builder:  tt.fields.builder,
//				children: tt.fields.children,
//			}
//			if got := rm.Builder(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Builder() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestReadme_CloseReader(t *testing.T) {
//	type fields struct {
//		Name     string
//		reader   io.Reader
//		builder  *strings.Builder
//		children map[string]rmmerge.Merger
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rm := &rmmerge.Readme{
//				Name:     tt.fields.Name,
//				reader:   tt.fields.reader,
//				builder:  tt.fields.builder,
//				children: tt.fields.children,
//			}
//			rm.CloseReader()
//		})
//	}
//}
//
//func TestReadme_GetChild(t *testing.T) {
//	type fields struct {
//		Name     string
//		reader   io.Reader
//		builder  *strings.Builder
//		children map[string]rmmerge.Merger
//	}
//	type args struct {
//		e string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantM   rmmerge.Merger
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rm := &rmmerge.Readme{
//				Name:     tt.fields.Name,
//				reader:   tt.fields.reader,
//				builder:  tt.fields.builder,
//				children: tt.fields.children,
//			}
//			gotM, err := rm.GetChild(tt.args.e)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetChild() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotM, tt.wantM) {
//				t.Errorf("GetChild() gotM = %v, want %v", gotM, tt.wantM)
//			}
//		})
//	}
//}
//
//func TestReadme_Reader(t *testing.T) {
//	type fields struct {
//		Name     string
//		reader   io.Reader
//		builder  *strings.Builder
//		children map[string]rmmerge.Merger
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    io.Reader
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rm := &rmmerge.Readme{
//				Name:     tt.fields.Name,
//				reader:   tt.fields.reader,
//				builder:  tt.fields.builder,
//				children: tt.fields.children,
//			}
//			got, err := rm.Reader()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Reader() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Reader() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
