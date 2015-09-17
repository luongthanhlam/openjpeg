
# OPENJPEG Library and Applications

## What is OpenJPEG ?

OpenJPEG is an open-source JPEG 2000 codec written in C language. It has been developed in order to promote the use of [JPEG 2000](http://www.jpeg.org/jpeg2000), a still-image compression standard from the Joint Photographic Experts Group ([JPEG](http://www.jpeg.org)).  Since April 2015, it is officially recognized by ISO/IEC and ITU-T as a [JPEG 2000 Reference Software](http://www.itu.int/rec/T-REC-T.804-201504-I!Amd2).

## Who can use the code ?

Anyone. As the OpenJPEG code is released under the [2-clauses BSD license](https://github.com/uclouvain/openjpeg/blob/master/LICENSE), anyone can use or modify the code, even for commercial applications. The only restriction is to retain the copyright in the sources or in the binaries documentation. Of course, if you modified the code in a way that might be of interest for other users, you are encouraged to share it (through a [github pull request](https://github.com/uclouvain/openjpeg/pulls) or by filling an [issue](https://github.com/uclouvain/openjpeg/issues)) but this is not a requirement.

## How to install and use OpenJPEG ?
API Documentation needs a major refactoring. Meanwhile, you can check [installation](https://github.com/uclouvain/openjpeg/wiki/Installation) instructions and [codec documentation](https://github.com/uclouvain/openjpeg/wiki/DocJ2KCodec).
    
## Current Status
* Build status: [![Build Status](https://travis-ci.org/uclouvain/openjpeg.svg?branch=master)](https://travis-ci.org/uclouvain/openjpeg)
    
## Who are the developers ?

The library is developed and maintained by the Image and Signal Processing Group ([ISPGroup](http://sites.uclouvain.be/ispgroup/)), in the Université catholique de Louvain ([UCL](http://www.uclouvain.be/en-index.html), with the support of the [CNES](https://cnes.fr/), the [CS](http://www.c-s.fr/) company and the [intoPIX](http://www.intopix.com) company. The JPWL module has been developed by the Digital Signal Processing Lab ([DSPLab](http://dsplab.diei.unipg.it/)) of the University of Perugia, Italy ([UNIPG](http://www.unipg.it/)).

## Details on folders hierarchy

* src
  * lib
    * openjp2: contains the sources of the openjp2 library (Part 1 & 2)
    * openjpwl: contains the additional sources if you want to build a JPWL-flavoured library.
    * openjpip: complete client-server architecture for remote browsing of jpeg 2000 images.
    * openjp3d: JP3D implementation
    * openmj2: MJ2 implementation
  * bin: contains all applications that use the openjpeg library
    * common: common files to all applications
    * jp2: a basic codec
    * mj2: motion jpeg 2000 executables
    * jpip: OpenJPIP applications (server and dec server)
      * java: a Java client viewer for JPIP
    * jp3d: JP3D applications
      * tcltk: a test tool for JP3D
    * wx
      * OPJViewer: gui for displaying j2k files (based on wxWidget)
* wrapping
  * java: java jni to use openjpeg in a java program
* thirdparty: thirdparty libraries used by some applications. These libraries will be built only if there are not found on the system. Note that libopenjpeg itself does not have any dependency.
* doc: doxygen documentation setup file and man pages
* tests: configuration files and utilities for the openjpeg test suite. All test images are located in [openjpeg-data](https://github.com/uclouvain/openjpeg-data) repository.
* cmake: cmake related files

See [LICENSE](https://github.com/uclouvain/openjpeg/blob/master/LICENSE) for license and copyright information.

See [INSTALL](https://github.com/uclouvain/openjpeg/blob/master/INSTALL) for installation procedures.

See [NEWS](https://github.com/uclouvain/openjpeg/blob/master/NEWS) for user visible changes in successive releases.

## API/ABI

OpenJPEG strives to provide a stable API/ABI for your applications. As such it
only exposes a limited subset of its functions.  It uses a mechanism of
exporting/hiding functions. If you are unsure which functions you can use in
your applications, you should compile OpenJPEG using something similar to gcc:
`fvisibility=hidden` compilation flag.
See also: http://gcc.gnu.org/wiki/Visibility

On windows, MSVC directly supports export/hiding function and as such the only
API available is the one supported by OpenJPEG.