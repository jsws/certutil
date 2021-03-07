# Certutil
Certutil is a cli to help you to easily look inside certificates and debug issues.  
![image](https://user-images.githubusercontent.com/5903484/110249098-4c67d200-7f6c-11eb-8fc1-870da891586f.png)

## Motivation
I got tired of googling for openssl commands and wanted something simple and easy to use with a memorable cli. 

## Usage
Install with `go get`

```
$ go get github.com/jsws/certutil
```

### Info
`info` will print out certificate information in an easily readable format. It will also perform AIA fetching.
```
$ certutil info connect github.com
Connecting to github.com:443
Connected to 140.82.121.4:443
Recieved 2 certificate(s).

Certificate 0
  Subject: CN=github.com,O=GitHub\, Inc.,L=San Francisco,ST=California,C=US
  Issuer:  CN=DigiCert SHA2 High Assurance Server CA,OU=www.digicert.com,O=DigiCert Inc,C=US
  DNS Names: github.com, www.github.com
  Vailidity Period
      Not Before:  2020-05-05 00:00:00 +0000 UTC ✅
      Not After:   2022-05-10 12:00:00 +0000 UTC ✅

Certificate 1
  Subject: CN=DigiCert SHA2 High Assurance Server CA,OU=www.digicert.com,O=DigiCert Inc,C=US
  Issuer:  CN=DigiCert High Assurance EV Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
  Vailidity Period
      Not Before:  2013-10-22 12:00:00 +0000 UTC ✅
      Not After:   2028-10-22 12:00:00 +0000 UTC ✅

Certificate is valid 🔒
```
### Save
`save` will save the certificates presented by a server to a file in PEM format. The file to the save the PEM chain in is given with the `--output` or `-o` flag. If no output file is set the PEM encoded certificate is outputted to stdout.
```
$ certutil save connect github.com -o file.pem
Connecting to jsws.co.uk:443
Connected to 185.199.111.153:443
Recieved 2 certificate(s).
Saving to file.pem
```

```
$ certutil save connect github.com
Connecting to github.com:443
Connected to 140.82.121.3:443
Recieved 2 certificate(s).
-----BEGIN CERTIFICATE-----
...
```


#### Connect
The `connect` subcommand can be used with the `info` and `save` command to connect to a server presenting TLS certificates. The `--servername` or `-s` flag can be used to set the SNI.

#### Read
The `read` subcommand can be used with the `info` command to display certificate information from a local PEM encoded certificate file. The path of the file is given as an argument.
```
$ certutil info read certs/file.pem
```

## Limitations
`info.IsValid()` shouldn't be used for anything important as it doesn't do revocation checking or certificate transparency checking.

## TODO:
- [ ] Ability to show full cert chain when AIA fetching is used.
- [ ] Add verbose output for AIA
- [ ] Add revocation checking
