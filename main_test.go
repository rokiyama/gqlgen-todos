package main_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rokiyama/gqlgen-todos/server"
)

func TestServer(t *testing.T) {
	s := server.New()
	queryDir := path.Join("testdata", "query")
	wantDir := path.Join("testdata", "want")
	files, err := ioutil.ReadDir(queryDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			q, err := ioutil.ReadFile(path.Join(queryDir, file.Name()))
			if err != nil {
				t.Fatal(err)
			}
			b, err := json.Marshal(struct {
				Query string `json:"query"`
			}{
				Query: string(q),
			})
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewReader(b))
			req.Header.Add("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			s.ServeHTTP(rec, req)
			resp := rec.Result()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			var m map[string]interface{}
			if err := json.Unmarshal(body, &m); err != nil {
				t.Fatal(err)
			}
			body, err = json.MarshalIndent(m, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			wantFile := path.Join(wantDir, file.Name()+".json")
			want, err := ioutil.ReadFile(wantFile)
			if os.IsNotExist(err) {
				if err := os.MkdirAll(path.Join(wantDir), 0755); err != nil {
					t.Fatal(err)
				}
				if err := ioutil.WriteFile(wantFile, body, 0444); err != nil {
					t.Fatal(err)
				}
				return
			} else if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(string(want), string(body)); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
				t.Logf("If you get an error unexpectedly, You should re-run tests with `rm -r %s`", wantDir)
			}
		})
	}
}
