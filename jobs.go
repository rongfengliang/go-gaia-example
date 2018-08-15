package main

import (
	sdk "github.com/gaia-pipeline/gosdk"
)

var jobs = sdk.Jobs{
	sdk.Job{
		Handler:     CreateUser,
		Title:       "Create DB User",
		Description: "Creates a database user with least privileged permissions.",
	},
	sdk.Job{
		Handler:     MigrateDB,
		Title:       "DB Migration",
		Description: "Imports newest test data dump and migrates to newest version.",
	},
	sdk.Job{
		Handler:     CreateNamespace,
		Title:       "Create K8S Namespace",
		Description: "Creates a new Kubernetes namespace for the new test environment.",
	},
	sdk.Job{
		Handler:     CreateDeployment,
		Title:       "Create K8S Deployment",
		Description: "Creates a new Kubernetes deployment for the new test environment.",
	},
	sdk.Job{
		Handler:     CreateService,
		Title:       "Create K8S Service",
		Description: "Creates a new Kubernetes service for the new test environment.",
	},
	sdk.Job{
		Handler:     CreateIngress,
		Title:       "Create K8S Ingress",
		Description: "Creates a new Kubernetes ingress for the new test environment.",
	},
	sdk.Job{
		Handler:     Cleanup,
		Title:       "Clean up",
		Description: "Removes all temporary files.",
	},
}
