#!/bin/bash
java -Xmx500M -cp "/usr/local/lib/antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go -visitor -package parsing -o ../cmd/parsing Blunderbuss.g4
