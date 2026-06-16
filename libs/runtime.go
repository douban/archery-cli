package libs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	openapi "github.intra.douban.com/sa/archery-cli"
)

type QueryEnvelope struct {
	Status int            `json:"status"`
	Msg    string         `json:"msg"`
	Data   map[string]any `json:"data"`
}

type TableLookupResult struct {
	Response *openapi.TableInstanceLookupResponse
	RawBody  []byte
}

type DescribeTableResult struct {
	Envelope QueryEnvelope
	RawBody  []byte
}

type ExecuteQueryResult struct {
	Envelope QueryEnvelope
	RawBody  []byte
}

type Runtime struct {
	ConfigPath string
	Config     Config
	Client     *openapi.APIClient
	Context    context.Context
}

func (r *Runtime) effectiveContext(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	}
	if r != nil && r.Context != nil {
		return r.Context
	}
	return context.Background()
}

func LoadRuntime(configPath string) (*Runtime, error) {
	resolvedPath, err := ResolveConfigPath(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	clientConfig := openapi.NewConfiguration()
	clientConfig.Servers = openapi.ServerConfigurations{{
		URL:         strings.TrimRight(cfg.BaseURL, "/"),
		Description: "Configured Archery API server",
	}}

	client := openapi.NewAPIClient(clientConfig)
	ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, map[string]openapi.APIKey{
		"knoxApiToken": {Key: cfg.APIToken, Prefix: "Token"},
	})

	return &Runtime{
		ConfigPath: resolvedPath,
		Config:     cfg,
		Client:     client,
		Context:    ctx,
	}, nil
}

func (r *Runtime) LookupTableInstances(ctx context.Context, tableName string) (*TableLookupResult, error) {
	if r == nil {
		return nil, fmt.Errorf("runtime is nil")
	}

	ctx = r.effectiveContext(ctx)
	request := openapi.TableInstanceLookup{TableName: strings.TrimSpace(tableName)}
	response, rawResponse, err := r.Client.ApiAPI.ApiV1InstanceTableInstancesCreate(ctx).TableInstanceLookup(request).Execute()
	if err != nil {
		return nil, err
	}
	if rawResponse == nil {
		return nil, fmt.Errorf("empty response")
	}

	body, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("read api response: %w", err)
	}
	rawResponse.Body.Close()
	rawResponse.Body = io.NopCloser(bytes.NewBuffer(body))

	return &TableLookupResult{Response: response, RawBody: body}, nil
}

func (r *Runtime) DescribeTableStructure(ctx context.Context, instanceName, dbName, tableName, schemaName string) (*DescribeTableResult, error) {
	if r == nil {
		return nil, fmt.Errorf("runtime is nil")
	}

	ctx = r.effectiveContext(ctx)
	payload := map[string]any{
		"instance_name": strings.TrimSpace(instanceName),
		"db_name":       strings.TrimSpace(dbName),
		"tb_name":       strings.TrimSpace(tableName),
		"schema_name":   strings.TrimSpace(schemaName),
	}

	body, err := r.postJSON(ctx, "/api/v1/sqlquery/describetable/", payload)
	if err != nil {
		return nil, err
	}

	var envelope QueryEnvelope
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, fmt.Errorf("decode describe-table response: %w", err)
	}

	return &DescribeTableResult{Envelope: envelope, RawBody: body}, nil
}

func (r *Runtime) ExecuteQuery(ctx context.Context, instanceName, dbName, schemaName, tableName, sqlContent string, limitNum int) (*ExecuteQueryResult, error) {
	if r == nil {
		return nil, fmt.Errorf("runtime is nil")
	}

	ctx = r.effectiveContext(ctx)
	payload := map[string]any{
		"instance_name": strings.TrimSpace(instanceName),
		"db_name":       strings.TrimSpace(dbName),
		"schema_name":   strings.TrimSpace(schemaName),
		"tb_name":       strings.TrimSpace(tableName),
		"sql_content":   sqlContent,
		"limit_num":     limitNum,
	}

	body, err := r.postJSON(ctx, "/api/v1/sqlquery/execute/", payload)
	if err != nil {
		return nil, err
	}

	var envelope QueryEnvelope
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, fmt.Errorf("decode execute response: %w", err)
	}

	return &ExecuteQueryResult{Envelope: envelope, RawBody: body}, nil
}

func (r *Runtime) postJSON(ctx context.Context, path string, payload any) ([]byte, error) {
	if r == nil {
		return nil, fmt.Errorf("runtime is nil")
	}

	baseURL := strings.TrimRight(r.Config.BaseURL, "/")
	if baseURL == "" {
		return nil, fmt.Errorf("base_url is required")
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}

	ctx = r.effectiveContext(ctx)

	localVarHeaderParams := make(map[string]string)
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(openapi.ContextAPIKeys).(map[string]openapi.APIKey); ok {
			if apiKey, ok := auth["knoxApiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+path, bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"
	for k, v := range localVarHeaderParams {
		request.Header.Set(k, v)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("call api: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read api response: %w", err)
	}

	if response.StatusCode >= http.StatusMultipleChoices {
		if len(body) == 0 {
			return nil, fmt.Errorf("api request failed: %s", response.Status)
		}
		return nil, fmt.Errorf("api request failed: %s: %s", response.Status, strings.TrimSpace(string(body)))
	}

	return body, nil
}
