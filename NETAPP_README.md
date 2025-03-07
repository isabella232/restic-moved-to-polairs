#Netapp Readme
# Overview
This fork has changes that make restic work with NetApp ONTAP and Storage Grid.

This repo uses main for its default branch and will update with master from restic/restic as we need.

The intent is to make these changes open source and let people enjoy them.  

NetApp does not imply or guarentee any support, but we are happy to help if we can.  File an issue or email the contributors if you need help

# NON ROOT RUN Changes
As of the 2.0 bump, this container is designed to run a non-root.
It needs a specific command line option for docker:

`docker run --cap-add DAC_READ_SEARCH -it --rm netapp-restic`

The --cap-add DAC_READ_SEARCH is needed by restic to do backups when running non root.

#License Notice
The license for Restic is 2 Clause BSD.  We will maintain that license but deep open source scans done by NetApp show some golang modules
include GPLv2 and LGPL v3 code.  Please keep that in mind when you use this software.  Again, this is in base Restic, nothing we have added.

