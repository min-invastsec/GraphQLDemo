echo /////go routine 5/////
echo -----GraphQL-----
go-wrk.exe -c 5 -d 10 -T 10 http://127.0.0.1:8080/graphql?query={hello}

echo -----Rest-----
go-wrk.exe -c 5 -d 10 -T 10 http://127.0.0.1:8080/rest?query={hello}


echo /////go routine 10/////
echo -----GraphQL-----
go-wrk.exe -c 10 -d 10 -T 10 http://127.0.0.1:8080/graphql?query={hello}

echo -----Rest-----
go-wrk.exe -c 10 -d 10 -T 10 http://127.0.0.1:8080/rest?query={hello}


echo /////go routine 30/////
echo -----GraphQL-----
go-wrk.exe -c 30 -d 10 -T 10 http://127.0.0.1:8080/graphql?query={hello}

echo -----Rest-----
go-wrk.exe -c 30 -d 10 -T 10 http://127.0.0.1:8080/rest?query={hello}