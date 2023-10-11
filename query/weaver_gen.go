// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package query

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "prototype/query/GithubUserQuery",
		Iface: reflect.TypeOf((*GithubUserQuery)(nil)).Elem(),
		Impl:  reflect.TypeOf(GithubQuery{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return githubUserQuery_local_stub{impl: impl.(GithubUserQuery), tracer: tracer, queryMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "prototype/query/GithubUserQuery", Method: "Query", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return githubUserQuery_client_stub{stub: stub, queryMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "prototype/query/GithubUserQuery", Method: "Query", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return githubUserQuery_server_stub{impl: impl.(GithubUserQuery), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return githubUserQuery_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[GithubUserQuery] = (*GithubQuery)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*GithubQuery)(nil)

// Local stub implementations.

type githubUserQuery_local_stub struct {
	impl         GithubUserQuery
	tracer       trace.Tracer
	queryMetrics *codegen.MethodMetrics
}

// Check that githubUserQuery_local_stub implements the GithubUserQuery interface.
var _ GithubUserQuery = (*githubUserQuery_local_stub)(nil)

func (s githubUserQuery_local_stub) Query(ctx context.Context, a0 string) (r0 []string, err error) {
	// Update metrics.
	begin := s.queryMetrics.Begin()
	defer func() { s.queryMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "query.GithubUserQuery.Query", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Query(ctx, a0)
}

// Client stub implementations.

type githubUserQuery_client_stub struct {
	stub         codegen.Stub
	queryMetrics *codegen.MethodMetrics
}

// Check that githubUserQuery_client_stub implements the GithubUserQuery interface.
var _ GithubUserQuery = (*githubUserQuery_client_stub)(nil)

func (s githubUserQuery_client_stub) Query(ctx context.Context, a0 string) (r0 []string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.queryMetrics.Begin()
	defer func() { s.queryMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "query.GithubUserQuery.Query", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_string_4af10117(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.21.2 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type githubUserQuery_server_stub struct {
	impl    GithubUserQuery
	addLoad func(key uint64, load float64)
}

// Check that githubUserQuery_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*githubUserQuery_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s githubUserQuery_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Query":
		return s.query
	default:
		return nil
	}
}

func (s githubUserQuery_server_stub) query(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Query(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type githubUserQuery_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that githubUserQuery_reflect_stub implements the GithubUserQuery interface.
var _ GithubUserQuery = (*githubUserQuery_reflect_stub)(nil)

func (s githubUserQuery_reflect_stub) Query(ctx context.Context, a0 string) (r0 []string, err error) {
	err = s.caller("Query", ctx, []any{a0}, []any{&r0})
	return
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}
