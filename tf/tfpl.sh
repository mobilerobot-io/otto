#!/bin/sh

terraform plan \
		-var "do_token=${DO_PAT}" \
		-var "pub_key=$HOME/.ssh/id_rsa.pub" \
		-var "pvt_key=$HOME/.ssh/id_rsa" \
		-var "ssh_fingerprint=6f:07:da:d2:f7:02:09:b5:ea:d5:75:b8:70:44:c6:01"