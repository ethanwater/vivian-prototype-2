// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package main

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
		Name:  "prototype/GithubUserQuery",
		Iface: reflect.TypeOf((*GithubUserQuery)(nil)).Elem(),
		Impl:  reflect.TypeOf(githubQuery{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return githubUserQuery_local_stub{impl: impl.(GithubUserQuery), tracer: tracer, queryMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "prototype/GithubUserQuery", Method: "Query", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return githubUserQuery_client_stub{stub: stub, queryMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "prototype/GithubUserQuery", Method: "Query", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return githubUserQuery_server_stub{impl: impl.(GithubUserQuery), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return githubUserQuery_reflect_stub{caller: caller}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(app{}),
		Listeners: []string{"proto"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦68914653:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→prototype/GithubUserQuery⟧\n⟦b00836f0:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→proto⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[GithubUserQuery] = (*githubQuery)(nil)
var _ weaver.InstanceOf[weaver.Main] = (*app)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*githubQuery)(nil)
var _ weaver.Unrouted = (*app)(nil)

// Local stub implementations.

type githubUserQuery_local_stub struct {
	impl         GithubUserQuery
	tracer       trace.Tracer
	queryMetrics *codegen.MethodMetrics
}

// Check that githubUserQuery_local_stub implements the GithubUserQuery interface.
var _ GithubUserQuery = (*githubUserQuery_local_stub)(nil)

func (s githubUserQuery_local_stub) Query(ctx context.Context) (err error) {
	// Update metrics.
	begin := s.queryMetrics.Begin()
	defer func() { s.queryMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.GithubUserQuery.Query", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Query(ctx)
}

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

// Client stub implementations.

type githubUserQuery_client_stub struct {
	stub         codegen.Stub
	queryMetrics *codegen.MethodMetrics
}

// Check that githubUserQuery_client_stub implements the GithubUserQuery interface.
var _ GithubUserQuery = (*githubUserQuery_client_stub)(nil)

func (s githubUserQuery_client_stub) Query(ctx context.Context) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.queryMetrics.Begin()
	defer func() { s.queryMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.GithubUserQuery.Query", trace.WithSpanKind(trace.SpanKindClient))
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

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

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

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Query(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

// Reflect stub implementations.

type githubUserQuery_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that githubUserQuery_reflect_stub implements the GithubUserQuery interface.
var _ GithubUserQuery = (*githubUserQuery_reflect_stub)(nil)

func (s githubUserQuery_reflect_stub) Query(ctx context.Context) (err error) {
	err = s.caller("Query", ctx, []any{}, []any{})
	return
}

type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that main_reflect_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_reflect_stub)(nil)

