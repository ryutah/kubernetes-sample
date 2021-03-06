#!/bin/sh -eux

gcloud container clusters get-credentials sample-gke --zone us-east1-d
kubectl set image deployment/go-sample go-sample=gcr.io/$PROJECT_ID/sample-gke:$REVISION_ID
kubectl set image deployment/go-sample go-sample-nginx=gcr.io/$PROJECT_ID/sample-gke-nginx:$REVISION_ID
