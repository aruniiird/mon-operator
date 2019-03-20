apiVersion: "operator.monstor.org/v1alpha1"
kind: #TBD
metadata:
  # Name for this node
  name: #TBD
  # CRD is namespaced
  namespace: #TBD
  annotations:
    # Applied by operator when it creates/manages this object from a template.
    # When this is present, contents will be dynamically adjusted accd to the
    # template in the cluster CR.
    # When this annotation is present, the admin may only modify
    # .spec.desiredState or delete the CR. Any other change will be
    # overwritten.
    anthill.gluster.org/template: template-name
  version: <string>
spec:
  # Nodes belong to a cluster
  cluster: my-cluster
  # Only 1 of external | storage
  storages:
    - name: [ceph | gluster | cassandra]
      version: xxx
      alert:
        - enabled: [yes | no]
          labelSelctor: <key_val_pair_map>
          prometheusNamespace: <optional_str_if_alert_is_yes>
      dashboard: 
        - enabled: [yes | no]
          labelSelctor: <key_val_pair_map>
          grafanaNamespace: <optional_str_if_dashboard_is_yes>
    ...
  nodeAffinity:
    # https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
    # For admin created GlusterNodes, this needs to specify a node selector
    # that matches exactly one node. For template-based GNs, this will inherit
    # from the template.
    ...
status:
  # TBD operator state
  # Possible states: (enabled | deleting | disabled)
  # see other 'currentState' keys in other operators
  currentState: enabled
