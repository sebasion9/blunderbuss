#!/bin/bash
java -Xmx500M -cp "/usr/local/lib/antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go -no-visitor -package parsing -o ../parsing Blunderbuss.g4
