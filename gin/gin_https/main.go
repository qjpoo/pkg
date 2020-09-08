package main

import (
	// "crypto/rand"
	// "crypto/rsa"
	// "crypto/x509"
	// "crypto/x509/pkix"
	// "encoding/pem"
	// "math/big"
	// "net"
	// "os"
	// "time"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	 max:=new(big.Int).Lsh(big.NewInt(1),128)
	 serialNumber,_:=rand.Int(rand.Reader,max)
	 subject:=pkix.Name{
	  Organization:[]string{"燕子李三"},
	  OrganizationalUnit:[]string{"Books"},
	  CommonName:"GO Web",
	 }

	 rootTemplate := x509.Certificate{
	  SerialNumber: serialNumber,
	  Subject:subject,
	  NotBefore: time.Now(),
	  NotAfter:time.Now().Add(365*time.Hour),
	  KeyUsage:x509.KeyUsageKeyEncipherment,
	  ExtKeyUsage:[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	  IPAddresses:[]net.IP{net.ParseIP("127.0.0.1")},
	 }
	 pk,_:=rsa.GenerateKey(rand.Reader,2048)
	 derBytes,_:=x509.CreateCertificate(rand.Reader,&rootTemplate,&rootTemplate,&pk.PublicKey,pk)

	 certOut,_:=os.Create("cert.pem")
	 pem.Encode(certOut,&pem.Block{Type:"CERTIFICATE",Bytes:derBytes})

	 certOut.Close()

	 keyOut,_:=os.Create("key.pem")
	 pem.Encode(keyOut,&pem.Block{Type:"RSA PRIVATE KEY",Bytes:x509.MarshalPKCS1PrivateKey(pk)})
	 keyOut.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, TLS!\n")
	})

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("About to listen on 8000. Go to https://127.0.0.1:8000/")
	err := http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", nil)
	log.Fatal(err)

}
