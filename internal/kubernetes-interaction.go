package internal

import (
	"context"
	errs "errors"
	"fmt"

	// "time"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	networkvibeta "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	// "encoding/json"

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
func StartSession(workspace *models.Workspace) (*appsv1.Deployment, error) {
	fmt.Printf("Starting Session for %s \n", workspace.ID)

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Client Configured \n")

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Client Created \n")

	deployment := makeDeployment(workspace)
	service := makeService(workspace)
	ingress := makeIngress(workspace)

	// Create Deployment
	fmt.Println("Creating deployment for workspace ")
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return nil, err
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	// Create Services
	fmt.Println("Creating service for workspace ")
	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	resultSvc, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return nil, err
	}
	fmt.Printf("Created service %q.\n", resultSvc.GetObjectMeta().GetName())

	// Create Ingress
	fmt.Println("Creating service for workspace ")
	ingressClient := clientset.NetworkingV1beta1().Ingresses(apiv1.NamespaceDefault)
	resultIngress, err := ingressClient.Create(context.TODO(), ingress, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return nil, err
	}
	fmt.Printf("Created ingress %q.\n", resultIngress.GetObjectMeta().GetName())

	return result, nil
}

func makeService(workspace *models.Workspace) *apiv1.Service {
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: sessionName(workspace.ID) + "-service",
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": sessionName(workspace.ID),
			},
			Ports: []apiv1.ServicePort{
				{
					Name:       "http",
					Protocol:   apiv1.ProtocolTCP,
					Port:       80,
					TargetPort: intstr.FromInt(8080),
				},
			},
		},
	}

	return service
}

func makeIngress(workspace *models.Workspace) *networkvibeta.Ingress {

	sname := sessionName(workspace.ID)

	prefix := networkvibeta.PathType("Prefix")
	path := networkvibeta.HTTPIngressPath{
		Path:     sname,
		PathType: &prefix,
		Backend: networkvibeta.IngressBackend{
			ServiceName: "/" + sname + "-service",
			ServicePort: intstr.FromInt(80),
		},
	}

	ruleValue := networkvibeta.HTTPIngressRuleValue{
		Paths: []networkvibeta.HTTPIngressPath{
			path,
		},
	}

	rule := networkvibeta.IngressRule{}
	rule.HTTP = &ruleValue

	ingress := &networkvibeta.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: sname + "-ingress",
			Annotations: map[string]string{
				"traefik.ingress.kubernetes.io/rewrite-target": "/",
			},
		},
		Spec: networkvibeta.IngressSpec{
			Rules: []networkvibeta.IngressRule{
				rule,
			},
		},
	}
	return ingress
}

func makeDeployment(workspace *models.Workspace) *appsv1.Deployment {
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
									ContainerPort: 8080,
								},
							},
							Args: []string{
								"--auth", "none",
							},
						},
					},
				},
			},
		},
	}
	return deployment
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

	result, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=" + deploymentName,
	})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", deploymentName, namespace)
		return nil, nil
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

	return nil, fmt.Errorf("NO PODS")
}

// DeleteSession - Deletes a workspace session
func DeleteSession(id string) error {
	fmt.Printf("Starting Delete for Workspace %s \n", id)
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

	// Delete Services
	fmt.Println("Deleting service for workspace ")
	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	err = serviceClient.Delete(context.TODO(), sessionName(id) + "-service", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}
	fmt.Printf("Deleted service.\n", )

	// Create Ingress
	fmt.Println("Creating service for workspace ")
	ingressClient := clientset.NetworkingV1beta1().Ingresses(apiv1.NamespaceDefault)
	err = ingressClient.Delete(context.TODO(), sessionName(id) + "-ingress", metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}
	fmt.Printf("Deleted ingress.\n")

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

	// Launch and connect
	status, err := LaunchAndConnect(params.ID)

	if err != nil {
		fmt.Printf("Error : %s \n", err)
		return operations.NewLaunchWorkpaceByIDInternalServerError()
	}

	rtn := &models.WorkspaceStatus{
		ID:     params.ID,
		Status: phaseToStr(status.Phase),
		URL:    "http://192.168.0.182:8080",
	}

	return operations.NewLaunchWorkpaceByIDOK().WithPayload(rtn)
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

// LaunchAndConnect = Launches a workspace and gets the status
func LaunchAndConnect(ID string) (*apiv1.PodStatus, error) {
	// First get the status. If there is a status then just return that status
	fmt.Printf("Checking for Session %s \n", ID)
	status, err := CheckSession(ID)

	// An error was encountered. Likely something wrong in Kubernetes or the RBAC Setup
	// if err != nil && err {
	// 	return nil, err
	// }

	if status != nil {
		return status, nil
	}

	// Find the workspace
	fmt.Printf("Looking for workspace %s \n", ID)
	workspace, err := DS.Retrieve(ID)

	// An error was encountered. Likely something with the datastore
	if err != nil {
		return nil, err
	}
	if workspace == nil {
		fmt.Printf("Workspace NOT FOUND: %s \n", ID)
		return nil, errs.New("No Workspace Found")
	}

	_, err = StartSession(workspace)
	if err != nil {
		fmt.Printf("Error Starting Session: %s \n", err)
		return nil, err
	}

	// First get the status. If there is a status then just return that status
	status, err = CheckSession(ID)

	// An error was encountered. Likely something wrong in Kubernetes or the RBAC Setup
	if err != nil {
		return nil, err
	}

	return status, nil
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
