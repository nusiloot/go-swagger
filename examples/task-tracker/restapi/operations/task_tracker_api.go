package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-swagger/go-swagger/examples/task-tracker/restapi/operations/tasks"
)

// NewTaskTrackerAPI creates a new TaskTracker instance
func NewTaskTrackerAPI(spec *loads.Document) *TaskTrackerAPI {
	return &TaskTrackerAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
	}
}

/*TaskTrackerAPI This application implements a very simple issue tracker.
It's implemented as an API which is described by this swagger spec document.

The go-swagger project uses this specification to test the code generation.
This document contains all possible values for a swagger definition.
This means that it exercises the framework relatively well.
*/
type TaskTrackerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/vnd.goswagger.examples.task-tracker.v1+json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/vnd.goswagger.examples.task-tracker.v1+json" mime type
	JSONProducer runtime.Producer

	// TokenHeaderAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-Token provided in the header
	TokenHeaderAuth func(string) (interface{}, error)

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key token provided in the query
	APIKeyAuth func(string) (interface{}, error)

	// TasksAddCommentToTaskHandler sets the operation handler for the add comment to task operation
	TasksAddCommentToTaskHandler tasks.AddCommentToTaskHandler
	// TasksCreateTaskHandler sets the operation handler for the create task operation
	TasksCreateTaskHandler tasks.CreateTaskHandler
	// TasksDeleteTaskHandler sets the operation handler for the delete task operation
	TasksDeleteTaskHandler tasks.DeleteTaskHandler
	// TasksGetTaskCommentsHandler sets the operation handler for the get task comments operation
	TasksGetTaskCommentsHandler tasks.GetTaskCommentsHandler
	// TasksGetTaskDetailsHandler sets the operation handler for the get task details operation
	TasksGetTaskDetailsHandler tasks.GetTaskDetailsHandler
	// TasksListTasksHandler sets the operation handler for the list tasks operation
	TasksListTasksHandler tasks.ListTasksHandler
	// TasksUpdateTaskHandler sets the operation handler for the update task operation
	TasksUpdateTaskHandler tasks.UpdateTaskHandler
	// TasksUploadTaskFileHandler sets the operation handler for the upload task file operation
	TasksUploadTaskFileHandler tasks.UploadTaskFileHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TaskTrackerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TaskTrackerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TaskTrackerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TaskTrackerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TaskTrackerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TaskTrackerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TaskTrackerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TaskTrackerAPI
func (o *TaskTrackerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.TokenHeaderAuth == nil {
		unregistered = append(unregistered, "XTokenAuth")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "TokenAuth")
	}

	if o.TasksAddCommentToTaskHandler == nil {
		unregistered = append(unregistered, "tasks.AddCommentToTaskHandler")
	}

	if o.TasksCreateTaskHandler == nil {
		unregistered = append(unregistered, "tasks.CreateTaskHandler")
	}

	if o.TasksDeleteTaskHandler == nil {
		unregistered = append(unregistered, "tasks.DeleteTaskHandler")
	}

	if o.TasksGetTaskCommentsHandler == nil {
		unregistered = append(unregistered, "tasks.GetTaskCommentsHandler")
	}

	if o.TasksGetTaskDetailsHandler == nil {
		unregistered = append(unregistered, "tasks.GetTaskDetailsHandler")
	}

	if o.TasksListTasksHandler == nil {
		unregistered = append(unregistered, "tasks.ListTasksHandler")
	}

	if o.TasksUpdateTaskHandler == nil {
		unregistered = append(unregistered, "tasks.UpdateTaskHandler")
	}

	if o.TasksUploadTaskFileHandler == nil {
		unregistered = append(unregistered, "tasks.UploadTaskFileHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TaskTrackerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TaskTrackerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "token_header":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, o.TokenHeaderAuth)

		case "api_key":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, o.APIKeyAuth)

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *TaskTrackerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/vnd.goswagger.examples.task-tracker.v1+json":
			result["application/vnd.goswagger.examples.task-tracker.v1+json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TaskTrackerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/vnd.goswagger.examples.task-tracker.v1+json":
			result["application/vnd.goswagger.examples.task-tracker.v1+json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TaskTrackerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the task tracker API
func (o *TaskTrackerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TaskTrackerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tasks/{id}/comments"] = tasks.NewAddCommentToTask(o.context, o.TasksAddCommentToTaskHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tasks"] = tasks.NewCreateTask(o.context, o.TasksCreateTaskHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tasks/{id}"] = tasks.NewDeleteTask(o.context, o.TasksDeleteTaskHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks/{id}/comments"] = tasks.NewGetTaskComments(o.context, o.TasksGetTaskCommentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks/{id}"] = tasks.NewGetTaskDetails(o.context, o.TasksGetTaskDetailsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks"] = tasks.NewListTasks(o.context, o.TasksListTasksHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/tasks/{id}"] = tasks.NewUpdateTask(o.context, o.TasksUpdateTaskHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tasks/{id}/files"] = tasks.NewUploadTaskFile(o.context, o.TasksUploadTaskFileHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TaskTrackerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

func (o *TaskTrackerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
