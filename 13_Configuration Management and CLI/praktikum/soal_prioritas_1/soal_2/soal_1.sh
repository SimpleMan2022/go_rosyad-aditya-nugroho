#!/bin/sh

mkdir foods drinks snacks

touch foods/menu.txt drinks/menu.txt snacks/menu.txt

echo "bubur sumsum" >> foods/menu.txt
echo "risol ayam" >> foods/menu.txt
echo "nasi padang" >> foods/menu.txt

echo "jus jeruk" >> drinks/menu.txt
echo "kopi tubruk" >> drinks/menu.txt
echo "air mineral" >> drinks/menu.txt

cd snacks

wget https://gist.githubusercontent.com/nadirbslmh/c84ee3527272e52a8e6257d46e627f91/raw/74593cc3fe61ca4ff36e11ed8e3fffcfb1d0c602/snacks.json

cat snacks.json >> menu.txt

rm snacks.json
