**Checkout and build for minishift

TBD


**Create project and service accounts
oc new-project insights-scan
oc create serviceaccount insights-scan
oc adm policy add-cluster-role-to-user cluster-admin system:serviceaccount:insights-scan:insights-scan
oc adm policy add-scc-to-user privileged system:serviceaccount:insights-scan:insights-scan

** Insights credentials 

cp  deploy/secrets.yaml.template to deploy/secrets.yaml and replace username and password
oc create -f secrets.yaml



** Run Job
oc project insights-scan
oc create -f deploy/scan-job.yaml

** Monitor
oc get jobs
oc get pods
oc logs -f <pod-name>
