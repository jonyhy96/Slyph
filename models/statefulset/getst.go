package statefulset

import (
	v1beta2 "k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s/models"
)

// GetStatefulset get the Statefulset info
func GetStatefulset() *v1beta2.StatefulSetList {
	log := models.Log{}
	sts, err := models.GetClientset().AppsV1beta2().StatefulSets("default").List(metav1.ListOptions{})
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to get Statefulset " + err.Error()
		go models.Insertlog(&log)
	}
	return sts
}
