
curl https://learn.zone01dakar.sn/assets/superhero/all.json | jq ' .[] | select( .id == 70) | .name'
