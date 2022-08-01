Generated with:
openssl req -new -newkey ec:<(openssl ecparam -name secp384r1) -sha256 -days 365 -nodes -x509 -subj '/CN=localhost' -keyout certificate.key -out certificate.crt
