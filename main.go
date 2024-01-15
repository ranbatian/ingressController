// main 包声明
package main

// 导入必要的包
import (
	"github.com/simon/ingressController/demo/pkg" // 导入自定义包
	"k8s.io/client-go/informers"                  // 导入 Kubernetes 客户端库中的 informers 包
	"k8s.io/client-go/kubernetes"                 // 导入 Kubernetes 客户端库
	"k8s.io/client-go/tools/clientcmd"            // 导入 Kubernetes 客户端工具包
)

// main 函数，程序入口
func main() {
	// 从用户的 kubeconfig 文件中构建 Kubernetes 配置
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic("failed to create client config.") // 处理错误情况
	}

	// 根据配置创建 Kubernetes 客户端
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("failed to create client.") // 处理错误情况
	}

	// 创建 shared informer 工厂，用于监视 Kubernetes 资源变更
	factory := informers.NewSharedInformerFactory(clientSet, 0)

	// 创建 Service 资源的 informer
	serviceInformer := factory.Core().V1().Services()

	// 创建 Ingress 资源的 informer
	ingressInformer := factory.Networking().V1().Ingresses()

	// 创建自定义控制器对象，传入 Kubernetes 客户端和 informers
	c := pkg.NewController(clientSet, serviceInformer, ingressInformer)

	// 创建用于停止 informer 的通道
	stopChan := make(chan struct{})

	// 启动 informer 工厂
	factory.Start(stopChan)

	// 等待缓存同步完成
	factory.WaitForCacheSync(stopChan)

	// 运行自定义控制器
	c.Run(stopChan)
}
