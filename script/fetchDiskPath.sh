#!/bin/bash
#1 ip 2 hostname
ssh root@$1 'virsh dumpxml '$2'|grep "source file"' > ./info/$2_disk_info