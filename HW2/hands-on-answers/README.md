# Questions

1.   
	1. Instead of listening on `0.0.0.0`, you should listen on an internal IP, such as `127.0.0.1`, or to your machine's private LAN IP address (something along the lines of `192.168.1.x` or `10.x.x.x`).    

	2. We want to use a private network with no route to the internet. No public IP or NAT or port forward to the service. In Docker you can use a network with `internal: true`.    

	3. You can use a firewall that drops external requests and only allows traffic from your local subnet range.

	4. Make the proxy the only gateway.    

2. The legacy service speaks in plain HTTP, and can't upgrade to support HTTPS. This will cause security problems and data must be encrypted. A simple way to do this is to make the proxy support TLS. This way we do not need to change the legacy server and it can continue only handling the business logic. JWT without TLS is a weakness because it can be sniffed and replayed.   
3. Code is in the `code` folder. First we build `docker compose up --build -d` and then we run the following six tests:   

    1. 

        Rooha@Delaraams-MacBook-Air code % curl -i http://localhost/get

        HTTP/1.1 401 Unauthorized
        Content-Type: text/plain; charset=utf-8
        Www-Authenticate: Bearer realm="proxy"
        X-Content-Type-Options: nosniff
        Date: Mon, 02 Nov 2026 19:11:55 GMT
        Content-Length: 13

        unauthorized

    Since there was no token, it was rejected by our proxy.    

    2. 

        Rooha@Delaraams-MacBook-Air code % curl -i -X POST http://localhost/login -d '{"username":"admin","password":"wrong"}'

        HTTP/1.1 401 Unauthorized
        Content-Type: text/plain; charset=utf-8
        X-Content-Type-Options: nosniff
        Date: Mon, 02 Nov 2026 19:12:31 GMT
        Content-Length: 20

        invalid credentials

    Quite straightforwards, wrong password -> invalid credentials.

    3. 

        Rooha@Delaraams-MacBook-Air code % TOKEN=$(curl -s -X POST http://localhost/login \
        -d '{"username":"admin","password":"secret"}' | sed 's/.*"token":"\([^"]*\)".*/\1/')
        echo $TOKEN
        eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTM2NTAzNjMsImlhdCI6MTc5MzY0Njc2Mywic3ViIjoiYWRtaW4ifQ.QG4iI642F0WITXVgiiRWDsa7uJsNSmMoeu00uhZhkrA

    We got a token!

    4.

        Rooha@Delaraams-MacBook-Air code % curl -i -H "Authorization: Bearer $TOKEN" http://localhost/get
        HTTP/1.1 200 OK
        Access-Control-Allow-Credentials: true
        Access-Control-Allow-Origin: *
        Content-Length: 527
        Content-Type: application/json; charset=utf-8
        Date: Mon, 02 Nov 2026 19:12:53 GMT

        {
        "args": {},
        "headers": {
            "Accept": [
            "*/*"
            ],
            "Accept-Encoding": [
            "gzip"
            ],
            "Authorization": [
            "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTM2NTAzNjMsImlhdCI6MTc5MzY0Njc2Mywic3ViIjoiYWRtaW4ifQ.QG4iI642F0WITXVgiiRWDsa7uJsNSmMoeu00uhZhkrA"
            ],
            "Host": [
            "localhost"
            ],
            "User-Agent": [
            "curl/8.7.1"
            ],
            "X-Forwarded-For": [
            "192.168.65.1"
            ]
        },
        "method": "GET",
        "origin": "192.168.65.1",
        "url": "http://localhost/get"
        }

    We did a request with our token and it worked.

    5. 

        Rooha@Delaraams-MacBook-Air code % curl -i -H "Authorization: Bearer ${TOKEN}x" http://localhost/get

        HTTP/1.1 401 Unauthorized
        Content-Type: text/plain; charset=utf-8
        Www-Authenticate: Bearer realm="proxy"
        X-Content-Type-Options: nosniff
        Date: Mon, 02 Nov 2026 19:13:05 GMT
        Content-Length: 13

        unauthorized

    Used a bad token (added x at the end), and didn't get results.

    6.

        Rooha@Delaraams-MacBook-Air code % curl -m 3 http://localhost:8081/get
        curl: (7) Failed to connect to localhost port 8081 after 1 ms: Couldn't connect to server

    You can't reach the backend directly and can only reach it on port 80 through our proxy, as intended.


# AI Usage

AI was used in answering all three questions. The first two questions were answered after searching in the internet, asking Claude and reading Google's AI overview, understanding it and writing down the answer in my own words. In the third question, AI was used in generating the code and tests. The code and test outputs were reviewed and studied. Test output explanations are written by me. 
    
