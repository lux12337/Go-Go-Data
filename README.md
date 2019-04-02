# Part 1: Setting up the local Postgres database

**Note:** The project was developed on a Mac

1. Install docker from : https://docs.docker.com/docker-for-mac/install/

2. Install Postgres in a docker container by running in terminal: `docker run -d --name postgresdb -p 5432:5432 postgres:11`
  -The name of the docker container is `postgresdb`
  -The database will be listening on the port 5432
  -The postgres version is 11

3. Open bash inside of the container: `docker exec -it postgresdb /bin/bash`

4. Once inside the container, run: `psql -U postgres`

5. Create a database called 'topos': 'CREATE DATABASE topos;'

6. Connect to it: `\c topos`

7. Create the 'ny_data' table to hold the data from the API: `CREATE TABLE IF NOT EXISTS ny_data(id serial PRIMARY KEY, bin INT, cnstrct_yr INT, lstmoddate VARCHAR(512), lststatype VARCHAR(512), doitt_id INT, heightroof DECIMAL, feat_code INT, groundelev INT, shape_area DECIMAL, shape_len INT, base_bbl VARCHAR(512), mpluto_bbl VARCHAR(512), geomsource VARCHAR(512));`
