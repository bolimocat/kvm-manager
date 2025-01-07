#!/bin/bash
#1 ip 2 hostname 3 snapshot name
ssh root@$1 'virsh snapshot-delete '$2' --snapshotname '$3