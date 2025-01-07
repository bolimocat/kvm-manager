# kvm-manager
A command-line tool for controlling KVM virtual machines

kvm manager instructions : 
-c help : Display help content 
 
-c source-list : Display of all available physical resource status from properties file
 
Generate clone hosts using existing images
./kmanager -framework x86 -host 10.30.0.13 -method clone -path /ssd -source 0 -machine test_uos
The meaning is : Select a machine with an IP of 0.13 from the x86 architecture, use cloning method,
and choose the first (0) image for replication. The name of the new host is test_uos. 
Detail :  
-framework : choose the build framework , the arguments is x86 or arm 
when user choose a framework by input that kmanager will judge the host is in list or not
-host : the host which you want to generate the new machine
the host is list in source-list
-method : Choose an installation method , the arguments is NEW or clone 
NEW means create a blank disk image then install a new system by iso (Currently not completed)
clone means copy from a exist image which you selected 
-path : the path which you want the new machine (Do not pay attention when using clones)
the path is list in source-list 
-source : the iso file or the templates file which you can select (Do not pay attention when using clones)
the file is different from x86 framework and arm framework 

Here's how to control a virtual host :  
1 List all running virtual machines on the specified physical machine :
./kmanager -c host-list -host hostIP 
2 Start a virtual machine on a specified physical machine
./kmanager -c host-start -host hostIP -machine hostName
3 Stop a virtual machine on a specified physical machine
./kmanager -c host-stop -host hostIP -machine hostName
4 Clear virtual machines from the list on the specified physical machine
./kmanager -c host-remove -host hostIP -machine hostName
5 Obtain physical file information generated by a certain host
./kmanager -c disk-path -host hostIP -machine hostName
Generate a disk information file starting with the current host name
 in the info directory. After clearing the virtual machine from the list, the actual physical files can be deleted based on this information.
6 Generate Snapshot
./kmanager -c snapshot-create -host hostIP -machine hostName
7 Display the existing snapshot list of a virtual host
./kmanager -c snapshot-create -host hostIP -machine hostName
8 Restore the state of the virtual host to a snapshot
./kmanager -c snapshot-revert -host hostIP -machine hostName -snapshot snapshotName
9 Delete snapshot
./kmanager -c snapshot-drop -host hostIP -machine hostName -snapshot snapshotName
Only by deleting all snapshots can this virtual host be deleted.
