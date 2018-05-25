package models

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func CreatePV(req *PV) (int, error) {
	log := Log{}
	quantity, err := resource.ParseQuantity(req.StorageQuantity)
	if err != nil {
		log.Time = Gettime()
		log.Etype = "error"
		log.Event = "Failed to ParseQuantity for PV " + req.Name + " " + err.Error()
		go Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	pv := &apiv1.PersistentVolume{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolume",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
			Annotations: map[string]string{
				"volume.beta.kubernetes.io/storage-class": req.StorageName,
			},
			Labels: map[string]string{
				"storage_name": req.StorageName,
			},
		},
		Spec: apiv1.PersistentVolumeSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.PersistentVolumeAccessMode(req.AccessModes),
			},
			// PersistentVolumeReclaimPolicy: apiv1.PersistentVolumeReclaimDelete,
			PersistentVolumeReclaimPolicy: apiv1.PersistentVolumeReclaimPolicy("Recycle"),
			Capacity: apiv1.ResourceList{
				"storage": quantity,
			},
			PersistentVolumeSource: apiv1.PersistentVolumeSource{
				NFS: &apiv1.NFSVolumeSource{
					Server: req.NfsServer,
					Path:   req.NfsPath,
				},
			},
		},
	}
	client := GetClientset()
	if _, err = client.CoreV1().PersistentVolumes().Create(pv); err != nil {
		log.Time = Gettime()
		log.Etype = "error"
		log.Event = "Failed to create  PV " + req.Name + " " + err.Error()
		go Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	log.Time = Gettime()
	log.Etype = "info"
	log.Event = "Created pv " + req.Name + " success"
	go Insertlog(&log)
	return http.StatusOK, nil
}

func DeletePV(req *Delete) (int, error) {
	log := Log{}
	client := GetClientset()
	if err := client.CoreV1().PersistentVolumes().Delete(req.Name, &metav1.DeleteOptions{}); err != nil {
		log.Time = Gettime()
		log.Etype = "error"
		log.Event = "Failed to Delete for PV " + req.Name + " " + err.Error()
		go Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	log.Time = Gettime()
	log.Etype = "info"
	log.Event = "Delete PV " + req.Name + "success"
	return http.StatusOK, nil
}

func GetPV() *apiv1.PersistentVolumeList {
	log := Log{}
	client := GetClientset()
	var pvs *apiv1.PersistentVolumeList
	var err error
	if pvs, err = client.CoreV1().PersistentVolumes().List(metav1.ListOptions{}); err != nil {
		log.Time = Gettime()
		log.Etype = "error"
		log.Event = "Failed to Get list of PV " + " " + err.Error()
		go Insertlog(&log)
		return nil
	}
	return pvs
}

func CreatePVC(req *PVC) (int, error) {
	log := Log{}
	quantity, err := resource.ParseQuantity(req.StorageQuantity)
	if err != nil {
		log.Time = Gettime()
		log.Etype = "error"
		log.Event = "Failed to ParseQuantity for PVC " + req.Name + " " + err.Error()
		go Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	pvc := &apiv1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
			Annotations: map[string]string{
				"volume.beta.kubernetes.io/storage-class": req.StorageName,
			},
			Labels: map[string]string{
				"storage_name": req.StorageName,
			},
		},
		Spec: apiv1.PersistentVolumeClaimSpec{
			AccessModes: []apiv1.PersistentVolumeAccessMode{
				apiv1.PersistentVolumeAccessMode(req.AccessModes),
			},
			VolumeName: req.Name,
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{
					"storage": quantity,
				},
			},
		},
	}
	client := GetClientset()
	if _, err = client.CoreV1().PersistentVolumeClaims(req.Namespace).Create(pvc); err != nil {
		log.Time = Gettime()
		log.Etype = "error"
		log.Event = "Failed to create  PVC " + req.Name + " " + err.Error()
		go Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	log.Time = Gettime()
	log.Etype = "info"
	log.Event = "Created pvc " + req.Name + " success"
	go Insertlog(&log)
	return http.StatusOK, nil
}
func DeletePVC(req *Delete) (int, error) {
	log := Log{}
	client := GetClientset()
	if err := client.CoreV1().PersistentVolumeClaims(req.Namespace).Delete(req.Name, &metav1.DeleteOptions{}); err != nil {
		log.Time = Gettime()
		log.Etype = "error"
		log.Event = "Failed to Delete for PVC " + req.Name + " " + err.Error()
		go Insertlog(&log)
		return http.StatusInternalServerError, err
	}
	log.Time = Gettime()
	log.Etype = "info"
	log.Event = "Delete PVC " + req.Name + "success"
	go Insertlog(&log)
	return http.StatusOK, nil
}

func GetPVC() *apiv1.PersistentVolumeClaimList {
	log := Log{}
	client := GetClientset()
	var pvcs *apiv1.PersistentVolumeClaimList
	var err error
	if pvcs, err = client.CoreV1().PersistentVolumeClaims("default").List(metav1.ListOptions{}); err != nil {
		log.Time = Gettime()
		log.Etype = "error"
		log.Event = "Failed to Get list of PVC " + " " + err.Error()
		go Insertlog(&log)
		return nil
	}
	return pvcs
}
