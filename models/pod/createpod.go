package pod

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"Slyph/models"
	"net/http"
)

func PodAdd(req *models.Pod) (int, error) {
	log := models.Log{}
	var pod = v1.Pod{}
	for index := range req.Containers {
		pod = v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      req.Name,
				Namespace: req.Namespace,
				Labels: map[string]string{
					"name": req.Name,
				},
			},
			Spec: v1.PodSpec{
				RestartPolicy: v1.RestartPolicyAlways,
				Containers: []v1.Container{
					v1.Container{
						Env:   req.Containers[index].Env,
						Name:  req.Containers[index].Name,
						Image: req.Containers[index].Image,
						Ports: []v1.ContainerPort{
							v1.ContainerPort{
								ContainerPort: req.Containers[index].ContainerPort,
							},
						},
					},
				},
			},
		}
	}
	if _, err := models.GetClientset().CoreV1().Pods(req.Namespace).Create(&pod); err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to create  pod " + req.Name + " " + err.Error()
		go models.Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Created pod " + req.Name + " success"
	go models.Insertlog(&log)
	return http.StatusOK, nil
}
