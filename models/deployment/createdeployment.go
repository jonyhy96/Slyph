package deployment

import (
	"encoding/json"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s/models"
	"net/http"
)

// AddDeployment adds a deployment
func AddDeployment(req *models.Deployment) (int, error) {
	log := models.Log{}
	var r apiv1.ResourceRequirements
	j := `{"limits": {"cpu":"2000m", "memory": "2Gi"}, "requests": {"cpu":"1000m", "memory": "1Gi"}}`
	json.Unmarshal([]byte(j), &r)
	deploymentsClient := models.GetClientset().AppsV1beta1().Deployments(req.Namespace)
	deployment := &appsv1beta1.Deployment{}
	for index := range req.Containers {
		deployment = &appsv1beta1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: req.Name,
			},
			Spec: appsv1beta1.DeploymentSpec{
				Replicas: int32Ptr(req.Replicas),
				Template: apiv1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"name": req.Containers[0].Name,
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							{
								Env:   req.Containers[index].Env,
								Name:  req.Containers[index].Name,
								Image: req.Containers[index].Image,
								Ports: []apiv1.ContainerPort{
									{
										Name:          "http",
										Protocol:      apiv1.ProtocolTCP,
										ContainerPort: req.Containers[0].ContainerPort,
									},
								},
								Resources: r,
							},
						},
					},
				},
			},
		}
	}
	// Create Deployment
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to create  deployment " + result.GetObjectMeta().GetName() + " " + err.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	models.Insertbaseimage(req.Name, req.Containers[0].Image, "base")
	go models.Insertbaseimage(req.Name, req.Containers[0].Image, "new")
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Created deployment " + result.GetObjectMeta().GetName() + " success"
	go models.Insertlog(&log)
	return http.StatusOK, nil
}
func int32Ptr(i int32) *int32 { return &i }
