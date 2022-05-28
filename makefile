server:
	go run .

clientcred: 
	curl -Lv --cacert "certificate.pem"  https://localhost:8000/

clientnocred: 
	curl -Lv https://localhost:8000/