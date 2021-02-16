# kubernetes-config

Buildar o dockerfile
 - docker build -t diorgenesbk/hello-go:1 .

Push na imagem docker
 - docker push diorgenesbk/hello-go:1

Gerar o pod
 - kubectl apply -f .\k8s\pod.yaml

Gerar a ReplicaSet
 - kubectl apply -f .\k8s\replicaset.yaml

Gerar o deployment
 - kubectl apply -f .\k8s\deployment.yaml

Voltar o deployment para uma determinada versão
 - kubectl rollout undo deployment goserver --to-revision=1

Ver status dos pods
 - kubectl get pods

Ver o detalhamento de um determinado pod
 - kubectl describe pod goserver-hqdr7
