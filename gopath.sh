#!/bin/sh

# setup
project_home=$(cd $(dirname $0); pwd)
cd $project_home

# set GOPATH for the prject
project_gopath="$project_home/vendor"
case $GOPATH in
    *${project_gopath}*)
        ;;
    *)
        export GOPATH=$GOPATH:$project_gopath
        ;;
esac

cd vendor
if [ ! -e src ] ; then
    ln -sf . src
fi

# cleanup
cd $project_home
