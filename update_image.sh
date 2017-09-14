#!/bin/sh -eux

kubectl set image deployment/go-sample go-sample=gcr.io/$PROJECT_ID/sample-gke:$REVISION_ID
kubectl set image deployment/go-sample go-sample-nginx=gcr.io/$PROJECT_ID/sample-gke-nginx:$REVISION_ID
