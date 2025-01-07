#!/bin/bash
#1 ip 2 source templates
ssh root@$1 'virsh snapshot-create '$2