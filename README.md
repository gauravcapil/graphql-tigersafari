#Using go => 1.21.6
#Dev Environment used => Windows 11
#go installation link: https://dl.google.com/go/go1.21.6.windows-amd64.msi


following link : https://www.apollographql.com/blog/using-graphql-with-golang for bootstrapping the service


#the script (init_service.sh)
follows the above link step by step for a quick installation

#to setup the DB
run initialize_postgres.sh

##PRE-REQUISITES:
Docker for desktop should me installed as we will be running a containerized postgres for this project.
Follow: https://docs.docker.com/desktop/install/windows-install/ for necessary actions
For e2e testing we required "jq" install using apt-get install jq

##To CONFIGURE PORTS AND DB CREDENTIALS:
open and configure appropriate values in config.cfg

##To Force migration for the first time
( to setup the tables, export MIGRATE_1=1)
( to truncate the tables, export MIGRATE_DROPALL=1)
More migration scripts can be enabled with addition of these env tags and implementing the sequence in the same way for future

##To RUN:

using moba-xterm or cygwin:
./run.sh

##To change and update the schema ( if required )
./updateSchema.sh

##To run an e2e test you can use the scripts in e2eTests folder
./start_e2eTests.sh

#For email via gmail support, follow:
https://support.google.com/accounts/answer/185833