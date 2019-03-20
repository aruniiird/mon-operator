Installer
-----------
Create a 'prometheus-rule object' from ceph-mixins
Requirements - Promethes and Ceph
Pass to the prometheus [if/when it is available] (use watch loop for prometheus CRD)
if Ceph is not available we should not deploy the alerts
  - Log level message to indicate Ceph's non availability

Uninstaller
-----------
Remove the given prometheus-rule object from the prometheus instance.
Catch here is to track the exact object which we installed.

Upgrade/Downgrade
-----------------
Detect the ceph version and replace with the docker image version
