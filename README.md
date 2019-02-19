# smagnet
Change the number of Spanner Nodes

## Setup

### Create PubSub Topic for Cloud Scheduler

```
gcloud pubsub topics create $YOUR_TOPIC
```

### Function's Deploy

```
gcloud functions deploy HandleSpannerMagnet --runtime go111 --trigger-topic $YOUR_TOPIC
```

### Create Cloud Scheduler Job

Create a schedule for which you want to change the number of Spanner's Nodes.
If you increase at the beginning of the peak time and decrease after the end, you need two Jobs.

#### Job Message

```
{
  "projectId" : "hoge",
  "instanceId" : "fuga",
  "nodeNumber" : 7
}
```

```
gcloud beta scheduler jobs create pubsub $JOB --schedule=$SCHEDULE --topic=smagnet --message-body-from-file=example-pubsub-message.json
```