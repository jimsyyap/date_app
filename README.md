# date app

golang back-end; <del>react</del> vue front-end, upload friday
- see jimnotes/dateapp notes for URS
- why a challenge?

## TODO

- [ ] api recipes available for users
- [ ] api for admin?
- [ ] create dashboard for user
- [ ] create blog platform for admin manage site/app
- [ ] matching algorithm using ML/AI
- [ ] redifine Recipe struct (backend/cmd/api) to include details for user admin & matching
- [ ] scrape data from ladders, present info with better interface, and allow users to do something on the site

### changelog

- 12/09/2024 - go run main.go served vue frontend connected to postgres
- 12/09/2024 - using vue so root go run main.go opens on port:8080.
- 21aug2024 ... decided on postgres, not mongodb
- friend needs an inventory management system. 21aug2024
    - dbname is inventorydb

#### POSTGRESQL quickies
- Open pgsql $psql
- List db $\l
- Use db $psql -d dbname, or $\c dbname
- List tables $\dt
- Describe tables $\d table_name
- To see what's inside of table: $> SELECT * FROM table_name;

#### vue quickies
- npm init vue@latest > npm i > run dev
