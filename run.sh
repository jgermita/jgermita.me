#!/bin/sh
ls -R -a

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/jgermita.me.exe" -importPath github.com/jgermita/jgermita.me -srcPath "$SCRIPTPATH/src" -runMode dev
