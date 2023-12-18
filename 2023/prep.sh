#! /bin/zsh

cd /Users/kurre/projects/advent-of-code/2023/
cp -R template $1
curl --cookie "session=53616c7465645f5f7f74711b5159679922980189b7ec39de6cbc4d68f1c917ca45405d742b1bf2e5d3a99ccf266f823474f2cf1d5ccedf440c9c081e4e5b2694" https://adventofcode.com/2023/day/$1/input -o $1/input.txt
curl --cookie "session=53616c7465645f5f7f74711b5159679922980189b7ec39de6cbc4d68f1c917ca45405d742b1bf2e5d3a99ccf266f823474f2cf1d5ccedf440c9c081e4e5b2694" https://adventofcode.com/2023/day/$1 -o temp.html && npx turndown-cli temp.html $1/assignment.md && rm temp.html
cd -