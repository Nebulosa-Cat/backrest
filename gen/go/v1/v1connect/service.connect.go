// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/service.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	types "github.com/garethgeorge/backrest/gen/go/types"
	v1 "github.com/garethgeorge/backrest/gen/go/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BackrestName is the fully-qualified name of the Backrest service.
	BackrestName = "v1.Backrest"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BackrestGetConfigProcedure is the fully-qualified name of the Backrest's GetConfig RPC.
	BackrestGetConfigProcedure = "/v1.Backrest/GetConfig"
	// BackrestSetConfigProcedure is the fully-qualified name of the Backrest's SetConfig RPC.
	BackrestSetConfigProcedure = "/v1.Backrest/SetConfig"
	// BackrestAddRepoProcedure is the fully-qualified name of the Backrest's AddRepo RPC.
	BackrestAddRepoProcedure = "/v1.Backrest/AddRepo"
	// BackrestGetOperationEventsProcedure is the fully-qualified name of the Backrest's
	// GetOperationEvents RPC.
	BackrestGetOperationEventsProcedure = "/v1.Backrest/GetOperationEvents"
	// BackrestGetOperationsProcedure is the fully-qualified name of the Backrest's GetOperations RPC.
	BackrestGetOperationsProcedure = "/v1.Backrest/GetOperations"
	// BackrestListSnapshotsProcedure is the fully-qualified name of the Backrest's ListSnapshots RPC.
	BackrestListSnapshotsProcedure = "/v1.Backrest/ListSnapshots"
	// BackrestListSnapshotFilesProcedure is the fully-qualified name of the Backrest's
	// ListSnapshotFiles RPC.
	BackrestListSnapshotFilesProcedure = "/v1.Backrest/ListSnapshotFiles"
	// BackrestBackupProcedure is the fully-qualified name of the Backrest's Backup RPC.
	BackrestBackupProcedure = "/v1.Backrest/Backup"
	// BackrestPruneProcedure is the fully-qualified name of the Backrest's Prune RPC.
	BackrestPruneProcedure = "/v1.Backrest/Prune"
	// BackrestForgetProcedure is the fully-qualified name of the Backrest's Forget RPC.
	BackrestForgetProcedure = "/v1.Backrest/Forget"
	// BackrestRestoreProcedure is the fully-qualified name of the Backrest's Restore RPC.
	BackrestRestoreProcedure = "/v1.Backrest/Restore"
	// BackrestUnlockProcedure is the fully-qualified name of the Backrest's Unlock RPC.
	BackrestUnlockProcedure = "/v1.Backrest/Unlock"
	// BackrestCancelProcedure is the fully-qualified name of the Backrest's Cancel RPC.
	BackrestCancelProcedure = "/v1.Backrest/Cancel"
	// BackrestClearHistoryProcedure is the fully-qualified name of the Backrest's ClearHistory RPC.
	BackrestClearHistoryProcedure = "/v1.Backrest/ClearHistory"
	// BackrestPathAutocompleteProcedure is the fully-qualified name of the Backrest's PathAutocomplete
	// RPC.
	BackrestPathAutocompleteProcedure = "/v1.Backrest/PathAutocomplete"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	backrestServiceDescriptor                  = v1.File_v1_service_proto.Services().ByName("Backrest")
	backrestGetConfigMethodDescriptor          = backrestServiceDescriptor.Methods().ByName("GetConfig")
	backrestSetConfigMethodDescriptor          = backrestServiceDescriptor.Methods().ByName("SetConfig")
	backrestAddRepoMethodDescriptor            = backrestServiceDescriptor.Methods().ByName("AddRepo")
	backrestGetOperationEventsMethodDescriptor = backrestServiceDescriptor.Methods().ByName("GetOperationEvents")
	backrestGetOperationsMethodDescriptor      = backrestServiceDescriptor.Methods().ByName("GetOperations")
	backrestListSnapshotsMethodDescriptor      = backrestServiceDescriptor.Methods().ByName("ListSnapshots")
	backrestListSnapshotFilesMethodDescriptor  = backrestServiceDescriptor.Methods().ByName("ListSnapshotFiles")
	backrestBackupMethodDescriptor             = backrestServiceDescriptor.Methods().ByName("Backup")
	backrestPruneMethodDescriptor              = backrestServiceDescriptor.Methods().ByName("Prune")
	backrestForgetMethodDescriptor             = backrestServiceDescriptor.Methods().ByName("Forget")
	backrestRestoreMethodDescriptor            = backrestServiceDescriptor.Methods().ByName("Restore")
	backrestUnlockMethodDescriptor             = backrestServiceDescriptor.Methods().ByName("Unlock")
	backrestCancelMethodDescriptor             = backrestServiceDescriptor.Methods().ByName("Cancel")
	backrestClearHistoryMethodDescriptor       = backrestServiceDescriptor.Methods().ByName("ClearHistory")
	backrestPathAutocompleteMethodDescriptor   = backrestServiceDescriptor.Methods().ByName("PathAutocomplete")
)

