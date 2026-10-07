package template

import (
	"fmt"
	"net/http"

	//	"strings"

	"go.uber.org/zap"

	"github.com/wundergraph/cosmo/router/core"
	"github.com/wundergraph/cosmo/router/pkg/pubsub/datasource"
)

// This is the name that the module will be known by in the router and must be unique. It does not have to match the package name.
const moduleID = "moduleTemplate"

// This is the priority for the module to be loaded among all other modules. Lower priorities are loaded first. If the priority is the
// same for multiple modules, the load order is non-deterministic among those modules.
const loadPriority = 1

// You do not need to change the type name here for additional modules.
type RouterModule struct {
	// Declare variables here -->
	// Example bool `mapstructure:"example_setting"`
	// Note that variables must be exportable (i.e., PascalCase) to be set.

	Logger *zap.Logger
}

// Module returns registration info with unique ID and execution priority
func (m *RouterModule) Module() core.ModuleInfo {
	return core.ModuleInfo{
		ID:       moduleID,
		Priority: loadPriority,
		New: func() core.Module {
			return &RouterModule{}
		},
	}
}

// Provision implements [core.Provisioner].
func (m *RouterModule) Provision(ctx *core.ModuleContext) error {
	m.Logger = ctx.Logger

	// Add your custom initialization, config validation, etc. code here.
	// This hook is called once on router startup and returning an error causes the router to fail to start.

	m.Logger.Info(fmt.Sprintf("custom module %v loaded", moduleID))

	return nil
}

// Cleanup implements [core.Cleaner].
func (m *RouterModule) Cleanup() error {
	// Place your cleanup code here that runs on router shutdown
	return nil
}

// RouterOnRequest implements [core.RouterOnRequestHandler].
func (m *RouterModule) RouterOnRequest(ctx core.RequestContext, next http.Handler) {
	// Place code here that runs as soon as a request is received.
	// Code here runs once per request, prior to authentication processing.

	// DO NOT REMOVE this line, as it's necessary to continue the request.
	next.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

// Middleware implements [core.RouterMiddlewareHandler].
func (m *RouterModule) Middleware(ctx core.RequestContext, next http.Handler) {
	// Place middleware code here that runs prior to GraphQL query plan execution.
	// Code here runs once per request, after authentication but before plan execution.

	// DO NOT REMOVE this line, as it's necessary to run any other middleware.
	next.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

// OnOriginRequest implements [core.EnginePreOriginHandler].
func (m *RouterModule) OnOriginRequest(req *http.Request, ctx core.RequestContext) (*http.Request, *http.Response) {
	// Place code here that runs before every subgraph request. Every request runs this hook in parallel and
	// in a nondeterministic order. You can use this to rewrite or add headers to the request, as well
	// as to short-circuit the response.

	// Returning a response short-circuits execution and is useful for returning custom responses.
	return req, nil
}

// OnOriginResponse implements [core.EnginePostOriginHandler].
func (m *RouterModule) OnOriginResponse(resp *http.Response, ctx core.RequestContext) *http.Response {
	// Place code here that runs after every subgraph request. Every request runs this hook in parallel and
	// in a nondeterministic order.

	// Returning a response short-circuits execution and is useful for returning custom responses.
	return nil
}

// Cosmo Streams module hooks
// SubscriptionOnCreate implements [core.SubscriptionOnCreateHandler].
func (m *RouterModule) SubscriptionOnCreate(ctx core.SubscriptionOnCreateHandlerContext) error {
	// Place code here that runs when a new subscription is requested, but before it's registered. This hook
	// runs before SubscriptionOnStart and can modify the event sourcing configuration, such as connecting
	// to a per-user topic.

	// Returning an error generates a GraphQL error and fails the subscription request.
	// use return &core.StreamHandlerError{} to return an error that is logged.
	return nil
}

// SubscriptionOnStart implements [core.SubscriptionOnStartHandler].
func (m *RouterModule) SubscriptionOnStart(ctx core.SubscriptionOnStartHandlerContext) error {
	// Place code here that runs when a new subscription is registered, but before events are generated.

	// Returning an error generates a GraphQL error and fails the subscription request.
	// use return &core.StreamHandlerError{} to return an error that is logged.
	return nil
}

// OnPublishEvents implements [core.StreamPublishEventHandler].
func (m *RouterModule) OnPublishEvents(ctx core.StreamPublishEventHandlerContext, events datasource.StreamEvents) (datasource.StreamEvents, error) {
	// Place hode here that runs prior to events being sent to event providers. This code runs once per provider, and the events object
	// is the set of events to be sent to the provider.

	// Returning an error here WILL send all events in the events set to the provider, unlike errors set in BeforeEventsDispatch.
	// The error is logged by the router and the client receives an edfs__PublishResult response with success = false.
	// If you want to return no events in the event of an error, return datasource.NewStreamEvents(nil).
	return events, nil
}

// BeforeEventsDispatch implements [core.StreamBeforeEventsDispatchHandler].
func (m *RouterModule) BeforeEventsDispatch(ctx core.StreamBeforeEventsDispatchHandlerContext, events datasource.StreamEvents) (datasource.StreamEvents, error) {
	// Place code here that runs when a batch of events is received from the message provider. The router reads events as a batch
	// and this code runs once per batch, as opposed to once per subscriber. A batch can contain 1 or more messages.
	// This hook only runs for messages to be delivered to subscribers. It does NOT run for events being sent to providers.

	// Returning an error prevents the entire batch from being processed. The error is *not* sent to any subscribed clients.
	// An error is logged by the router and no events from the batch are delivered to *any* subscriber. If an error is recoverable,
	// or all events are filtered out, return an empty set from datasource.NewStreamEvents(nil) and nil for error.
	return events, nil
}

// OnReceiveEvents implements [core.StreamReceiveEventHandler].
func (m *RouterModule) OnReceiveEvents(ctx core.StreamReceiveEventHandlerContext, events datasource.StreamEvents) (datasource.StreamEvents, error) {
	// Place code here that runs prior to events being sent to subscribers. This hook runs once per subscribed client, in parallel, with
	// a nondeterministic order.

	// Returning an error DOES send any events in the events set to the client. It does NOT send an error to the client, but instead
	// ends the subscription and closes the connection. The error is logged by the router.
	// If you don't want to send any events due to the error, return an empty set from datasource.NewStreamEvents(nil).
	return events, nil
}

// Interface guard
var (
	_ core.Provisioner             = (*RouterModule)(nil)
	_ core.Cleaner                 = (*RouterModule)(nil)
	_ core.RouterMiddlewareHandler = (*RouterModule)(nil)
	_ core.RouterOnRequestHandler  = (*RouterModule)(nil)
	_ core.EnginePreOriginHandler  = (*RouterModule)(nil)
	_ core.EnginePostOriginHandler = (*RouterModule)(nil)

	// Cosmo Streams (EDFS) handlers
	_ core.SubscriptionOnCreateHandler       = (*RouterModule)(nil)
	_ core.SubscriptionOnStartHandler        = (*RouterModule)(nil)
	_ core.StreamBeforeEventsDispatchHandler = (*RouterModule)(nil)
	_ core.StreamReceiveEventHandler         = (*RouterModule)(nil)
	_ core.StreamPublishEventHandler         = (*RouterModule)(nil)
)

func init() {
	// Register your module here
	core.RegisterModule(&RouterModule{})
}
