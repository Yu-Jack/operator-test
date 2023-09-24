package main

import (
	"context"
	"fmt"
	"os"

	api "github.com/Yu-Jack/operator-test/apis/cronjob/v1alpha1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

type reconciler struct {
	client.Client
}

func (r *reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("cronjob", req.NamespacedName)
	log.V(1).Info("reconciling cronjob")

	var cronjob api.CronJob
	if err := r.Get(ctx, req.NamespacedName, &cronjob); err != nil {
		log.Error(err, "unable to get cronjob")
		return ctrl.Result{}, err
	}

	fmt.Printf("Sync/Add/Update for foo %s\n", cronjob.GetName())
	fmt.Println("cronjob.Spec.Foo: ", cronjob.Spec.Foo)
	return ctrl.Result{}, nil
}

func main() {
	ctrl.SetLogger(zap.New())
	// 创建 Manager，创建时设置 Leader Election 相关的参数
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		LeaderElection:          true,
		LeaderElectionID:        "sample-controller",
		LeaderElectionNamespace: "kube-system",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// in a real controller, we'd create a new scheme for this
	err = api.AddToScheme(mgr.GetScheme())
	if err != nil {
		setupLog.Error(err, "unable to add scheme")
		os.Exit(1)
	}
	// 创建对 CronJob 进行调谐的 controller
	err = ctrl.NewControllerManagedBy(mgr).
		For(&api.CronJob{}).
		Complete(&reconciler{
			Client: mgr.GetClient(),
		})
	if err != nil {
		setupLog.Error(err, "unable to create controller")
		os.Exit(1)
	}

	// 启动 Manager，Manager 将启动其管理的所有 controller 以及 webhook server
	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
