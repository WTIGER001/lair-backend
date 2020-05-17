// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/wtiger001/lair-backend/internal"
	"github.com/wtiger001/lair-backend/restapi/operations"
)

//go:generate swagger generate server --target ../../lair-backend --name Lair --spec ../../../../../../swagger/swagger.yaml

func configureFlags(api *operations.LairAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LairAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Configure the data store
	internal.DS = internal.NewDataStore()

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetWorkspacesHandler = operations.GetWorkspacesHandlerFunc(internal.GetWorkspaces)
	api.GetWorkspacesIDHandler = operations.GetWorkspacesIDHandlerFunc(internal.GetWorkspaceByID)
	api.PostWorkpaceByIDHandler = operations.PostWorkpaceByIDHandlerFunc(internal.SaveWorkspace)
	
	if api.PutWorkpaceByIDHandler == nil {
		api.PutWorkpaceByIDHandler = operations.PutWorkpaceByIDHandlerFunc(func(params operations.PutWorkpaceByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PutWorkpaceByID has not yet been implemented")
		})
	}

	if api.CancelLaunchHandler == nil {
		api.CancelLaunchHandler = operations.CancelLaunchHandlerFunc(func(params operations.CancelLaunchParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CancelLaunch has not yet been implemented")
		})
	}
	if api.DeleteWorkpaceByIDHandler == nil {
		api.DeleteWorkpaceByIDHandler = operations.DeleteWorkpaceByIDHandlerFunc(func(params operations.DeleteWorkpaceByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteWorkpaceByID has not yet been implemented")
		})
	}
	if api.GetWorkpaceLaunchStatusHandler == nil {
		api.GetWorkpaceLaunchStatusHandler = operations.GetWorkpaceLaunchStatusHandlerFunc(func(params operations.GetWorkpaceLaunchStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetWorkpaceLaunchStatus has not yet been implemented")
		})
	}
	if api.LaunchWorkpaceByIDHandler == nil {
		api.LaunchWorkpaceByIDHandler = operations.LaunchWorkpaceByIDHandlerFunc(func(params operations.LaunchWorkpaceByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.LaunchWorkpaceByID has not yet been implemented")
		})
	}
	

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
