package internal

import (
	"context"
	errs "errors"
	"fmt"

	// "time"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

	"github.com/wtiger001/lair-backend/models"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/wtiger001/lair-backend/restapi/operations"
)

// StartSession - Starts a session for a given workspace
func StartSession(workspace *models.Workspace) error {
	fmt.Printf("Starting Session for %s \n", workspace.ID)

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Client Configured \n")

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Client Created \n")

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: sessionName(workspace.ID),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": sessionName(workspace.ID),
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": sessionName(workspace.ID),
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  sessionName(workspace.ID) + "-container",
							Image: workspace.DockerImage,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment for workspace ")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return err
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	return nil
}

// CheckSession - Checks the status of a session pod. First find the deployment and then find the pod...
func CheckSession(id string) (*apiv1.PodStatus, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Look for the deployment
	namespace := "default"
	// deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deploymentName := sessionName(id)
	// deployment, err := deploymentsClient.Get(context.TODO(), deploymentName,  metav1.GetOptions{})
	// if errors.IsNotFound(err) {
	// 	fmt.Printf("Deployment %s in namespace %s not found\n", deploymentName, namespace)
	// 	return nil, err
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	fmt.Printf("Error getting deployment %s in namespace %s: %v\n",
	// 	deploymentName, namespace, statusError.ErrStatus.Message)
	// 	return nil, err
	// } else if err != nil {
	// 	return nil, err
	// }

	result, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=" + deploymentName,
	})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", deploymentName, namespace)
		return nil, err
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			deploymentName, namespace, statusError.ErrStatus.Message)
		return nil, err
	} else if err != nil {
		return nil, err
	}

	if len(result.Items) == 1 {
		return &result.Items[0].Status, nil
	}

	return nil, errs.New(" PODS != 1")
}

// DeleteSession - Deletes a workspace session
func DeleteSession(id string) error {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	deploymentName := sessionName(id)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deletePolicy := metav1.DeletePropagationForeground
	err = deploymentsClient.Delete(context.TODO(), deploymentName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	return err
}

// GetWorkspaceLaunchStatus - Gets the session status
func GetWorkspaceLaunchStatus(params operations.GetWorkpaceLaunchStatusParams) middleware.Responder {
	fmt.Printf("Checking Session for %s \n", params.ID)

	status, err := CheckSession(params.ID)

	if err != nil {
		fmt.Printf("Workspace NOT FOUND: %s \n", params.ID)

		return operations.NewGetWorkpaceLaunchStatusNotFound()
	}

	rtn := &models.WorkspaceStatus{
		ID:     params.ID,
		Status: phaseToStr(status.Phase),
		URL:    "http://192.168.0.182:8080",
	}

	return operations.NewGetWorkpaceLaunchStatusOK().WithPayload(rtn)

}

// LaunchWorkspace - Launches a Workspace
func LaunchWorkspace(params operations.LaunchWorkpaceByIDParams) middleware.Responder {
	// Find the workspace
	fmt.Printf("Looking for workspace %s \n", params.ID)

	workspace, err := DS.Retrieve(params.ID)

	if err != nil {
		fmt.Printf("Error : %s \n", err)
		return operations.NewLaunchWorkpaceByIDInternalServerError()
	}
	if workspace == nil {
		fmt.Printf("Workspace NOT FOUND: %s \n", params.ID)
		return operations.NewLaunchWorkpaceByIDNotFound()
	}

	// Workspace found
	err = StartSession(workspace)
	if err != nil {

		fmt.Printf("Error Starting Session: %s \n", err)

		return operations.NewLaunchWorkpaceByIDInternalServerError().WithPayload(err.Error())
	}

	return operations.NewLaunchWorkpaceByIDOK()
}

// TerminateSession - Terminates a session
func TerminateSession(params operations.CancelLaunchParams) middleware.Responder {
	// Find the workspace
	fmt.Printf("Looking for workspace %s \n", params.ID)

	workspace, err := DS.Retrieve(params.ID)

	if err != nil {
		fmt.Printf("Error : %s \n", err)
		return operations.NewCancelLaunchNotFound()
	}
	if workspace == nil {
		fmt.Printf("Workspace NOT FOUND: %s \n", params.ID)
		return operations.NewCancelLaunchNotFound()
	}

	// Workspace found
	err = DeleteSession(workspace.ID)
	if err != nil {
		fmt.Printf("Error Canceling Session: %s \n", err)
		return operations.NewCancelLaunchInternalServerError().WithPayload(err.Error())
	}

	return operations.NewCancelLaunchOK()
}

func int32Ptr(i int32) *int32 { return &i }

func phaseToStr(phase apiv1.PodPhase) string {
	switch phase {
	case apiv1.PodPending:
		return models.WorkspaceStatusStatusScheduling
	case apiv1.PodRunning:
		return models.WorkspaceStatusStatusRunning
	case apiv1.PodSucceeded:
		return models.WorkspaceStatusStatusStopping
	case apiv1.PodFailed:
		return models.WorkspaceStatusStatusStopping
	case apiv1.PodUnknown:
		return models.WorkspaceStatusStatusStopping
	}
	return "ERROR--NO STATUS"
}

func sessionName(ID string) string {
	return "workspace-session-" + ID
}
