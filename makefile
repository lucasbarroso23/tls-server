server:
	go run .

client: 
	curl --cacert "./credentials/certificate.pem"  https://localhost:8000/