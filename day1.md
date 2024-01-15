main.go 逻辑描述：
导入包： 导入了程序所需的包，包括自定义包 github.com/simon/ingressController/demo/pkg 和 Kubernetes 客户端库。

构建 Kubernetes 配置： 使用 clientcmd.BuildConfigFromFlags 从用户的 kubeconfig 文件中构建 Kubernetes 配置。

创建 Kubernetes 客户端： 使用构建的配置创建 Kubernetes 客户端。

创建 Shared Informer 工厂： 使用客户端创建一个 Shared Informer 工厂，该工厂用于监视 Kubernetes 资源变更。

创建 Service 和 Ingress Informers： 使用工厂创建 Service 和 Ingress 资源的 informers。

创建自定义控制器： 调用 pkg.NewController 创建自定义控制器对象，并传入 Kubernetes 客户端和 informers。

启动 Shared Informer 工厂： 启动 informer 工厂，并等待缓存同步完成。

运行自定义控制器： 调用控制器的 Run 方法运行自定义控制器。

main.go 简略伪代码：
go
Copy code
// main 包声明
package main

// 导入必要的包

// main 函数，程序入口
func main() {
// 从用户的 kubeconfig 文件中构建 Kubernetes 配置

    // 根据配置创建 Kubernetes 客户端

    // 创建 shared informer 工厂，用于监视 Kubernetes 资源变更

    // 创建 Service 资源的 informer

    // 创建 Ingress 资源的 informer

    // 创建自定义控制器对象，传入 Kubernetes 客户端和 informers

    // 创建用于停止 informer 的通道

    // 启动 informer 工厂

    // 等待缓存同步完成

    // 运行自定义控制器

}
controller.go 逻辑描述：
导入包： 导入了程序所需的包，包括 Kubernetes 客户端 informers 包、网络 informers 包等。

控制器结构体定义： 定义了控制器结构体，包括 Kubernetes 客户端、Service 和 Ingress 资源列表器以及工作队列。

回调函数定义： 定义了处理 Service 和 Ingress 资源事件的回调函数，包括更新、添加和删除操作。

将对象添加到工作队列函数： 定义了将对象添加到工作队列的函数，其中使用了 cache.MetaNamespaceKeyFunc 生成键。

运行控制器主函数： 定义了运行控制器的主函数，该函数会不断从工作队列中取出键，并处理相应的资源对象。

创建控制器实例函数： 定义了创建控制器实例的函数，其中初始化了监听事件的 informer，并添加了相应的事件处理函数。

controller.go 简略伪代码：
go
Copy code
// 包声明
package pkg

// 导入必要的包

// 控制器结构体

// 更新 Service 资源的回调函数

// 添加 Service 资源的回调函数

// 删除 Ingress 资源的回调函数

// 将对象添加到工作队列的函数

// 运行控制器的主函数

// 创建控制器实例函数
这是对你提供的两个文件的简要逻辑描述和简略伪代码。如果你有特定的部分想要详细描述或者有其他问题，随时告诉我。
