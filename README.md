# date app

golang back-end; react front-end
- see jimnotes/dateapp notes for URS

## TODO

- [ ] api recipes available for users
- [ ] api for admin?
- [ ] set up db (postgres)
- [ ] matching algorithm using ML/AI
- [ ] redifine Recipe struct (backend/cmd/api) to include details for user admin & matching
- [ ] scrape data from ladders, present info with better interface, and allow users to do something on the site

### changelog

- looking into mongodb over postgres. See terms&conditions for details.
- rm backend/cmd/api
- changed backend/cmd/api to user-recipes, not movies

#### POSTGRESQL
- Open pgsql $psql
- List db $\l
- Use db $psql -d dbname, or $\c dbname
- List tables $\dt
- Describe tables $\d table_name
- To see what's inside of table: $> SELECT * FROM table_name;
