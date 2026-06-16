package libs

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	openapi "github.intra.douban.com/sa/archery-cli"
)

func newTestRuntime(t *testing.T, serverURL string, token string) *Runtime {
	t.Helper()

	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{{URL: serverURL}}
	client := openapi.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, map[string]openapi.APIKey{
		"knoxApiToken": {Key: token},
	})

	return &Runtime{
		Config: Config{
			BaseURL:  serverURL,
			APIToken: token,
		},
		Client:  client,
		Context: ctx,
	}
}

func TestLookupTableInstances(t *testing.T) {
	const token = "token-lookup"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/instance/table-instances/" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != token {
			t.Fatalf("unexpected auth header: %s", got)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		var payload map[string]any
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		if payload["table_name"] != "users" {
			t.Fatalf("unexpected table_name payload: %v", payload["table_name"])
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":0,"msg":"ok","count":1,"data":[{"id":1,"name":"ins-a","db_type":"mysql","db_name":"archery","table_name":"users"}]}`))
	}))
	defer server.Close()

	rt := newTestRuntime(t, server.URL, token)
	result, err := rt.LookupTableInstances(nil, " users ")
	if err != nil {
		t.Fatalf("LookupTableInstances returned error: %v", err)
	}

	if result.Response.GetStatus() != 0 {
		t.Fatalf("unexpected status: %d", result.Response.GetStatus())
	}
	if len(result.Response.GetData()) != 1 {
		t.Fatalf("unexpected data length: %d", len(result.Response.GetData()))
	}
	if !strings.Contains(string(result.RawBody), `"count":1`) {
		t.Fatalf("raw body was not returned correctly: %s", string(result.RawBody))
	}
}

func TestDescribeTableStructure(t *testing.T) {
	const token = "token-describe"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/sqlquery/describetable/" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != token {
			t.Fatalf("unexpected auth header: %s", got)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		var payload map[string]any
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		if payload["instance_name"] != "ins-a" || payload["db_name"] != "archery" || payload["tb_name"] != "users" {
			t.Fatalf("unexpected payload: %v", payload)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":0,"msg":"ok","data":{"column_list":["Field","Type"],"rows":[["id","bigint"]]}}`))
	}))
	defer server.Close()

	rt := newTestRuntime(t, server.URL, token)
	result, err := rt.DescribeTableStructure(nil, "ins-a", "archery", "users", "")
	if err != nil {
		t.Fatalf("DescribeTableStructure returned error: %v", err)
	}

	if result.Envelope.Status != 0 {
		t.Fatalf("unexpected status: %d", result.Envelope.Status)
	}
	if result.Envelope.Msg != "ok" {
		t.Fatalf("unexpected msg: %s", result.Envelope.Msg)
	}
	if !strings.Contains(string(result.RawBody), `"column_list"`) {
		t.Fatalf("raw body was not returned correctly: %s", string(result.RawBody))
	}
}

func TestExecuteQuery(t *testing.T) {
	const token = "token-execute"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/sqlquery/execute/" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != token {
			t.Fatalf("unexpected auth header: %s", got)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read request body: %v", err)
		}
		var payload map[string]any
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		if payload["sql_content"] != "select 1" {
			t.Fatalf("unexpected sql_content: %v", payload["sql_content"])
		}
		if int(payload["limit_num"].(float64)) != 100 {
			t.Fatalf("unexpected limit_num: %v", payload["limit_num"])
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":0,"msg":"ok","data":{"column_list":["1"],"rows":[[1]]}}`))
	}))
	defer server.Close()

	rt := newTestRuntime(t, server.URL, token)
	result, err := rt.ExecuteQuery(nil, "ins-a", "archery", "", "", "select 1", 100)
	if err != nil {
		t.Fatalf("ExecuteQuery returned error: %v", err)
	}

	if result.Envelope.Status != 0 {
		t.Fatalf("unexpected status: %d", result.Envelope.Status)
	}
	if result.Envelope.Msg != "ok" {
		t.Fatalf("unexpected msg: %s", result.Envelope.Msg)
	}
	if !strings.Contains(string(result.RawBody), `"rows"`) {
		t.Fatalf("raw body was not returned correctly: %s", string(result.RawBody))
	}
}
