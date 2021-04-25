package kube

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getPodLabel(namespace, podName string) {
	clientSet := GetKubeClient()
	pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(pod.GetLabels())
}

func listPod(namespace string, labels map[string]string) {
	clientSet := GetKubeClient()
	selector, err := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
		MatchLabels: labels,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	podList, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: selector.String(),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, pod := range podList.Items {
		log.Println(pod.GetLabels())
	}
}
