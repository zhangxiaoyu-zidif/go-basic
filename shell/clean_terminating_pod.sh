#!/bin/bash

echo "execute delete terminating pod time:" >> /var/log/delete_pod_log
current=$(date +%Y-%m-%dT%T)
echo $current >> delete_pod_log

target_podname=$(/usr/bin/kubectl get pods -n appos | grep ' Terminating ' | awk '{print $1}')

count=0
for podname in ${target_podname[@]}
do
        deleteStamp=$(/usr/bin/kubectl get pod $podname -n appos -o yaml | grep deletionTimestamp | awk '{print $2}')
        deleteDate=$(date -d $deleteStamp +%s)
        currentTime=$(date -u +%s)
        timeMinus=$(echo "($currentTime-$deleteDate)/(60*60)"|bc)

        if [ $timeMinus -ge 1 ]; then
                /usr/bin/kubectl get pod $podname -n appos -o wide >> /var/log/delete_pod_log
                count=$(($count+1))
                if [ $count == 10 ]; then
                        exit 0
                fi
                # kubectl patch pod $podname --patch '{"metadata": {"$deleteFromPrimitiveList/finalizers":["pod.beta1.sigma.ali/cni-allocated"]}}' -n appos
        fi
done
