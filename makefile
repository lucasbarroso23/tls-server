server:
	go run .

clientcred: 
	curl --cacert "certificate.pem"  https://localhost:8000/

clientnocred: 
	curl https://localhost:8000/