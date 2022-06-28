/*
Copyright 2022 The Koordinator Authors.

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

package groupidentity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"

	ext "github.com/koordinator-sh/koordinator/apis/extension"
	runtimeapi "github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	"github.com/koordinator-sh/koordinator/pkg/util"
	"github.com/koordinator-sh/koordinator/pkg/util/system"
)

func Test_bvtPlugin_SetPodBvtValue_Proxy(t *testing.T) {
	defaultRule := &bvtRule{
		podQOSParams: map[ext.QoSClass]int64{
			ext.QoSLSR: 2,
			ext.QoSLS:  2,
			ext.QoSBE:  -1,
		},
		kubeQOSDirParams: map[corev1.PodQOSClass]int64{
			corev1.PodQOSGuaranteed: 0,
			corev1.PodQOSBurstable:  2,
			corev1.PodQOSBestEffort: -1,
		},
		kubeQOSPodParams: map[corev1.PodQOSClass]int64{
			corev1.PodQOSGuaranteed: 2,
			corev1.PodQOSBurstable:  2,
			corev1.PodQOSBestEffort: -1,
		},
	}
	type fields struct {
		systemSupported *bool
	}
	type args struct {
		request  *runtimeapi.PodSandboxHookRequest
		response *runtimeapi.PodSandboxHookResponse
	}
	type want struct {
		bvtValue *int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "set ls pod bvt",
			fields: fields{
				systemSupported: pointer.Bool(true),
			},
			args: args{
				request: &runtimeapi.PodSandboxHookRequest{
					Labels: map[string]string{
						ext.LabelPodQoS: string(ext.QoSLS),
					},
					CgroupParent: "kubepods/pod-guaranteed-test-uid/",
				},
				response: &runtimeapi.PodSandboxHookResponse{},
			},
			want: want{
				bvtValue: pointer.Int64(2),
			},
		},
		{
			name: "set be pod bvt",
			fields: fields{
				systemSupported: pointer.Bool(true),
			},
			args: args{
				request: &runtimeapi.PodSandboxHookRequest{
					Labels: map[string]string{
						ext.LabelPodQoS: string(ext.QoSBE),
					},
					CgroupParent: "kubepods/besteffort/pod-besteffort-test-uid/",
				},
				response: &runtimeapi.PodSandboxHookResponse{},
			},
			want: want{
				bvtValue: pointer.Int64(-1),
			},
		},
		{
			name: "set be pod bvt but system not support",
			fields: fields{
				systemSupported: pointer.Bool(false),
			},
			args: args{
				request: &runtimeapi.PodSandboxHookRequest{
					Labels: map[string]string{
						ext.LabelPodQoS: string(ext.QoSBE),
					},
					CgroupParent: "kubepods/besteffort/pod-besteffort-test-uid/",
				},
				response: &runtimeapi.PodSandboxHookResponse{},
			},
			want: want{
				bvtValue: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testHelper := system.NewFileTestUtil(t)
			initCPUBvt(tt.args.request.CgroupParent, 0, testHelper)

			b := &bvtPlugin{
				rule:         defaultRule,
				sysSupported: tt.fields.systemSupported,
			}
			ctx := &protocol.PodContext{}
			ctx.FromProxy(tt.args.request)
			err := b.SetPodBvtValue(ctx)
			ctx.ProxyDone(tt.args.response)
			assert.NoError(t, err)

			if tt.want.bvtValue == nil {
				assert.Nil(t, ctx.Response.Resources.CPUBvt, "bvt value should be nil")
			} else {
				assert.Equal(t, *tt.want.bvtValue, *ctx.Response.Resources.CPUBvt, "pod bvt in response should be equal")
				gotBvt := getPodCPUBvt(tt.args.request.CgroupParent, testHelper)
				assert.Equal(t, *tt.want.bvtValue, gotBvt, "pod bvt should equal")
			}
		})
	}
}

func Test_bvtPlugin_SetKubeQOSBvtValue_Reconciler(t *testing.T) {
	defaultRule := &bvtRule{
		podQOSParams: map[ext.QoSClass]int64{
			ext.QoSLSR: 2,
			ext.QoSLS:  2,
			ext.QoSBE:  -1,
		},
		kubeQOSDirParams: map[corev1.PodQOSClass]int64{
			corev1.PodQOSGuaranteed: 0,
			corev1.PodQOSBurstable:  2,
			corev1.PodQOSBestEffort: -1,
		},
		kubeQOSPodParams: map[corev1.PodQOSClass]int64{
			corev1.PodQOSGuaranteed: 2,
			corev1.PodQOSBurstable:  2,
			corev1.PodQOSBestEffort: -1,
		},
	}
	type fields struct {
		sysSupported *bool
	}
	type args struct {
		kubeQOS corev1.PodQOSClass
	}
	type want struct {
		bvtValue *int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "set guaranteed dir bvt",
			fields: fields{
				sysSupported: pointer.BoolPtr(true),
			},
			args: args{
				kubeQOS: corev1.PodQOSGuaranteed,
			},
			want: want{
				bvtValue: pointer.Int64(0),
			},
		},
		{
			name: "set burstable dir bvt",
			fields: fields{
				sysSupported: pointer.BoolPtr(true),
			},
			args: args{
				kubeQOS: corev1.PodQOSBurstable,
			},
			want: want{
				bvtValue: pointer.Int64(2),
			},
		},
		{
			name: "set be dir bvt",
			fields: fields{
				sysSupported: pointer.BoolPtr(true),
			},
			args: args{
				kubeQOS: corev1.PodQOSBestEffort,
			},
			want: want{
				bvtValue: pointer.Int64(-1),
			},
		},
		{
			name: "set be dir bvt but system not support",
			fields: fields{
				sysSupported: pointer.BoolPtr(false),
			},
			args: args{
				kubeQOS: corev1.PodQOSBestEffort,
			},
			want: want{
				bvtValue: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testHelper := system.NewFileTestUtil(t)
			kubeQOSDir := util.GetKubeQosRelativePath(tt.args.kubeQOS)
			initCPUBvt(kubeQOSDir, 0, testHelper)

			b := &bvtPlugin{
				rule:         defaultRule,
				sysSupported: tt.fields.sysSupported,
			}
			ctx := &protocol.KubeQOSContext{}
			ctx.FromReconciler(tt.args.kubeQOS)
			err := b.SetKubeQOSBvtValue(ctx)
			ctx.ReconcilerDone()

			assert.NoError(t, err)

			if tt.want.bvtValue == nil {
				assert.Nil(t, ctx.Response.Resources.CPUBvt, "bvt value should be nil")
			} else {
				assert.Equal(t, *tt.want.bvtValue, *ctx.Response.Resources.CPUBvt, "kube qos bvt in response should be equal")
				gotBvt := getPodCPUBvt(ctx.Request.CgroupParent, testHelper)
				assert.Equal(t, *tt.want.bvtValue, gotBvt, "kube qos bvt should be equal")
			}
		})
	}
}