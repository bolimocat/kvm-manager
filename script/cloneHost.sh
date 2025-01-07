#!/bin/bash
#1 ip 2 source templates 3 new host
ssh root@$1 'virt-clone --auto-clone -o '$2' -n '$3