# sh file
# Description: Testing script for journal

#!/bin/bash
rm /Users/moustafamakhlouf/Documents/go_notes/saved/journals.json



./journal --create="word1 word2 word1" --title="abc" --tags="summer,trip"
./journal --create="words words words" --title="word2"  --tags="winter,trip"
./journal --create="word3" --title="def"
./journal --create="def" --title="def" --tags="trip"

./journal --analyze