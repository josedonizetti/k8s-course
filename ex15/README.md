```
kubectl set image deployment ex14-deployment my-app=myapp:v3
kubectl rollout status deployment ex14-deployment
kubectl rollout history deployment ex14-deployment
kubectl rollout undo deployment ex14-deployment
```
