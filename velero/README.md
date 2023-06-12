## Backup
``` sh
BUCKET=dev-eks-backup-velero
REGION=ap-southeast-1
velero install \
    --provider aws \
    --plugins velero/velero-plugin-for-aws:v1.6.0 \
    --bucket $BUCKET \
    --secret-file ./credentials \
    --backup-location-config region=$REGION \
    --snapshot-location-config region=$REGION 
```

# migration-case
``` sh
velero install --provider aws \
    --image velero/velero:v1.8.0 \
    --plugins velero/velero-plugin-for-aws:v1.4.0 \
    --bucket velero-migration-demo \
    --secret-file xxxx/aws-credentials-cluster1 \
    --backup-location-config region=us-east-2 \
    --snapshot-location-config region=us-east-2
```
