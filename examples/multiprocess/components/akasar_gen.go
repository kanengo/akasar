// Code generated by "akasar generate". DO NOT EDIT.
//go:build !ignoreAkasarGen

package components

import (
	"context"
	"errors"
	"github.com/kanengo/akasar"
	"github.com/kanengo/akasar/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/kanengo/akasar/examples/multiprocess/components/Store",
		Iface: reflect.TypeOf((*Store)(nil)).Elem(),
		Impl:  reflect.TypeOf(store{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return storeLocalStub{impl: impl.(Store), tracer: tracer, buyGoodsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/kanengo/akasar/examples/multiprocess/components/Store", Method: "BuyGoods", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string, tracer trace.Tracer) any {
			return storeClientStub{stub: stub, buyGoodsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/kanengo/akasar/examples/multiprocess/components/Store", Method: "BuyGoods", Remote: true})}
		},
		ServerStubFn: func(impl any) codegen.Server {
			return storeServerStub{impl: impl.(Store)}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/kanengo/akasar/examples/multiprocess/components/User",
		Iface: reflect.TypeOf((*User)(nil)).Elem(),
		Impl:  reflect.TypeOf(user{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return userLocalStub{impl: impl.(User), tracer: tracer, loginMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/kanengo/akasar/examples/multiprocess/components/User", Method: "Login", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string, tracer trace.Tracer) any {
			return userClientStub{stub: stub, loginMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/kanengo/akasar/examples/multiprocess/components/User", Method: "Login", Remote: true})}
		},
		ServerStubFn: func(impl any) codegen.Server {
			return userServerStub{impl: impl.(User)}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/kanengo/akasar/examples/multiprocess/components/VIP",
		Iface: reflect.TypeOf((*VIP)(nil)).Elem(),
		Impl:  reflect.TypeOf(vip{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return vIPLocalStub{impl: impl.(VIP), tracer: tracer, getVipInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/kanengo/akasar/examples/multiprocess/components/VIP", Method: "GetVipInfo", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string, tracer trace.Tracer) any {
			return vIPClientStub{stub: stub, getVipInfoMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/kanengo/akasar/examples/multiprocess/components/VIP", Method: "GetVipInfo", Remote: true})}
		},
		ServerStubFn: func(impl any) codegen.Server {
			return vIPServerStub{impl: impl.(VIP)}
		},
	})
}

// akasar.InstanceOf checks.
var _ akasar.InstanceOf[Store] = (*store)(nil)
var _ akasar.InstanceOf[User] = (*user)(nil)
var _ akasar.InstanceOf[VIP] = (*vip)(nil)

// Local stub implementations.

type storeLocalStub struct {
	impl            Store
	tracer          trace.Tracer
	buyGoodsMetrics *codegen.MethodMetrics
}

// Check that storeLocalStub implements the Store interface
var _ Store = (*storeLocalStub)(nil)

func (s storeLocalStub) BuyGoods(ctx context.Context, a0 int64, a1 int32) (err error) {
	// Update metrics.
	begin := s.buyGoodsMetrics.Begin()
	defer func() { s.buyGoodsMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "components.Store.BuyGoods", trace.WithSpanKind((trace.SpanKindInternal)))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}
	defer func() {
		if err == nil {
			err = codegen.CatchResultUnwrapPanic(recover())
		}
	}()

	err = s.impl.BuyGoods(ctx, a0, a1)
	return
}

type userLocalStub struct {
	impl         User
	tracer       trace.Tracer
	loginMetrics *codegen.MethodMetrics
}

// Check that userLocalStub implements the User interface
var _ User = (*userLocalStub)(nil)

func (s userLocalStub) Login(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	begin := s.loginMetrics.Begin()
	defer func() { s.loginMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "components.User.Login", trace.WithSpanKind((trace.SpanKindInternal)))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}
	defer func() {
		if err == nil {
			err = codegen.CatchResultUnwrapPanic(recover())
		}
	}()

	err = s.impl.Login(ctx, a0, a1)
	return
}

type vIPLocalStub struct {
	impl              VIP
	tracer            trace.Tracer
	getVipInfoMetrics *codegen.MethodMetrics
}

// Check that vIPLocalStub implements the VIP interface
var _ VIP = (*vIPLocalStub)(nil)

func (s vIPLocalStub) GetVipInfo(ctx context.Context, a0 int64) (err error) {
	// Update metrics.
	begin := s.getVipInfoMetrics.Begin()
	defer func() { s.getVipInfoMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "components.VIP.GetVipInfo", trace.WithSpanKind((trace.SpanKindInternal)))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}
	defer func() {
		if err == nil {
			err = codegen.CatchResultUnwrapPanic(recover())
		}
	}()

	err = s.impl.GetVipInfo(ctx, a0)
	return
}

// Client stub implementations.

type storeClientStub struct {
	stub            codegen.Stub
	tracer          trace.Tracer
	buyGoodsMetrics *codegen.MethodMetrics
}

// Check that storeClientStub implements the Store interface
var _ Store = (*storeClientStub)(nil)

func (s storeClientStub) BuyGoods(ctx context.Context, a0 int64, a1 int32) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.buyGoodsMetrics.Begin()
	defer func() { s.buyGoodsMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "components.Store.BuyGoods", trace.WithSpanKind((trace.SpanKindInternal)))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(akasar.RemoteCallError, err)
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
	size += 8
	size += 4
	enc := codegen.NewSerializer(size)

	// Encode arguments.
	enc.Int64(a0)
	enc.Int32(a1)
	var shardKey uint64

	// Call the remote method.
	data := enc.Data()
	requestBytes = len(data)
	var results []byte
	results, err = s.stub.Invoke(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(akasar.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDeserializer(results)
	err = dec.Error()

	return
}

type userClientStub struct {
	stub         codegen.Stub
	tracer       trace.Tracer
	loginMetrics *codegen.MethodMetrics
}

// Check that userClientStub implements the User interface
var _ User = (*userClientStub)(nil)

func (s userClientStub) Login(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.loginMetrics.Begin()
	defer func() { s.loginMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "components.User.Login", trace.WithSpanKind((trace.SpanKindInternal)))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(akasar.RemoteCallError, err)
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
	size += (4 + len(a1))
	enc := codegen.NewSerializer(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)
	var shardKey uint64

	// Call the remote method.
	data := enc.Data()
	requestBytes = len(data)
	var results []byte
	results, err = s.stub.Invoke(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(akasar.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDeserializer(results)
	err = dec.Error()

	return
}

type vIPClientStub struct {
	stub              codegen.Stub
	tracer            trace.Tracer
	getVipInfoMetrics *codegen.MethodMetrics
}

// Check that vIPClientStub implements the VIP interface
var _ VIP = (*vIPClientStub)(nil)

func (s vIPClientStub) GetVipInfo(ctx context.Context, a0 int64) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getVipInfoMetrics.Begin()
	defer func() { s.getVipInfoMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		ctx, span = s.tracer.Start(ctx, "components.VIP.GetVipInfo", trace.WithSpanKind((trace.SpanKindInternal)))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(akasar.RemoteCallError, err)
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
	size += 8
	enc := codegen.NewSerializer(size)

	// Encode arguments.
	enc.Int64(a0)
	var shardKey uint64

	// Call the remote method.
	data := enc.Data()
	requestBytes = len(data)
	var results []byte
	results, err = s.stub.Invoke(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(akasar.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDeserializer(results)
	err = dec.Error()

	return
}

// Server stub implementation.

type storeServerStub struct {
	impl Store
}

// Check that storeServerStub is implements the codegen.Server interface.
var _ codegen.Server = (*storeServerStub)(nil)

// GetHandleFn implements the codegen.Server interface.
func (s storeServerStub) GetHandleFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "BuyGoods":
		return s.buyGoods
	default:
		return nil
	}
}

func (s *storeServerStub) buyGoods(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Encode arguments.
	dec := codegen.NewDeserializer(args)
	var a0 int64
	a0 = dec.Int64()
	var a1 int32
	a1 = dec.Int32()

	appErr := s.impl.BuyGoods(ctx, a0, a1)

	//Encode the results.
	enc := codegen.NewSerializer()
	enc.Error(appErr)
	return enc.Data(), nil
}

type userServerStub struct {
	impl User
}

// Check that userServerStub is implements the codegen.Server interface.
var _ codegen.Server = (*userServerStub)(nil)

// GetHandleFn implements the codegen.Server interface.
func (s userServerStub) GetHandleFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Login":
		return s.login
	default:
		return nil
	}
}

func (s *userServerStub) login(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Encode arguments.
	dec := codegen.NewDeserializer(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()

	appErr := s.impl.Login(ctx, a0, a1)

	//Encode the results.
	enc := codegen.NewSerializer()
	enc.Error(appErr)
	return enc.Data(), nil
}

type vIPServerStub struct {
	impl VIP
}

// Check that vIPServerStub is implements the codegen.Server interface.
var _ codegen.Server = (*vIPServerStub)(nil)

// GetHandleFn implements the codegen.Server interface.
func (s vIPServerStub) GetHandleFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetVipInfo":
		return s.getVipInfo
	default:
		return nil
	}
}

func (s *vIPServerStub) getVipInfo(ctx context.Context, args []byte) (res []byte, err error) {
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Encode arguments.
	dec := codegen.NewDeserializer(args)
	var a0 int64
	a0 = dec.Int64()

	appErr := s.impl.GetVipInfo(ctx, a0)

	//Encode the results.
	enc := codegen.NewSerializer()
	enc.Error(appErr)
	return enc.Data(), nil
}
