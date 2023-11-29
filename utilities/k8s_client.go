package utilities

import (
	"flag"
	entandoclients "github.com/antromeo/entando-clients/pkg/client/clientset/versioned"
	pipelineclients "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sync"
)

type KubeClient struct {
	ClientSet         *kubernetes.Clientset
	EntandoClientSet  *entandoclients.Clientset
	PipelineClientSet *pipelineclients.Clientset
	DynamicClient     *dynamic.DynamicClient
	Namespace         string
}

var kubeClient *KubeClient
var onceK8s sync.Once

func GetKubeClientInstance() *KubeClient {
	onceK8s.Do(func() {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err)
		}
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		dynamicClient, err := dynamic.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		entandoclientset, err := entandoclients.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		pipelineclients, err := pipelineclients.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		clientCfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
		namespace := clientCfg.Contexts[clientCfg.CurrentContext].Namespace
		kubeClient = &KubeClient{
			ClientSet:         clientset,
			EntandoClientSet:  entandoclientset,
			PipelineClientSet: pipelineclients,
			DynamicClient:     dynamicClient,
			Namespace:         namespace,
		}

	})
	return kubeClient
}
