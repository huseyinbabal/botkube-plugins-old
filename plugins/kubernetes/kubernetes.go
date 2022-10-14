package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/huseyinbabal/botkube-plugins/internal"
	botkubeplugin "github.com/huseyinbabal/botkube-plugins/internal/source"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"
)

type KubernetesSource struct{}

func (p KubernetesSource) Consume(ch chan interface{}) error {
	events := make(chan interface{}, 1)
	go listenEvents(events)

	for {
		select {
		case event := <-events:
			ch <- event
		}
	}
	return nil
}

func listenEvents(ch chan interface{}) {
	home, _ := os.UserHomeDir()
	config, _ := clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
	clientset, _ := kubernetes.NewForConfig(config)

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(clientset, time.Second*30)
	svcInformer := kubeInformerFactory.Core().V1().Pods().Informer()

	svcInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ch <- obj
		},
		DeleteFunc: func(obj interface{}) {
			ch <- obj
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			ch <- newObj
		},
	})

	stop := make(chan struct{})
	defer close(stop)
	kubeInformerFactory.Start(stop)
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	internal.Serve(map[string]plugin.Plugin{
		"kubernetes": &botkubeplugin.SourcePlugin{Impl: &KubernetesSource{}},
	})
}
