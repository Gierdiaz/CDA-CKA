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
# minikube mount /home/allison/workspace/Docker/html:/home/allison/workspace/Docker/html => preciso rodar isso para eu acessar o nginx
# kubeclt logs nginx-proxy
# minikube tunnel => o minikube não consegue provisionar um LoadBalancer então ele fica como pending. Usando o minkube tunnel ele cria um endereço IP externo permitindo acesso ao LoadBalance. Quando o serviço é encerrado ele volta ao normal. 
# kubectl port-forwasrd svc/go-api 8081:3000 => acessando uma porta externa pelo cluster
# kubectl port-forward svc/nginx-proxy 8080:80
# kubectl exec -it nginx-proxy-5865f7bf96-xdpdf -- /bin/bash