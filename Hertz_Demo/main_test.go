package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/stretchr/testify/assert"
)


func TestDecode_ValidRequest(t *testing.T) {
	body := `{"userId": "test id", "message": "test"}`

	ctx := &app.RequestContext{
		Request: *protocol.NewRequest(http.MethodGet, "/ServiceA/methodA", nil),
	}
	ctx.Request.SetHeader("Content-Type", "application/json")
	ctx.Request.AppendBodyString(body)

	testC := context.Background()

	decode(testC, ctx)

	assert.Equal(t, http.StatusOK, ctx.Response.StatusCode())

	expected := "{\"message\":\"Usertest id Connected to ServiceA, methodA.\\nMessage content:test\"}"

	assert.JSONEq(t, expected, string(ctx.Response.Body()))
}

func TestInvalidContentType_ValidRequest(t *testing.T) {
	ctx := &app.RequestContext{}
	ctx.Request.SetHeader("Content-Type", "application/json")

	boolean := invalidContentType(ctx)
	assert.False(t, boolean)
}

func TestInvalidContentType_InvalidRequest(t *testing.T) {
	ctx := &app.RequestContext{}
	ctx.Request.SetHeader("Content-Type", "text/plain")

	boolean := invalidContentType(ctx)
	assert.True(t, boolean)
}

func TestReadPath(t *testing.T) {
	ctx := &app.RequestContext{}
	ctx.Request.SetRequestURI("/foo/bar")

	arr := readPath(ctx)
	assert.Equal(t, []string{"", "foo", "bar"}, arr)
}

func TestParseRequestBody_ValidRequest(t *testing.T) {
	ctx := &app.RequestContext{}
	ctx.Request.SetBodyString("{\"key\":\"value\"}")

	mapping, err := parseRequestBody(ctx.Request.Body())

	assert.Nil(t, err)

	expected := map[string]interface{}{"key": "value"}
	assert.Equal(t, expected, mapping)
}

func TestParseRequestBody_InvalidRequest1(t *testing.T) {
	ctx := &app.RequestContext{}
	ctx.Request.SetBodyString("{\"hello\"")

	mapping, err := parseRequestBody(ctx.Request.Body())

	assert.Error(t, err)

	expected := map[string]interface{}(nil)
	assert.Equal(t, expected, mapping)
}

func TestParseRequestBody_InvalidRequest2(t *testing.T) {
	ctx := &app.RequestContext{}
	mapping, err := parseRequestBody(ctx.Request.Body())

	assert.Error(t, err)

	expected := map[string]interface{}(nil)
	assert.Equal(t, expected, mapping)
}
