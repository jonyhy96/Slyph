package statefulset

import (
	"encoding/json"
	"k8s.io/api/apps/v1beta2"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s/models"
	"net/http"
)

// AddStatefulset add a Statefulset
func AddStatefulset(req *models.Statefulset) (int, error) {
	log := models.Log{}
	var r apiv1.ResourceRequirements
	j := `{"limits": {"cpu":"2000m", "memory": "2Gi"}, "requests": {"cpu":"1000m", "memory": "1Gi"}}`
	json.Unmarshal([]byte(j), &r)
	var podspec = apiv1.PodSpec{}
	for index := range req.Containers {
		podspec = apiv1.PodSpec{
			RestartPolicy: apiv1.RestartPolicyAlways,
			Containers: []apiv1.Container{
				apiv1.Container{
					Env:   req.Containers[index].Env,
					Name:  req.Containers[index].Name,
					Image: req.Containers[index].Image,
					Ports: []apiv1.ContainerPort{
						apiv1.ContainerPort{
							ContainerPort: req.Containers[index].ContainerPort,
						},
					},
					Resources: r,
					VolumeMounts: []apiv1.VolumeMount{
						apiv1.VolumeMount{
							Name:      req.VMname,
							MountPath: req.MountPath,
						},
					},
				},
			},
		}
	}
	st := v1beta2.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
			Labels: map[string]string{
				"name": req.Name,
			},
		},
	}
	vcts := apiv1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.VMname,
			Annotations: map[string]string{
				"volume.beta.kubernetes.io/storage-class": req.StorageName,
			},
			Labels: map[string]string{
				"name": req.Name,
			},
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.PersistentVolumeAccessMode(req.AccessModes),
			},
		},
	}
	requestStorage, err := resource.ParseQuantity(req.Requests)
	if err != nil {
		return 1, err
	}
	vcts.Spec.Resources.Requests = apiv1.ResourceList{apiv1.ResourceStorage: requestStorage}
	spec := v1beta2.StatefulSetSpec{
		Replicas: &req.Replicas,
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"name": req.Name,
			},
		},
		ServiceName: req.Name,
		Template: apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"name": req.Name,
				},
			},
		},
	}
	spec.VolumeClaimTemplates = append(spec.VolumeClaimTemplates, vcts)
	spec.Template.Spec = podspec
	st.Spec = spec
	// Create service
	if _, err := models.GetClientset().AppsV1beta2().StatefulSets(req.Namespace).Create(&st); err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Created StatefulSets " + req.Name + " error"
		go models.Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	log.Time = models.Gettime()
	log.Etype = "info"
	log.Event = "Created StatefulSets " + req.Name + " success"
	go models.Insertlog(&log)
	return http.StatusOK, nil
}
