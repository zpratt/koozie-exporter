* Lead time
  * needs to come from the code repository
    * will need to consider multiple remote git repo managers, such as github, github enterprise, gitlab
    * cannot assume this could be provided by a pipeline
    * github could post to a webhook for this app
      * if public github, then assumes this app is publicly exposed
      * if github enterprise, then assumes github enterprise can reach the cluster
      * **note**: this requires more consideration from a network security standpoint
    * polling of github/gitlab could be used to provide simplified network rules
    * will require some kind of configuration or convention to correlate a kubernetes object to a github/gitlab repo
      * could be achieved with annotations
  * tracking commit to deploy: require docker images to be tagged with the SHA of the commit? otherwise, how will we correlate the 2?
* Change failure rate
  * if we track failure start time and resolution time, then that can be used to track MTTR
  * this could get very complex, depending on what datasource(s) is to be used to track failures
  * simple MVP approach: track pod restarts
  * more elaborate approaches could get signals from logging systems. maybe loki has something cool for this
* MTTR
  * dependent upon tracking the start and stop time of incidents
  * once change failure rate is figured out, it provides the data necessary to track MTTR
* Deployment frequency
  * can be tracked with the admission webhook
  * just need to record the timestamp of each deployment
  * does prometheus record the time when a counter is incremented? if so, a counter could be exposed for this