// BackrestClient is a client for the v1.Backrest service.
type BackrestClient interface {
	GetConfig(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Config], error)
	SetConfig(context.Context, *connect.Request[v1.Config]) (*connect.Response[v1.Config], error)
	AddRepo(context.Context, *connect.Request[v1.Repo]) (*connect.Response[v1.Config], error)
	GetOperationEvents(context.Context, *connect.Request[emptypb.Empty]) (*connect.ServerStreamForClient[v1.OperationEvent], error)
	GetOperations(context.Context, *connect.Request[v1.GetOperationsRequest]) (*connect.Response[v1.OperationList], error)
	ListSnapshots(context.Context, *connect.Request[v1.ListSnapshotsRequest]) (*connect.Response[v1.ResticSnapshotList], error)
	ListSnapshotFiles(context.Context, *connect.Request[v1.ListSnapshotFilesRequest]) (*connect.Response[v1.ListSnapshotFilesResponse], error)
	// Backup schedules a backup operation. It accepts a plan id and returns empty if the task is enqueued.
	Backup(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error)
	// Prune schedules a prune operation.
	Prune(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error)
	// Forget schedules a forget operation.
	Forget(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error)
	// Restore schedules a restore operation.
	Restore(context.Context, *connect.Request[v1.RestoreSnapshotRequest]) (*connect.Response[emptypb.Empty], error)
	// Unlock synchronously attempts to unlock the repo. Will block if other operations are in progress.
	Unlock(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error)
	// Cancel attempts to cancel a task with the given operation ID. Not guaranteed to succeed.
	Cancel(context.Context, *connect.Request[types.Int64Value]) (*connect.Response[emptypb.Empty], error)
	ClearHistory(context.Context, *connect.Request[v1.ClearHistoryRequest]) (*connect.Response[emptypb.Empty], error)
	// PathAutocomplete provides path autocompletion options for a given filesystem path.
	PathAutocomplete(context.Context, *connect.Request[types.StringValue]) (*connect.Response[types.StringList], error)
}

