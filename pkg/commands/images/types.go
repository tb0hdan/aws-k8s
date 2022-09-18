package images

import "time"

// "v1.25.0"
// kubectl get pods -o json
// https://kubernetes.io/docs/reference/kubectl/jsonpath/
type GetPodsResponse struct {
	ApiVersion string `json:"apiVersion"`
	Items      []struct {
		ApiVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations struct {
				KubernetesIoPsp     string `json:"kubernetes.io/psp"`
				ChecksumTokenSecret string `json:"checksum/token-secret,omitempty"`
				ChecksumConfigmap   string `json:"checksum/configmap,omitempty"`
				ChecksumHealth      string `json:"checksum/health,omitempty"`
				ChecksumScripts     string `json:"checksum/scripts,omitempty"`
				ChecksumSecret      string `json:"checksum/secret,omitempty"`
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			GenerateName      string    `json:"generateName"`
			Labels            struct {
				App                            string `json:"app,omitempty"`
				PodTemplateHash                string `json:"pod-template-hash,omitempty"`
				AppKubernetesIoInstance        string `json:"app.kubernetes.io/instance,omitempty"`
				AppKubernetesIoManagedBy       string `json:"app.kubernetes.io/managed-by,omitempty"`
				AppKubernetesIoName            string `json:"app.kubernetes.io/name,omitempty"`
				ControllerRevisionHash         string `json:"controller-revision-hash,omitempty"`
				HelmShChart                    string `json:"helm.sh/chart,omitempty"`
				StatefulsetKubernetesIoPodName string `json:"statefulset.kubernetes.io/pod-name,omitempty"`
				AppKubernetesIoComponent       string `json:"app.kubernetes.io/component,omitempty"`
			} `json:"labels"`
			Name            string `json:"name"`
			Namespace       string `json:"namespace"`
			OwnerReferences []struct {
				ApiVersion         string `json:"apiVersion"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
				Controller         bool   `json:"controller"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				Uid                string `json:"uid"`
			} `json:"ownerReferences"`
			ResourceVersion string `json:"resourceVersion"`
			Uid             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			Containers []struct {
				Env []struct {
					Name      string `json:"name"`
					ValueFrom struct {
						FieldRef struct {
							ApiVersion string `json:"apiVersion"`
							FieldPath  string `json:"fieldPath"`
						} `json:"fieldRef,omitempty"`
						SecretKeyRef struct {
							Key      string `json:"key"`
							Name     string `json:"name"`
							Optional bool   `json:"optional,omitempty"`
						} `json:"secretKeyRef,omitempty"`
					} `json:"valueFrom,omitempty"`
					Value string `json:"value,omitempty"`
				} `json:"env,omitempty"`
				EnvFrom []struct {
					SecretRef struct {
						Name string `json:"name"`
					} `json:"secretRef"`
				} `json:"envFrom,omitempty"`
				Image           string `json:"image"`
				ImagePullPolicy string `json:"imagePullPolicy"`
				LivenessProbe   struct {
					FailureThreshold    int `json:"failureThreshold"`
					InitialDelaySeconds int `json:"initialDelaySeconds"`
					PeriodSeconds       int `json:"periodSeconds"`
					SuccessThreshold    int `json:"successThreshold"`
					TcpSocket           struct {
						Port int `json:"port"`
					} `json:"tcpSocket,omitempty"`
					TimeoutSeconds int `json:"timeoutSeconds"`
					Exec           struct {
						Command []string `json:"command"`
					} `json:"exec,omitempty"`
				} `json:"livenessProbe,omitempty"`
				Name  string `json:"name"`
				Ports []struct {
					ContainerPort int    `json:"containerPort"`
					Name          string `json:"name"`
					Protocol      string `json:"protocol"`
				} `json:"ports,omitempty"`
				Resources struct {
					Limits struct {
						Cpu    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"limits,omitempty"`
					Requests struct {
						Cpu    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"requests,omitempty"`
				} `json:"resources"`
				StartupProbe struct {
					FailureThreshold    int `json:"failureThreshold"`
					InitialDelaySeconds int `json:"initialDelaySeconds"`
					PeriodSeconds       int `json:"periodSeconds"`
					SuccessThreshold    int `json:"successThreshold"`
					TcpSocket           struct {
						Port int `json:"port"`
					} `json:"tcpSocket"`
					TimeoutSeconds int `json:"timeoutSeconds"`
				} `json:"startupProbe,omitempty"`
				TerminationMessagePath   string `json:"terminationMessagePath"`
				TerminationMessagePolicy string `json:"terminationMessagePolicy"`
				VolumeMounts             []struct {
					MountPath string `json:"mountPath"`
					Name      string `json:"name"`
					ReadOnly  bool   `json:"readOnly,omitempty"`
				} `json:"volumeMounts"`
				Lifecycle struct {
					PreStop struct {
						Exec struct {
							Command []string `json:"command"`
						} `json:"exec"`
					} `json:"preStop"`
				} `json:"lifecycle,omitempty"`
				ReadinessProbe struct {
					Exec struct {
						Command []string `json:"command"`
					} `json:"exec"`
					FailureThreshold    int `json:"failureThreshold"`
					InitialDelaySeconds int `json:"initialDelaySeconds"`
					PeriodSeconds       int `json:"periodSeconds"`
					SuccessThreshold    int `json:"successThreshold"`
					TimeoutSeconds      int `json:"timeoutSeconds"`
				} `json:"readinessProbe,omitempty"`
				SecurityContext struct {
					RunAsNonRoot bool `json:"runAsNonRoot,omitempty"`
					RunAsUser    int  `json:"runAsUser"`
				} `json:"securityContext,omitempty"`
				Args    []string `json:"args,omitempty"`
				Command []string `json:"command,omitempty"`
			} `json:"containers"`
			DnsPolicy          string `json:"dnsPolicy"`
			EnableServiceLinks bool   `json:"enableServiceLinks"`
			NodeName           string `json:"nodeName"`
			PreemptionPolicy   string `json:"preemptionPolicy"`
			Priority           int    `json:"priority"`
			RestartPolicy      string `json:"restartPolicy"`
			SchedulerName      string `json:"schedulerName"`
			SecurityContext    struct {
				FsGroup int `json:"fsGroup,omitempty"`
			} `json:"securityContext"`
			ServiceAccount                string `json:"serviceAccount"`
			ServiceAccountName            string `json:"serviceAccountName"`
			TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
			Tolerations                   []struct {
				Effect            string `json:"effect"`
				Key               string `json:"key"`
				Operator          string `json:"operator"`
				TolerationSeconds int    `json:"tolerationSeconds"`
			} `json:"tolerations"`
			Volumes []struct {
				Name      string `json:"name"`
				Projected struct {
					DefaultMode int `json:"defaultMode"`
					Sources     []struct {
						ServiceAccountToken struct {
							ExpirationSeconds int    `json:"expirationSeconds"`
							Path              string `json:"path"`
						} `json:"serviceAccountToken,omitempty"`
						ConfigMap struct {
							Items []struct {
								Key  string `json:"key"`
								Path string `json:"path"`
							} `json:"items"`
							Name string `json:"name"`
						} `json:"configMap,omitempty"`
						DownwardAPI struct {
							Items []struct {
								FieldRef struct {
									ApiVersion string `json:"apiVersion"`
									FieldPath  string `json:"fieldPath"`
								} `json:"fieldRef"`
								Path string `json:"path"`
							} `json:"items"`
						} `json:"downwardAPI,omitempty"`
					} `json:"sources"`
				} `json:"projected,omitempty"`
				PersistentVolumeClaim struct {
					ClaimName string `json:"claimName"`
				} `json:"persistentVolumeClaim,omitempty"`
				Secret struct {
					DefaultMode int    `json:"defaultMode"`
					SecretName  string `json:"secretName"`
				} `json:"secret,omitempty"`
				ConfigMap struct {
					DefaultMode int    `json:"defaultMode"`
					Name        string `json:"name"`
				} `json:"configMap,omitempty"`
				EmptyDir struct {
				} `json:"emptyDir,omitempty"`
			} `json:"volumes"`
			Affinity struct {
				PodAntiAffinity struct {
					PreferredDuringSchedulingIgnoredDuringExecution []struct {
						PodAffinityTerm struct {
							LabelSelector struct {
								MatchLabels struct {
									AppKubernetesIoInstance  string `json:"app.kubernetes.io/instance"`
									AppKubernetesIoName      string `json:"app.kubernetes.io/name"`
									AppKubernetesIoComponent string `json:"app.kubernetes.io/component,omitempty"`
								} `json:"matchLabels"`
							} `json:"labelSelector"`
							Namespaces  []string `json:"namespaces"`
							TopologyKey string   `json:"topologyKey"`
						} `json:"podAffinityTerm"`
						Weight int `json:"weight"`
					} `json:"preferredDuringSchedulingIgnoredDuringExecution"`
				} `json:"podAntiAffinity"`
			} `json:"affinity,omitempty"`
			Hostname  string `json:"hostname,omitempty"`
			Subdomain string `json:"subdomain,omitempty"`
		} `json:"spec"`
		Status struct {
			Conditions []struct {
				LastProbeTime      interface{} `json:"lastProbeTime"`
				LastTransitionTime time.Time   `json:"lastTransitionTime"`
				Status             string      `json:"status"`
				Type               string      `json:"type"`
				Message            string      `json:"message,omitempty"`
				Reason             string      `json:"reason,omitempty"`
			} `json:"conditions"`
			ContainerStatuses []struct {
				ContainerID string `json:"containerID"`
				Image       string `json:"image"`
				ImageID     string `json:"imageID"`
				LastState   struct {
					Terminated struct {
						ContainerID string    `json:"containerID"`
						ExitCode    int       `json:"exitCode"`
						FinishedAt  time.Time `json:"finishedAt"`
						Reason      string    `json:"reason"`
						StartedAt   time.Time `json:"startedAt"`
					} `json:"terminated,omitempty"`
				} `json:"lastState"`
				Name         string `json:"name"`
				Ready        bool   `json:"ready"`
				RestartCount int    `json:"restartCount"`
				Started      bool   `json:"started"`
				State        struct {
					Running struct {
						StartedAt time.Time `json:"startedAt"`
					} `json:"running,omitempty"`
					Waiting struct {
						Message string `json:"message"`
						Reason  string `json:"reason"`
					} `json:"waiting,omitempty"`
				} `json:"state"`
			} `json:"containerStatuses"`
			HostIP string `json:"hostIP"`
			Phase  string `json:"phase"`
			PodIP  string `json:"podIP"`
			PodIPs []struct {
				Ip string `json:"ip"`
			} `json:"podIPs"`
			QosClass  string    `json:"qosClass"`
			StartTime time.Time `json:"startTime"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
}
