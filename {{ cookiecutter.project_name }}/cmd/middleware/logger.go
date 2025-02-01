package middleware

import (
	"bytes"
	"io"
	"os"
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ResponseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseBodyWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger(config cfg.AppConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &ResponseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = bodyWriter

		start_time := time.Now()
		method := ctx.Request.Method
		url := ctx.Request.URL
		version := "12345"
		request := get_request_body(ctx, config)
		client_id := ""

		ctx.Next()

		status_code := ctx.Writer.Status()
		duration := time.Since(start_time)
		response := get_response_body(ctx, config, bodyWriter)
		trace_id := ctx.GetString("x-trace-id")

		if !ignore_tests() && !is_ignored_path(ctx) {
			log.Info().
				Str("version", utils.IF(version != "", version, "UNVERSIONED").(string)).
				Str("trace_id", trace_id).
				Str("client_id", utils.IF(client_id != "", client_id, "UNKNOWN").(string)).
				Str("method", method).
				Str("url", url.String()).
				Int("status_code", status_code).
				Dur("duration", duration).
				RawJSON("request", []byte(utils.IF(request != "", request, "{}").(string))).
				RawJSON("response", []byte(utils.IF(response != "", response, "{}").(string))).
				Msg("http_logger")
		}
	}
}

func ignore_tests() bool {
	return os.Getenv("environment") == "integration"
}

func is_ignored_path(ctx *gin.Context) bool {
	url := ctx.Request.URL.Path
	ignored_paths := []string{"/metrics"}

	paths := utils.Filter(ignored_paths, func(path string) bool { return path == url })
	return len(paths) != 0
}

func is_sensitive_req(ctx *gin.Context) bool {
	url := ctx.Request.URL.Path
	sensitive_paths := []string{"/login", "/signup", "/me"}

	paths := utils.Filter(sensitive_paths, func(path string) bool { return path == url })
	return len(paths) != 0
}

func get_request_body(ctx *gin.Context, config cfg.AppConfig) string {
	debug := ctx.GetHeader("debugger")

	if debug == config.OBSERVABILITY.DEBUGGER_KEY && !is_sensitive_req(ctx) {
		var buf bytes.Buffer
		tee := io.TeeReader(ctx.Request.Body, &buf)
		body, _ := io.ReadAll(tee)
		ctx.Request.Body = io.NopCloser(&buf)
		result := string(body)
		return result
	} else {
		return ""
	}
}

func get_response_body(ctx *gin.Context, config cfg.AppConfig, writer *ResponseBodyWriter) string {
	debug := ctx.GetHeader("debugger")

	if debug == config.OBSERVABILITY.DEBUGGER_KEY && !is_sensitive_req(ctx) {
		result := writer.body.String()
		return result
	} else {
		return ""
	}
}
