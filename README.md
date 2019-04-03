# Part 1: Setting up the local Postgres database

**Note:** The project was developed on a Mac

1. Install docker from : https://docs.docker.com/docker-for-mac/install/

2. Install Postgres in a docker container by running in terminal: `docker run -d --name postgresdb -e POSTGRES_PASSWORD=your-password -p 5432:5432 postgres:11`
  - The name of the docker container is `postgresdb`
  - The database will be listening on the port 5432
  - The postgres version is 11
  - **Note:** Put a password of your choice instead of 'your-password'

3. Open bash inside of the container: `docker exec -it postgresdb /bin/bash`

4. Once inside the container, run: `psql -U postgres`

5. Create a database called 'topos': 'CREATE DATABASE topos;'

6. Connect to it: `\c topos`

7. Create the 'ny_data' table to hold the data from the API: `CREATE TABLE IF NOT EXISTS ny_data(id serial PRIMARY KEY, bin INT, cnstrct_yr INT, lstmoddate VARCHAR(512), lststatype VARCHAR(512), doitt_id INT, heightroof DECIMAL, feat_code INT, groundelev INT, shape_area DECIMAL, shape_len DECIMAL, base_bbl VARCHAR(512), mpluto_bbl VARCHAR(512), geomsource VARCHAR(512));`


# Part 2: Setting up the environment for Go

1. Follow tutorial to setup go environment for Mac: https://www.youtube.com/watch?v=I5XCvYs0tGo

2. Download this repository and move the folder 'topos' to your src folder for Go

3. Navigate to 'topos' folder and edit the '.env' by entering the password for your postgres database instead of 'password'

4. Run `./dependencies.sh` to fetch all the dependencies needed to run the code

5. Run `go run api-server.go` to start the API server
