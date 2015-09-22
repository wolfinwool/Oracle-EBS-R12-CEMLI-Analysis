#Oracle EBS R12 CEMLI Analysis
This repository holds source code for a CEMLI analysis tool. The development is restricted to support for upgrades to Oracle E-business Suite, Release 12 from an existing Oracle database.

##Goals
* Create a lightweight analysis tool to assist in CEMLI analysis.
* Accurately detect changes between legacy releases and R12.
  * Detect sources of breakage due to the upgrades.
* TBD

##Usage
Currently the CEMLI analysis tool is in alpha development and is currently not designed for dynamic use. It is also separated into two separate tools (parser/analyzer). However, if you are GO savvy you may hardcode the source with the proper directories to get it working.

##Todo

* Lots of code refactoring since this tool is going public.
* Application restructuring.
* Combination of both tools into one package.
* Proper Go formatting.
