#!/bin/bash

#settings
winbinname="herbwebdavd.exe"
winbuildername="build-win.sh"
linuxbinname="herbwebdavd"
linuxbuildername="build-linux.sh"

if [ -z "$1" ]
then 
    basename=$(basename "$0")
    echo "Usage $basename [targetpah]"
    exit 0
fi
if [ -e "$1" ]
then 
    echo "Target path $1 exists."
    exit 0
fi
path=$(dirname "$0")
cd $path

echo "Publish to $1."
echo "Building"
bash ./$winbuildername
bash ./$linuxbuildername
echo "Creating folder $1."
mkdir $1
mkdir $1/appdata
cp ../../appdata/readme.md $1/appdata/readme.md
echo "Copying bin file."
mkdir $1/bin
cp -rpf ../../bin/$winbinname $1/bin/$winbinname
cp -rpf ../../bin/$linuxbinname $1/bin/$linuxbinname

echo "Copying system files."
cp -rpf ../../system $1/system
echo "Copying resources files."
cp -rpf ../../resources $1/resources
echo "Copying config skeleton files."
cp -rpf ../../system/configskeleton $1/config
