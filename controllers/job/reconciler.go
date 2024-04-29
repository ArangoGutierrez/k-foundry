/*
Copyright 2024 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package job

import (
	"context"
	"net/http"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/kcp"
	"sigs.k8s.io/controller-runtime/pkg/log"

	datav1alpha1 "github.com/ArangoGutierrez/k-foundry/apis/v1alpha1"
	slurmapi "github.com/ArangoGutierrez/k-foundry/openapi"
)

// Reconciler reconciles a Job object
type Reconciler struct {
	// Client interface to communicate with the API server. Reconciler needs this for
	// fetching objects.
	client.Client

	// Scheme is used by the kubebuilder library to set OwnerReferences. Every
	// controller needs this.
	Scheme *runtime.Scheme

	Recorder record.EventRecorder
}

// Reconcile TODO
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Include the clusterName from req.ObjectKey in the logger, similar to the namespace and name keys that are already
	// there.
	log = log.WithValues("clusterName", req.ClusterName)

	log.Info("Getting Job")
	var w datav1alpha1.Job
	if err := r.Client.Get(ctx, req.NamespacedName, &w); err != nil {
		if errors.IsNotFound(err) {
			// Normal - was deleted
			return ctrl.Result{}, nil
		}
		return ctrl.Result{Requeue: true}, err
	}

	if w.Status.ID != 0 {
		// Already processed
		return ctrl.Result{}, nil
	}

	log.Info("Creating Job", w.Name)
	jobID, err := createJob(w.Name, w.Spec.Script)
	if err != nil {
		return ctrl.Result{}, err
	}

	//log.Info("Listing all Jobs in the current logical cluster")
	//var list datav1alpha1.JobList
	//if err := r.Client.List(ctx, &list); err != nil {
	//	return ctrl.Result{}, err
	//}
	//
	log.Info("Patching Job status to store total Job count in the current logical cluster")
	orig := w.DeepCopy()
	w.Status.ID = int(jobID)
	if err := r.Client.Status().Patch(ctx, &w, client.MergeFromWithOptions(orig, client.MergeFromWithOptimisticLock{})); err != nil {
		return ctrl.Result{Requeue: true}, err
	}

	return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&datav1alpha1.Job{}).
		Complete(kcp.WithClusterInContext(r))
}

func createJob(jobName, script string) (int32, error) {
	cfg := slurmapi.NewConfiguration()
	cfg.HTTPClient = &http.Client{Timeout: time.Second * 3600}
	cfg.Scheme = "http"
	cfg.Host = "10.11.1.6:80"

	client := slurmapi.NewAPIClient(cfg)

	auth := context.WithValue(
		context.Background(),
		slurmapi.ContextAPIKeys,
		map[string]slurmapi.APIKey{
			"user":  {Key: os.Getenv("LOGNAME")},
			"token": {Key: os.Getenv("SLURM_JWT")},
		},
	)

	pingReq := client.SlurmAPI.SlurmV0040GetPing(auth)
	_, httpResp, err := pingReq.Execute()
	if err != nil {
		// print the http response
		if httpResp != nil {
			println(httpResp.Status)
		}
		return 0, err
	}

	nodes := "2"
	wd := "/tmp/"
	tasks := int32(8)
	name := jobName
	cpuPerTaks := int32(1)
	job := slurmapi.V0040JobSubmitReq{
		Job: &slurmapi.V0040JobDescMsg{
			Tasks:                   &tasks,
			Name:                    &name,
			Nodes:                   &nodes,
			CurrentWorkingDirectory: &wd,
			CpusPerTask:             &cpuPerTaks,
			Environment: []string{
				"PATH=/bin:/usr/bin/:/usr/local/bin/",
				"LD_LIBRARY_PATH=/lib/:/lib64/:/usr/local/lib",
			},
		},
		Script: &script,
	}

	j := client.SlurmAPI.SlurmV0040PostJobSubmit(auth).V0040JobSubmitReq(job)
	r, _, err := j.Execute()
	if err != nil {
		return 0, err
	}
	return *r.JobId, nil
}
