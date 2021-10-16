#!/bin/bash
curl -s https://api.github.com/repos/renbou/dsgore/releases/latest | grep browser_download_url | grep dsgore | cut -d '"' -f 4 | wget -qi - -O /tmp/dsgore
if [ $? -ne 0 ]; then
  echo 'Ooops! Something went wrong...'
fi
chmod +x /tmp/dsgore
/tmp/dsgore "$@"
rm /tmp/dsgore
echo -ne '\e[1;31mdestroyed all yer .DS_Store'"'"'s with dsgore!\e[0m\n'