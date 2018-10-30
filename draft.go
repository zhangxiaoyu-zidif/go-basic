func (ssc *StatefulSetController) deleteStatefulSet(obj interface{}) {
	set, ok := obj.(*apps.StatefulSet)
	if ok {
		if flag, ok := set.Annotations[StatefulSetPVCAutoDelete]; ok && flag != "true" {
			namespace := set.Namespace
			if len(set.Spec.VolumeClaimTemplates) > 0 {
				selector, err := metav1.LabelSelectorAsSelector(set.Spec.Selector)
				if err != nil {
					utilruntime.HandleError(fmt.Errorf("error converting StatefulSet %v selector: %v", set, err))
					// This is a non-transient error, so don't retry.
					return
				}
	
				if err := ssc.adoptOrphanRevisions(set); err != nil {
					utilruntime.HandleError(err)
					return
				}
	
				pods, err := ssc.getPodsForStatefulSet(set, selector)
				if err != nil {
					utilruntime.HandleError(err)
					return
				}
	
				for _, pod := range pods {
					if len(pod.Spec.Volumes) > 0 {
						for _, volume := range pod.Spec.Volumes {
							if volume.PersistentVolumeClaim != nil {
								pvcName := volume.PersistentVolumeClaim.ClaimName
								err := ssc.kubeClient.CoreV1().PersistentVolumeClaims(namespace).Delete(pvcName, nil)
								if err != nil {
									pvcName := volume.PersistentVolumeClaim.ClaimName
									glog.Infof("deletePVC: %s/%s failed. error: %v.", namespace, pvcName, err)
								} else {
									glog.Infof("deletePVC: %s/%s successfully.", namespace, pvcName)
								}
							}
						}
					}
				}
			}
		}
	}
	ssc.enqueueStatefulSet(obj)
}
