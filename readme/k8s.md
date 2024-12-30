# kubectl apply -f k8s/  => Rodando com kubernets
# minikube service nginx-proxy => acessando a aplicação
#  publicar minha imagem no docker hub 
# docker login
# docker tag docker-go-api:latest schzimmyd/docker-go-api:latest
# docker push schzimmyd/odcker-go-api:latest
# publicar no arquivo yaml também
# kubectl apply -f k8s/
# kubectl get pods
# kubectl get replicasets
# kubectl get svc => vejo as portas como se fosse um container
# minikube ip
# curl http://192.168.49.2:30191
# kubectl describe pod nginx-proxy-797646587d-fdv44
# minikube service nginx-proxy => cria um túnel para acessar o serviço localmente
# minikube service go-api
# NodePort e ClusterIP