// 包声明
package pkg

// 导入必要的包
import (
	"fmt"
	"reflect"

	v1 "k8s.io/api/networking/v1"
	informers "k8s.io/client-go/informers/core/v1"          // 导入 Kubernetes 客户端 informers 包
	netInformers "k8s.io/client-go/informers/networking/v1" // 导入 Kubernetes 网络 informers 包
	"k8s.io/client-go/kubernetes"                           // 导入 Kubernetes 客户端库
	core "k8s.io/client-go/listers/core/v1"                 // 导入 Kubernetes 核心资源列表包
	network "k8s.io/client-go/listers/networking/v1"        // 导入 Kubernetes 网络资源列表包
	"k8s.io/client-go/tools/cache"                          // 导入 Kubernetes 缓存工具包
	"k8s.io/client-go/util/workqueue"
)

// 控制器结构体
type controller struct {
	client        *kubernetes.Clientset           // Kubernetes 客户端
	serviceLister core.ServiceLister              // Service 资源列表器
	ingressList   network.IngressLister           // Ingress 资源列表器
	queue         workqueue.RateLimitingInterface // 用于处理工作队列的接口
}

// 更新 Service 资源的回调函数
func (c *controller) updateService(oldObj interface{}, newObj interface{}) {
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	c.enQueue(newObj)
}

// 添加 Service 资源的回调函数
func (c *controller) addService(obj interface{}) {
	c.enQueue(obj)
}

// 删除 Ingress 资源的回调函数
func (c *controller) deleteIngress(obj interface{}) {
	i, ok := obj.(*v1.Ingress)
	if !ok {
		return
	}
	c.queue.Add(i.Namespace + "/" + i.Name)
}

// 将对象添加到工作队列的函数
func (c *controller) enQueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		fmt.Printf("error creating key: %v\n", err)
	}

	c.queue.Add(key)
}

// 运行控制器的主函数
func (c *controller) Run(stopChan chan struct{}) {
	<-stopChan
}

// NewController 函数，用于创建新的控制器实例
func NewController(client *kubernetes.Clientset, service informers.ServiceInformer, ingress netInformers.IngressInformer) controller {
	// 创建控制器对象
	c := controller{
		client:        client,
		serviceLister: service.Lister(),
		ingressList:   ingress.Lister(),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingressManager"),
	}

	// 增加 Service 资源的监听事件
	service.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})

	// 增加 Ingress 资源的监听事件
	ingress.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})

	return c // 返回创建的控制器对象
}
