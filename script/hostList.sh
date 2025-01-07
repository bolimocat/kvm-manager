#!/bin/bash
#1 ip
if test $2 = 'running'
	then ssh root@$1 'virsh list '
fi
if test $2 = 'all'
	then ssh root@$1 'virsh list --all'
fi