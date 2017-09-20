# Deploying Static Sites and your actual site(s)

* These are step by step *manual* instructions.
* They can be scripted into *CodeShip* or equivalent.
* They are an example only; other options exist.

## Installing onto a Digital Ocean droplet

Replace the IP address in the commands below with the one for the new droplet that the commands create. As these commands assume you are SSHing as *root*, the paths for everything are assumed to be relative to the *root* account's *home* (eg ```/root```).

* Add your SSH keys to your account.
* Create a new Droplet (Debian, minimum spec, ensure your SSH keys are added).
* Test your droplet via: ```ssh root@46.101.92.76```
* Upload the Static Sites application by: ```scp ./StaticSites root@46.101.92.76:StaticSites```
* Compress your site locally: ```zip -r kcartlidge.zip kcartlidge.com/*```
* Upload your site by: ```scp ./kcartlidge.zip root@46.101.92.76:kcartlidge.zip```
* Clear the last deployed version (if any): ```rm -rf sites/kcartlidge```
* Unzip your site *on the server*: ```unzip -d sites kcartlidge.zip```
* If the first time round, add a config: ```echo "kcartlidge.com = /root/sites/kcartlidge.com" > sites/sites.ini```
* Start: ```./StaticSites -sites sites/sites.ini -port 443```

Note that the ```sites.ini``` configuration should have *full paths* for the site folders, so regardless of the launching process they can be found. 

### Installing unzip if it is missing from your droplet

``` sh
apt-get update && sudo apt-get upgrade
apt-get install unzip
```

## Ensuring the sites survive a reboot

Configuring supervisor:

* Check if it is already installed: ```supervisord --version```
* Install *supervisor*: ```apt-get install supervisor```
* Test the install: ```supervisord --version```
* Clone the sample config: ```echo_supervisord_conf > /etc/supervisor/conf.d/sites.conf```
* Test the config: ```supervisord -n -c /etc/supervisor/conf.d/sites.conf``` (ctrl-c stops)
* Create a Static Sites command:
``` sh
echo "" >> /etc/supervisor/conf.d/sites.conf
echo "[program:staticsites]" >> /etc/supervisor/conf.d/sites.conf
echo "command=/root/StaticSites -sites /root/sites/sites.ini -port 443" >> /etc/supervisor/conf.d/sites.conf
echo "" >> /etc/supervisor/conf.d/sites.conf
```
* Give it a try: ```supervisord -n -c /etc/supervisor/conf.d/sites.conf``` (ctrl-c stops)
* Whilst the above runs, the site should be being served as expected. Test it.
* See if it survives a reboot: ```reboot``` (no warnings)
* Check the status: ```systemctl status supervisor.service```

The sites should survive a reboot. If not, the following may help:

* Enable the service: ```systemctl enable supervisor.service```
* Confirm it is linked in: ```ls /lib/systemd/system```
* Check it is active: ```systemctl is-active supervisor.service```
* Also confirm it's status: ```systemctl status supervisor.service```
* To see all available systemd services: ```systemctl list-units```

Surviving crashes should happen anyway, also, but just in case:

* Edit the supervisor systemd script: ```vi /etc/systemd/system/multi-user.target.wants/supervisor.service```
* Set *restart* to *always*: ```Restart=always```
* Reload and restart the systemd service: ```systemctl daemon-reload && systemctl restart supervisor.service```

