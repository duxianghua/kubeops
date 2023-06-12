github token:
    username: kepten
    token: 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

Github repository: https://github.com/duxianghua/sockshop.git

## Install
```sh
helm upgrade --install keptn keptn --repo=https://charts.keptn.sh \
-n keptn --create-namespace \
--set=apiGatewayNginx.type=LoadBalancer \
--set=continuousDelivery.enabled=true \
--wait
```

## AUTH
``` sh
KEPTN_ENDPOINT=http://a3f35fb2272d6446eb151c51dc46f48b-941087129.ap-southeast-1.elb.amazonaws.com/api

keptn auth --endpoint=${KEPTN_ENDPOINT} --api-token=PqioywJEEFCRPYtkfEwJXoNwz4laFmLPDvk15xUSqhKUj
```

## Create Project
``` sh
keptn create project shipyard-sockshop --shipyard=./fast-project.yaml --git-user=keptn --git-token=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 --git-remote-url=https://github.com/duxianghua/sockshop.git
```

## Update Project
``` sh
keptn trigger delivery --project=shipyard-sockshop --service=test --image=test --tag=v0.0.1 --sequence=hotfix-delivery
```

## trigger delivery
``` sh
keptn trigger delivery --project=shipyard-sockshop --service=test --image=test:latest --sequence=dev
```

## Add Resources
```sh
keptn add-resource --project=shipyard-sockshop --service=test --all-stages --resource=helloservice.tgz --resourceUri=helm/helloservice.tgz

keptn add-resource --project=shipyard-sockshop --service=test --all-stages --resource=./endpoints.yaml --resourceUri=helm/endpoints.yaml
```