curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq ' .[] | select( ID == '$HERO_ID' ) | .connections .relatives' | cut -d '"' -f2
