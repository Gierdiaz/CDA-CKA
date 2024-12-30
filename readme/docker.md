# docker container list = docker container ls
# docker container list -a 
# docker container start go-api ou adicionar mais de um container se for preciso. docker container start go-api nginx-proxy
# docker container stop go-api. Você pode passar mais de um e também pode usar o name do container ou o ID dele
# docker container pause go-api
# docker container unpuase go-api
# docker container rm go-api
# docker container restart go-api
# docker container logs -f go-api => seguindo o log em tempo real
# docker container logs --tail 10 go-api => opção caso eu queira ver as últimas 10 linhas do log
# docker container exec -it go-api bash => caso eu queira executar algo dentro do container do go. Usando o bash eu ainda vejo o ID do container