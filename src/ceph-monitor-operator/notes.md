Installer
-----------
Create a 'prometheus-rule object' from ceph-mixins
Requirements - Prometheus and Ceph
Pass to the prometheus [if/when it is available] (use watch loop for prometheus CRD)
if Ceph is not available we should not deploy the alerts
  - Log level message to indicate Ceph's non availability
  - Mixins not deployed, should be indicated through watch events
Minimum Ceph-Mixin version check for the Ceph versions

Namespace selection ?
Monitoring operator can be deployed in any namespace,
Rule file should be deployed in the same name space where prometheus is running

Docker Image build has tobe done through operator CI
Nightly builds can tag Ceph-Mixin master
Release builds should have particular Ceph-Mixin version

Misc details
------------
RBAC, SA, SCC and CRD
Node affinity
CRD design doc should be added first

Uninstaller
-----------
Remove the given prometheus-rule object from the prometheus instance.
Catch here is to track the exact object which we installed.

Upgrade/Downgrade
-----------------
Detect the ceph version and replace with the docker image version

______

Timelines
---------
March-2019 end operator work should be finished
April first week - Deployment / Integration 

