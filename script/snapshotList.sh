#!/bin/bash
#1 ip 2 hostname
ssh root@$1 'virsh snapshot-list '$2 