# sh file
# Description: Testing script for journal

#!/bin/bash
rm /Users/moustafamakhlouf/Documents/apple_interview/saved/journals.json
./journal --create "entry1" --title "Title1"
./journal --create="entry2" --title="Title2"
./journal -create "entry3" -title "Title3"
./journal -create="entry4" -title="Title4"
./journal --list
./journal --delete "2" 
./journal --delete="2" 
./journal --list
# ./journal -list
# ./journal -list=true
# ./journal --list=true

