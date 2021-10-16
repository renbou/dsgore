#!/bin/bash
curl -s https://api.github.com/repos/renbou/dsgore/releases/latest | grep browser_download_url | grep dsgore | cut -d '"' -f 4 | wget -i - -O /usr/local/bin/dsgore
if [ $? -ne 0 ]; then
  echo 'Ooops! Something went wrong...'
fi
chmod +x /usr/local/bin/dsgore
echo -ne '\e[1;31mdsgore installed... time to slay some .DS_Store'"'"'s!\e[0m\n'