// NewBackrestClient constructs a client for the v1.Backrest service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBackrestClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BackrestClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &backrestClient{
		getConfig: connect.NewClient[emptypb.Empty, v1.Config](
			httpClient,
			baseURL+BackrestGetConfigProcedure,
			connect.WithSchema(backrestGetConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setConfig: connect.NewClient[v1.Config, v1.Config](
			httpClient,
			baseURL+BackrestSetConfigProcedure,
			connect.WithSchema(backrestSetConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addRepo: connect.NewClient[v1.Repo, v1.Config](
			httpClient,
			baseURL+BackrestAddRepoProcedure,
			connect.WithSchema(backrestAddRepoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getOperationEvents: connect.NewClient[emptypb.Empty, v1.OperationEvent](
			httpClient,
			baseURL+BackrestGetOperationEventsProcedure,
			connect.WithSchema(backrestGetOperationEventsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getOperations: connect.NewClient[v1.GetOperationsRequest, v1.OperationList](
			httpClient,
			baseURL+BackrestGetOperationsProcedure,
			connect.WithSchema(backrestGetOperationsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listSnapshots: connect.NewClient[v1.ListSnapshotsRequest, v1.ResticSnapshotList](
			httpClient,
			baseURL+BackrestListSnapshotsProcedure,
			connect.WithSchema(backrestListSnapshotsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listSnapshotFiles: connect.NewClient[v1.ListSnapshotFilesRequest, v1.ListSnapshotFilesResponse](
			httpClient,
			baseURL+BackrestListSnapshotFilesProcedure,
			connect.WithSchema(backrestListSnapshotFilesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		backup: connect.NewClient[types.StringValue, emptypb.Empty](
			httpClient,
			baseURL+BackrestBackupProcedure,
			connect.WithSchema(backrestBackupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		prune: connect.NewClient[types.StringValue, emptypb.Empty](
			httpClient,
			baseURL+BackrestPruneProcedure,
			connect.WithSchema(backrestPruneMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		forget: connect.NewClient[types.StringValue, emptypb.Empty](
			httpClient,
			baseURL+BackrestForgetProcedure,
			connect.WithSchema(backrestForgetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		restore: connect.NewClient[v1.RestoreSnapshotRequest, emptypb.Empty](
			httpClient,
			baseURL+BackrestRestoreProcedure,
			connect.WithSchema(backrestRestoreMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		unlock: connect.NewClient[types.StringValue, emptypb.Empty](
			httpClient,
			baseURL+BackrestUnlockProcedure,
			connect.WithSchema(backrestUnlockMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		cancel: connect.NewClient[types.Int64Value, emptypb.Empty](
			httpClient,
			baseURL+BackrestCancelProcedure,
			connect.WithSchema(backrestCancelMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		clearHistory: connect.NewClient[v1.ClearHistoryRequest, emptypb.Empty](
			httpClient,
			baseURL+BackrestClearHistoryProcedure,
			connect.WithSchema(backrestClearHistoryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		pathAutocomplete: connect.NewClient[types.StringValue, types.StringList](
			httpClient,
			baseURL+BackrestPathAutocompleteProcedure,
			connect.WithSchema(backrestPathAutocompleteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// backrestClient implements BackrestClient.
type backrestClient struct {
	getConfig          *connect.Client[emptypb.Empty, v1.Config]
	setConfig          *connect.Client[v1.Config, v1.Config]
	addRepo            *connect.Client[v1.Repo, v1.Config]
	getOperationEvents *connect.Client[emptypb.Empty, v1.OperationEvent]
	getOperations      *connect.Client[v1.GetOperationsRequest, v1.OperationList]
	listSnapshots      *connect.Client[v1.ListSnapshotsRequest, v1.ResticSnapshotList]
	listSnapshotFiles  *connect.Client[v1.ListSnapshotFilesRequest, v1.ListSnapshotFilesResponse]
	backup             *connect.Client[types.StringValue, emptypb.Empty]
	prune              *connect.Client[types.StringValue, emptypb.Empty]
	forget             *connect.Client[types.StringValue, emptypb.Empty]
	restore            *connect.Client[v1.RestoreSnapshotRequest, emptypb.Empty]
	unlock             *connect.Client[types.StringValue, emptypb.Empty]
	cancel             *connect.Client[types.Int64Value, emptypb.Empty]
	clearHistory       *connect.Client[v1.ClearHistoryRequest, emptypb.Empty]
	pathAutocomplete   *connect.Client[types.StringValue, types.StringList]
}

// GetConfig calls v1.Backrest.GetConfig.
func (c *backrestClient) GetConfig(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.Config], error) {
	return c.getConfig.CallUnary(ctx, req)
}

// SetConfig calls v1.Backrest.SetConfig.
func (c *backrestClient) SetConfig(ctx context.Context, req *connect.Request[v1.Config]) (*connect.Response[v1.Config], error) {
	return c.setConfig.CallUnary(ctx, req)
}

// AddRepo calls v1.Backrest.AddRepo.
func (c *backrestClient) AddRepo(ctx context.Context, req *connect.Request[v1.Repo]) (*connect.Response[v1.Config], error) {
	return c.addRepo.CallUnary(ctx, req)
}

// GetOperationEvents calls v1.Backrest.GetOperationEvents.
func (c *backrestClient) GetOperationEvents(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.ServerStreamForClient[v1.OperationEvent], error) {
	return c.getOperationEvents.CallServerStream(ctx, req)
}

// GetOperations calls v1.Backrest.GetOperations.
func (c *backrestClient) GetOperations(ctx context.Context, req *connect.Request[v1.GetOperationsRequest]) (*connect.Response[v1.OperationList], error) {
	return c.getOperations.CallUnary(ctx, req)
}

// ListSnapshots calls v1.Backrest.ListSnapshots.
func (c *backrestClient) ListSnapshots(ctx context.Context, req *connect.Request[v1.ListSnapshotsRequest]) (*connect.Response[v1.ResticSnapshotList], error) {
	return c.listSnapshots.CallUnary(ctx, req)
}

// ListSnapshotFiles calls v1.Backrest.ListSnapshotFiles.
func (c *backrestClient) ListSnapshotFiles(ctx context.Context, req *connect.Request[v1.ListSnapshotFilesRequest]) (*connect.Response[v1.ListSnapshotFilesResponse], error) {
	return c.listSnapshotFiles.CallUnary(ctx, req)
}

// Backup calls v1.Backrest.Backup.
func (c *backrestClient) Backup(ctx context.Context, req *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error) {
	return c.backup.CallUnary(ctx, req)
}

// Prune calls v1.Backrest.Prune.
func (c *backrestClient) Prune(ctx context.Context, req *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error) {
	return c.prune.CallUnary(ctx, req)
}

// Forget calls v1.Backrest.Forget.
func (c *backrestClient) Forget(ctx context.Context, req *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error) {
	return c.forget.CallUnary(ctx, req)
}

// Restore calls v1.Backrest.Restore.
func (c *backrestClient) Restore(ctx context.Context, req *connect.Request[v1.RestoreSnapshotRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.restore.CallUnary(ctx, req)
}

// Unlock calls v1.Backrest.Unlock.
func (c *backrestClient) Unlock(ctx context.Context, req *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error) {
	return c.unlock.CallUnary(ctx, req)
}

// Cancel calls v1.Backrest.Cancel.
func (c *backrestClient) Cancel(ctx context.Context, req *connect.Request[types.Int64Value]) (*connect.Response[emptypb.Empty], error) {
	return c.cancel.CallUnary(ctx, req)
}

// ClearHistory calls v1.Backrest.ClearHistory.
func (c *backrestClient) ClearHistory(ctx context.Context, req *connect.Request[v1.ClearHistoryRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.clearHistory.CallUnary(ctx, req)
}

// PathAutocomplete calls v1.Backrest.PathAutocomplete.
func (c *backrestClient) PathAutocomplete(ctx context.Context, req *connect.Request[types.StringValue]) (*connect.Response[types.StringList], error) {
	return c.pathAutocomplete.CallUnary(ctx, req)
}

// BackrestHandler is an implementation of the v1.Backrest service.
type BackrestHandler interface {
	GetConfig(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Config], error)
	SetConfig(context.Context, *connect.Request[v1.Config]) (*connect.Response[v1.Config], error)
	AddRepo(context.Context, *connect.Request[v1.Repo]) (*connect.Response[v1.Config], error)
	GetOperationEvents(context.Context, *connect.Request[emptypb.Empty], *connect.ServerStream[v1.OperationEvent]) error
	GetOperations(context.Context, *connect.Request[v1.GetOperationsRequest]) (*connect.Response[v1.OperationList], error)
	ListSnapshots(context.Context, *connect.Request[v1.ListSnapshotsRequest]) (*connect.Response[v1.ResticSnapshotList], error)
	ListSnapshotFiles(context.Context, *connect.Request[v1.ListSnapshotFilesRequest]) (*connect.Response[v1.ListSnapshotFilesResponse], error)
	// Backup schedules a backup operation. It accepts a plan id and returns empty if the task is enqueued.
	Backup(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error)
	// Prune schedules a prune operation.
	Prune(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error)
	// Forget schedules a forget operation.
	Forget(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error)
	// Restore schedules a restore operation.
	Restore(context.Context, *connect.Request[v1.RestoreSnapshotRequest]) (*connect.Response[emptypb.Empty], error)
	// Unlock synchronously attempts to unlock the repo. Will block if other operations are in progress.
	Unlock(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error)
	// Cancel attempts to cancel a task with the given operation ID. Not guaranteed to succeed.
	Cancel(context.Context, *connect.Request[types.Int64Value]) (*connect.Response[emptypb.Empty], error)
	ClearHistory(context.Context, *connect.Request[v1.ClearHistoryRequest]) (*connect.Response[emptypb.Empty], error)
	// PathAutocomplete provides path autocompletion options for a given filesystem path.
	PathAutocomplete(context.Context, *connect.Request[types.StringValue]) (*connect.Response[types.StringList], error)
}

// NewBackrestHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBackrestHandler(svc BackrestHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	backrestGetConfigHandler := connect.NewUnaryHandler(
		BackrestGetConfigProcedure,
		svc.GetConfig,
		connect.WithSchema(backrestGetConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestSetConfigHandler := connect.NewUnaryHandler(
		BackrestSetConfigProcedure,
		svc.SetConfig,
		connect.WithSchema(backrestSetConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestAddRepoHandler := connect.NewUnaryHandler(
		BackrestAddRepoProcedure,
		svc.AddRepo,
		connect.WithSchema(backrestAddRepoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestGetOperationEventsHandler := connect.NewServerStreamHandler(
		BackrestGetOperationEventsProcedure,
		svc.GetOperationEvents,
		connect.WithSchema(backrestGetOperationEventsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestGetOperationsHandler := connect.NewUnaryHandler(
		BackrestGetOperationsProcedure,
		svc.GetOperations,
		connect.WithSchema(backrestGetOperationsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestListSnapshotsHandler := connect.NewUnaryHandler(
		BackrestListSnapshotsProcedure,
		svc.ListSnapshots,
		connect.WithSchema(backrestListSnapshotsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestListSnapshotFilesHandler := connect.NewUnaryHandler(
		BackrestListSnapshotFilesProcedure,
		svc.ListSnapshotFiles,
		connect.WithSchema(backrestListSnapshotFilesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestBackupHandler := connect.NewUnaryHandler(
		BackrestBackupProcedure,
		svc.Backup,
		connect.WithSchema(backrestBackupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestPruneHandler := connect.NewUnaryHandler(
		BackrestPruneProcedure,
		svc.Prune,
		connect.WithSchema(backrestPruneMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestForgetHandler := connect.NewUnaryHandler(
		BackrestForgetProcedure,
		svc.Forget,
		connect.WithSchema(backrestForgetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestRestoreHandler := connect.NewUnaryHandler(
		BackrestRestoreProcedure,
		svc.Restore,
		connect.WithSchema(backrestRestoreMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestUnlockHandler := connect.NewUnaryHandler(
		BackrestUnlockProcedure,
		svc.Unlock,
		connect.WithSchema(backrestUnlockMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestCancelHandler := connect.NewUnaryHandler(
		BackrestCancelProcedure,
		svc.Cancel,
		connect.WithSchema(backrestCancelMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestClearHistoryHandler := connect.NewUnaryHandler(
		BackrestClearHistoryProcedure,
		svc.ClearHistory,
		connect.WithSchema(backrestClearHistoryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	backrestPathAutocompleteHandler := connect.NewUnaryHandler(
		BackrestPathAutocompleteProcedure,
		svc.PathAutocomplete,
		connect.WithSchema(backrestPathAutocompleteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/v1.Backrest/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BackrestGetConfigProcedure:
			backrestGetConfigHandler.ServeHTTP(w, r)
		case BackrestSetConfigProcedure:
			backrestSetConfigHandler.ServeHTTP(w, r)
		case BackrestAddRepoProcedure:
			backrestAddRepoHandler.ServeHTTP(w, r)
		case BackrestGetOperationEventsProcedure:
			backrestGetOperationEventsHandler.ServeHTTP(w, r)
		case BackrestGetOperationsProcedure:
			backrestGetOperationsHandler.ServeHTTP(w, r)
		case BackrestListSnapshotsProcedure:
			backrestListSnapshotsHandler.ServeHTTP(w, r)
		case BackrestListSnapshotFilesProcedure:
			backrestListSnapshotFilesHandler.ServeHTTP(w, r)
		case BackrestBackupProcedure:
			backrestBackupHandler.ServeHTTP(w, r)
		case BackrestPruneProcedure:
			backrestPruneHandler.ServeHTTP(w, r)
		case BackrestForgetProcedure:
			backrestForgetHandler.ServeHTTP(w, r)
		case BackrestRestoreProcedure:
			backrestRestoreHandler.ServeHTTP(w, r)
		case BackrestUnlockProcedure:
			backrestUnlockHandler.ServeHTTP(w, r)
		case BackrestCancelProcedure:
			backrestCancelHandler.ServeHTTP(w, r)
		case BackrestClearHistoryProcedure:
			backrestClearHistoryHandler.ServeHTTP(w, r)
		case BackrestPathAutocompleteProcedure:
			backrestPathAutocompleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBackrestHandler returns CodeUnimplemented from all methods.
type UnimplementedBackrestHandler struct{}

func (UnimplementedBackrestHandler) GetConfig(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.Config], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.GetConfig is not implemented"))
}

func (UnimplementedBackrestHandler) SetConfig(context.Context, *connect.Request[v1.Config]) (*connect.Response[v1.Config], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.SetConfig is not implemented"))
}

func (UnimplementedBackrestHandler) AddRepo(context.Context, *connect.Request[v1.Repo]) (*connect.Response[v1.Config], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.AddRepo is not implemented"))
}

func (UnimplementedBackrestHandler) GetOperationEvents(context.Context, *connect.Request[emptypb.Empty], *connect.ServerStream[v1.OperationEvent]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.GetOperationEvents is not implemented"))
}

func (UnimplementedBackrestHandler) GetOperations(context.Context, *connect.Request[v1.GetOperationsRequest]) (*connect.Response[v1.OperationList], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.GetOperations is not implemented"))
}

func (UnimplementedBackrestHandler) ListSnapshots(context.Context, *connect.Request[v1.ListSnapshotsRequest]) (*connect.Response[v1.ResticSnapshotList], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.ListSnapshots is not implemented"))
}

func (UnimplementedBackrestHandler) ListSnapshotFiles(context.Context, *connect.Request[v1.ListSnapshotFilesRequest]) (*connect.Response[v1.ListSnapshotFilesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.ListSnapshotFiles is not implemented"))
}

func (UnimplementedBackrestHandler) Backup(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.Backup is not implemented"))
}

func (UnimplementedBackrestHandler) Prune(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.Prune is not implemented"))
}

func (UnimplementedBackrestHandler) Forget(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.Forget is not implemented"))
}

func (UnimplementedBackrestHandler) Restore(context.Context, *connect.Request[v1.RestoreSnapshotRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.Restore is not implemented"))
}

func (UnimplementedBackrestHandler) Unlock(context.Context, *connect.Request[types.StringValue]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.Unlock is not implemented"))
}

func (UnimplementedBackrestHandler) Cancel(context.Context, *connect.Request[types.Int64Value]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.Cancel is not implemented"))
}

func (UnimplementedBackrestHandler) ClearHistory(context.Context, *connect.Request[v1.ClearHistoryRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.ClearHistory is not implemented"))
}

func (UnimplementedBackrestHandler) PathAutocomplete(context.Context, *connect.Request[types.StringValue]) (*connect.Response[types.StringList], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("v1.Backrest.PathAutocomplete is not implemented"))
}
