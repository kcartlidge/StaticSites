# Deploying Static Sites and your actual site(s)

* These are step by step *manual* instructions.
* They can be scripted into *Buddy*, *CodeShip* or equivalent.
* They are an *example only*; other options exist and may work better.
* Re-deploying a site *does not* require a StaticSites restart.

## Installing onto a Digital Ocean droplet

Replace the IP address in the commands below with the one for the new droplet that the commands create. As these commands assume you are SSHing as *root*, the paths for everything are assumed to be relative to the *root account's home* (eg `/root`).

Locally, these example commands should be run from the `cmd` folder.

* Add your SSH keys to your account.
* Create a new Droplet (eg Debian, minimum spec, ensuring your SSH keys are added).
* Perform the following *locally*
* Test connect to your droplet via: `ssh root@46.101.92.76` (disconnect when done)
* Upload the Static Sites *application* by: `scp ./builds/linux/StaticSites root@46.101.92.76:StaticSites`
* Compress your site locally: `zip -r ./staticsites.io.zip ../sites/staticsites.io/*`
* Upload your site by: `scp ./staticsites.io.zip root@46.101.92.76:staticsites.io.zip`
* Now continue *on the server* via `ssh` as before
* Clear the last deployed version (if any): `rm -rf sites/staticsites.io`
* Unzip your site: `unzip -d . staticsites.io.zip` (assumes your `.zip` contains a `sites` top folder)
* Start: `./StaticSites -sites sites -port 443`

### Installing unzip if it is missing from your droplet

``` sh
apt-get update && sudo apt-get upgrade
apt-get install unzip
```

## Ensuring the sites survive a reboot

Configuring supervisor:

* Still *on the server*
* Check if it is already installed: `supervisord --version`
* Install *supervisor* if missing: `apt-get install supervisor`
* Test the install: `supervisord --version`
* Clone the sample config: `echo_supervisord_conf > /etc/supervisor/conf.d/sites.conf`
* Test the config: `supervisord -n -c /etc/supervisor/conf.d/sites.conf` (ctrl-c stops)
* Add a Static Sites command:

``` sh
echo "" >> /etc/supervisor/conf.d/sites.conf
echo "[program:staticsites]" >> /etc/supervisor/conf.d/sites.conf
echo "command=/root/StaticSites -sites /root/sites -port 443" >> /etc/supervisor/conf.d/sites.conf
echo "" >> /etc/supervisor/conf.d/sites.conf
```

* Give it a try: `supervisord -n -c /etc/supervisor/conf.d/sites.conf` (ctrl-c stops)
* Whilst the above runs, the site should be being served as expected. Test it.
* See if it survives a reboot: `reboot` (no warnings)
* Check the status: `systemctl status supervisor.service`

The sites should survive a reboot. If not, the following may help:

* Enable the service: `systemctl enable supervisor.service`
* Confirm it is linked in: `ls /lib/systemd/system`
* Check it is active: `systemctl is-active supervisor.service`
* Also confirm it's status: `systemctl status supervisor.service`
* To see all available systemd services: `systemctl list-units`

Surviving crashes should happen anyway, also, but just in case:

* Edit the supervisor systemd script: `vi /etc/systemd/system/multi-user.target.wants/supervisor.service`
* Set *restart* to *always*: `Restart=always`
* Reload and restart the systemd service: `systemctl daemon-reload && systemctl restart supervisor.service`
