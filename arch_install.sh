sudo pacman -S postgresql
sudo -u postgres initdb --locale $LANG -E UTF8 -D '/var/lib/postgres/data'

# enable and kick off
sudo systemctl enable postgresql.service --now

# make our database 
sudo -u postgres createdb mikes_db 

# create user
sudo -u postgres createuser --interactive

# set password
echo "Run: ALTER USER username WITH PASSWORD 'password';" in following shell:
sudo -u postgres psql

# install memcached
sudo pacman -S memcached

# enable and start 
sudo systemctl enable memcached --now

# alter memcached to listen on our local server on port 11211
sudo nano /etc/memcached.conf
#add the line "-l 127.0.0.1 -p 11211"

# install pgmemcache extension for postgresql
sudo pacman -S postgresql-pgmemcache-git
# NOTE: this might not work, I had to build from source:
    # https://github.com/ohmu/pgmemcache


sudo nano /var/lib/postgres/data/postgresql.conf
# add shared_preload_libraries = 'pgmemcache'
sudo systemctl restart postgres

# IT WORKS!! that was so easy!