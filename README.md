Сервис и клиент для выставления имени хоста в Linux, а также изменения списка DNS серверов. Использование gRPC, Cobra.

cd testovoe
sudo /usr/local/go/bin/go run cmd/server/main.go - запуск сервера
go run cmd/client/main.go set-hostname <name> - запуск клиента изменение hostname
go run cmd/client/main.go get-dns-servers - запуск клиента вывод dns из файла /etc/resolv.conf
go run cmd/client/main.go add-dns-server <dns> - запуск клиента добавление dns в файл /etc/resolv.conf
go run cmd/client/main.go remove-dns-server <dns> - запуск клиента удаление dns из файла /etc/resolv.conf
