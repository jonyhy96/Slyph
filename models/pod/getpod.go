package pod

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"Slyph/models"
)

func Getpodinfo() *v1.PodList {
	log := new(models.Log)
	pods, err := models.GetClientset().CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Get pod error " + err.Error()
		models.Insertlog(log)
		return nil
	}
	return pods
}
