Сервис и клиент для выставления имени хоста в Linux, а также изменения списка DNS серверов. Использование gRPC, Cobra.


cd testovoe

protoc -I proto --go_out pb --go_opt paths=source_relative host.proto - команда для получения *.pb.go
    
protoc -I proto --go-grpc_out pb --go-grpc_opt paths=source_relative host.proto - - команда для получения *.pb.go
    

sudo /usr/local/go/bin/go run cmd/server/main.go - запуск сервера

go run cmd/client/main.go set-hostname name - запуск клиента изменение hostname

go run cmd/client/main.go get-dns-servers - запуск клиента вывод dns из файла /etc/resolv.conf

go run cmd/client/main.go add-dns-server dns - запуск клиента добавление dns в файл /etc/resolv.conf

go run cmd/client/main.go remove-dns-server dns - запуск клиента удаление dns из файла /etc/resolv.conf
