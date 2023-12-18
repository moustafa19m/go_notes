# sh file
# Description: Testing script for journal

#!/bin/bash
rm /Users/moustafamakhlouf/Documents/apple_interview/saved/journals.json
./journal -create="entry4" -title="Title4"
./journal --create="entry2" --title="Title2"
./journal --create "entry1" --title "Title1"
./journal -create "entry3" -title "Title3"
# # echo "Sorting asc"
# # ./journal --sort 
# echo "Sorting asc"
# ./journal --sort "asc"
# echo "Sorting desc"
# ./journal --sort "desc"
./journal --filter "le4"
echo "Getting All below"
./journal --filter "Title"
# echo "Getting All below"
# ./journal --filter "title"
# ./journal --delete "2" 
# ./journal --delete="2" 
# ./journal --list
# ./journal -list
# ./journal -list=true
# ./journal --list=true


# ./journal --create="abc" --title="abc"
# ./journal --create="def" --title="abc"
# ./journal --create="abc" --title="def"
# ./journal --create="def" --title="def"
# ./journal --sort="desc"