# Projet Go
##Installation

Utilisez le makefile pour:
- initialiser le projet : $ make build
- lancer le projet : $ make run
- supprimer le docker : $ make down

##Postman
#####pour une utilisation postman voici les différentes routes avec des examples de json:

Post: localhost:8888/api/user/new

{
"first_name":"Bruce",
"last_name":"Wayne",
"email":"t@wayne.ioa",
"password":"marthajjjj88",
"birthdate":"1995-01-01T00:00:00Z"
}


Post: localhost:8888/api/user/login

{
"email":"t@wayne.ioa",
"password":"marthajjjj88"
}


Delete: localhost:8888/api/user/delete/{uuid_voulu}

Put: localhost:8888/api/user/put/{uuid_voulu}

{
"first_name":"Brucee",
"last_name":"Waynee",
"email": "t@wayne.io",
"password":"marthajjjj77777"
}

