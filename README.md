#Oracle EBS R12 CEMLI Analysis
This repository holds source code for a CEMLI analysis tool. The development is restricted to support for upgrades to Oracle E-business Suite, Release 12 from an existing Oracle EBS release.

Uses AD diff html files that can be downloaded from Oracle Support: EBS Data Model Comparison Report Overview (Doc ID 1290886.1), or use one of the provided AD diff files in the /resources directory.

To be used for upgrades to the following EBS releases:

* Release 12.1.3
* Release 12.2.2
* Release 12.2.3
* Release 12.2.4


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
