#! /bin/bash

helm template rancher ./rancher-2.5.9.tgz --output-dir . \
 --namespace cattle-system \
 --set hostname=rancher.scs.buaa.edu.cn \
 --set rancherImage=harbor.scs.buaa.edu.cn/rancher/rancher \
 --set systemDefaultRegistry=harbor.scs.buaa.edu.cn \
 --set useBundledSystemChart=true \
 --set tls=external